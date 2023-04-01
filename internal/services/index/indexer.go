package index

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"path"
	"strings"
	"time"

	"github.com/daniilty/hackaton-keysystems-ftp-proxy/internal/model"
	"github.com/daniilty/hackaton-keysystems-ftp-proxy/internal/repository"
	"github.com/daniilty/hackaton-keysystems-ftp-proxy/internal/transport/mq/rabbitmq/publisher"
	"github.com/secsy/goftp"
	"github.com/valyala/bytebufferpool"
)

type Indexer struct {
	bufPool   *bytebufferpool.Pool
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
		interval:  interval,
		client:    client,
		cache:     newSumCache(),
		bufPool:   &bytebufferpool.Pool{},
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

			err = i.readDir(dPath, d)
			if err != nil {
				log.Println(err)
			}

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
	buf := i.bufPool.Get()
	byteBuf := bytes.NewBuffer(buf.B)

	err := i.client.Retrieve(name, byteBuf)
	if err != nil {
		return fmt.Errorf("retrieve file: %w", err)
	}

	buf.B = byteBuf.Bytes()

	if len(name) < 4 {
		log.Println("bad file", name)
		return nil
	}

	defer func() {
		i.bufPool.Put(buf)
	}()

	if isZip(fName) {
		buf.B, err = i.handleZip(buf.B)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *Indexer) handleZip(buf []byte) ([]byte, error) {
	bufferAt := bytes.NewReader(buf)

	zipReader, err := zip.NewReader(bufferAt, int64(len(buf)))
	if err != nil {
		return buf, fmt.Errorf("zip new reader: %w", err)
	}

	for _, f := range zipReader.File {
		if f.FileInfo().IsDir() {
			continue
		}

		fIsXml := isXml(f.Name)
		fIsZip := isZip(f.Name)

		if !fIsXml && !fIsZip {
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

		if len(bb) == 0 {
			fReader.Close()
			continue
		}

		if fIsZip {
			_, err = i.handleZip(bb)
			if err != nil {
				log.Println("handle zip", err)
			}

			fReader.Close()
			continue
		}

		md5Sum := fmt.Sprintf("%x", md5.Sum(bb))
		oldSum := i.cache.get(f.Name)
		if md5Sum != oldSum {
			xmlDec := xml.NewDecoder(bytes.NewReader(bb))
			data := &model.Data{}
			err = xmlDec.Decode(data)
			if err != nil {
				if errors.Is(err, io.EOF) {
					log.Println("eof:", err)
				}
				fReader.Close()

				continue
			}

			var pubFn func(context.Context, []byte) error
			if data.Contract != nil {
				pubFn = i.publisher.SendContract
				bb, err = json.Marshal(data.Contract)
			} else if data.ContractProcedure != nil {
				pubFn = i.publisher.SendContractProcedure
				bb, err = json.Marshal(data.ContractProcedure)
			} else {
				fReader.Close()
				continue
			}

			if err != nil {
				log.Println("json marshal:", err)
				fReader.Close()

				continue
			}

			err = pubFn(context.Background(), bb)
			if err != nil {
				log.Println("send contract amqp:", err)
				fReader.Close()

				continue
			}

			i.cache.set(f.Name, md5Sum)

			i.syncer.pushOp(&model.MD5Sum{
				Name: f.Name,
				Sum:  md5Sum,
			})
		}

		fReader.Close()
	}

	return buf, nil
}

func isZip(name string) bool {
	return strings.HasSuffix(strings.ToLower(name), ".zip")
}

func isXml(name string) bool {
	return strings.HasSuffix(strings.ToLower(name), ".xml")
}
