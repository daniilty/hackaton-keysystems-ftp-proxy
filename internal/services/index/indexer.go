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
}

func NewIndexer(interval time.Duration, client *goftp.Client, publisher publisher.Publisher, path string) *Indexer {
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
	}
}

func (i *Indexer) Run(ctx context.Context) {
	ticker := time.NewTicker(i.interval)

	for {
		err := i.checkDirs()
		if err != nil {
			log.Println("check dirs", err)
		}

		select {
		case <-ticker.C:
		case <-ctx.Done():
			log.Println("context done, exiting")
			return
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
		err := i.handleFile(path.Join(name, info.Name()), info.Size())
		if err != nil {
			return fmt.Errorf("handle file: %w", err)
		}
	}

	return nil
}

func (i *Indexer) handleFile(name string, size int64) error {
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

	log.Println("retrieve file", name, size)
	buf = buffer.Bytes()
	defer func() {
		i.bufPool.Put(&buf)
	}()
	bufferAt := bytes.NewReader(buf)
	switch strings.ToLower(name[len(name)-3:]) {
	case "zip":
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

				err = i.publisher.SendContract(bb)
				if err != nil {
					log.Println("send contract amqp:", err)
				}
			}

			fReader.Close()
		}
	case "xml":
		dec := xml.NewDecoder(buffer)
		dat := &model.Data{}

		err = dec.Decode(dat)
		if err != nil {
			return fmt.Errorf("xml decode: %w", err)
		}

		bb, err := json.Marshal(dat.Contract)
		if err != nil {
			return fmt.Errorf("json marshal: %w", err)
		}

		log.Println(string(bb))
	}
	return nil
}
