package index

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/fs"
	"log"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/daniilty/hackaton-keysystems-ftp-proxy/internal/model"
	"github.com/daniilty/hackaton-keysystems-ftp-proxy/internal/repository"
	"github.com/daniilty/hackaton-keysystems-ftp-proxy/internal/transport/mq/rabbitmq/publisher"
	"github.com/secsy/goftp"
)

const bufSize = 10 * 1024 * 1024

type Indexer struct {
	bufPool   *sync.Pool
	interval  time.Duration
	client    *goftp.Client
	path      string
	cache     *sumCache
	publisher publisher.Publisher
	db        repository.DB
	syncer    *syncer
}

func NewIndexer(interval time.Duration, syncInterval time.Duration, client *goftp.Client, publisher publisher.Publisher, db repository.DB, path string) *Indexer {
	return &Indexer{
		interval: interval,
		client:   client,
		cache:    newSumCache(),
		bufPool: &sync.Pool{
			New: func() any {
				buf := make([]byte, bufSize)
				return &buf
			},
		},
		path:      path,
		publisher: publisher,
		db:        db,
		syncer:    newSyncer(syncInterval, db),
	}
}

func (i *Indexer) initCache(ctx context.Context) error {
	sums, err := i.db.GetSums(ctx)
	if err != nil {
		return fmt.Errorf("get db sums: %w", err)
	}

	for _, s := range sums {
		i.cache.set(s.Name, s.Sum)
	}

	log.Println("initialized cache, records:", len(sums))
	return nil
}

func (i *Indexer) Run(ctx context.Context) error {
	ticker := time.NewTicker(i.interval)

	err := i.initCache(ctx)
	if err != nil {
		return fmt.Errorf("init cache: %w", err)
	}

	go func() {
		err := i.syncer.run(ctx)
		if err != nil {
			log.Println("run syncer", err)
		}
	}()

	for {
		err := i.checkDirs()
		if err != nil {
			log.Println("check dirs", err)
		}

		select {
		case <-ticker.C:
		case <-ctx.Done():
			log.Println("context done, exiting")
			return nil
		}
	}
}

func (i *Indexer) checkDirs() error {
	dir, err := i.client.ReadDir(i.path)
	if err != nil {
		return fmt.Errorf("read dir: %w", err)
	}

	return i.readDir(i.path, dir)
}

func (i *Indexer) readDir(name string, dir []fs.FileInfo) error {
	log.Println("read dir", name)
	for _, info := range dir {
		if info.IsDir() {
			dPath := path.Join(name, info.Name())
			d, err := i.client.ReadDir(dPath)
			if err != nil {
				return fmt.Errorf("read dir: %w", err)
			}

			go func() {
				err = i.readDir(dPath, d)
				if err != nil {
					log.Println(err)
				}
			}()

			continue
		}

		log.Println("handle file", path.Join(name, info.Name()))
		err := i.handleFile(path.Join(name, info.Name()), info.Name())
		if err != nil {
			return fmt.Errorf("handle file: %w", err)
		}
	}

	return nil
}

func (i *Indexer) handleFile(name string, fName string) error {
	var buf []byte

	bufP, ok := i.bufPool.Get().(*[]byte)
	if !ok {
		buf = make([]byte, bufSize)
	} else {
		buf = *bufP
	}

	buffer := bytes.NewBuffer(buf)
	err := i.client.Retrieve(name, buffer)
	if err != nil {
		return fmt.Errorf("retrieve file: %w", err)
	}

	if len(name) < 4 {
		log.Println("bad file", name)
		return nil
	}

	log.Println("retrieve file", name, fName)
	buf = buffer.Bytes()
	defer func() {
		i.bufPool.Put(&buf)
	}()

	switch strings.ToLower(name[len(name)-3:]) {
	case "zip":
		bufferAt := bytes.NewReader(buf)

		zipReader, err := zip.NewReader(bufferAt, int64(len(buf)))
		if err != nil {
			return fmt.Errorf("zip new reader: %w", err)
		}

		for _, f := range zipReader.File {
			if f.FileInfo().IsDir() {
				continue
			}

			fReader, err := f.Open()
			if err != nil {
				log.Println("failed to open zip file:", err)

				continue
			}

			bb, err := io.ReadAll(fReader)
			if err != nil {
				fReader.Close()
				log.Println("read file:", err)

				continue
			}

			md5Sum := fmt.Sprintf("%x", md5.Sum(bb))
			oldSum := i.cache.get(f.Name)
			if md5Sum != oldSum {
				i.cache.set(f.Name, md5Sum)

				i.syncer.pushOp(&model.MD5Sum{
					Name: f.Name,
					Sum:  md5Sum,
				})

				xmlDec := xml.NewDecoder(bytes.NewReader(bb))
				data := &model.Data{}
				err = xmlDec.Decode(data)
				if err != nil {
					log.Println("xml decode:", err)
					fReader.Close()

					continue
				}

				bb, err = json.Marshal(&data.Contract)
				if err != nil {
					log.Println("json marshal:", err)
					fReader.Close()

					continue
				}
				log.Println(string(bb))

				err = i.publisher.SendContract(bb)
				if err != nil {
					log.Println("send contract amqp:", err)
				}
			}

			fReader.Close()
		}
	case "xml":
		md5Sum := fmt.Sprintf("%x", md5.Sum(buf))
		oldSum := i.cache.get(fName)
		if md5Sum != oldSum {
			i.cache.set(fName, md5Sum)

			i.syncer.pushOp(&model.MD5Sum{
				Name: fName,
				Sum:  md5Sum,
			})

			xmlDec := xml.NewDecoder(bytes.NewReader(buf))
			data := &model.Data{}
			err = xmlDec.Decode(data)
			if err != nil {
				return fmt.Errorf("xml decode: %w", err)
			}

			bb, err := json.Marshal(&data.Contract)
			if err != nil {
				return fmt.Errorf("json marshal: %w", err)
			}

			err = i.publisher.SendContract(bb)
			if err != nil {
				return fmt.Errorf("send contract amqp: %w", err)
			}
		}
	}

	return nil
}
