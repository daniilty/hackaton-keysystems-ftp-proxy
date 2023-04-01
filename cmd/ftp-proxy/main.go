package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os/signal"
	"syscall"
	"time"

	"github.com/daniilty/hackaton-keysystems-ftp-proxy/internal/config"
	"github.com/daniilty/hackaton-keysystems-ftp-proxy/internal/repository/postgres"
	"github.com/daniilty/hackaton-keysystems-ftp-proxy/internal/services/index"
	"github.com/daniilty/hackaton-keysystems-ftp-proxy/internal/transport/mq/rabbitmq/publisher"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/secsy/goftp"
)

// Just for fun, walk an ftp server in parallel. I make no claim that this is
// correct or a good idea.
func main() {
	f := initFlags()

	cfg, err := config.Init(f.configPath)
	if err != nil {
		panic(err)
	}

	ftpCfg := goftp.Config{
		User:               cfg.UserName,
		Password:           cfg.Password,
		ConnectionsPerHost: cfg.ConnsPerHost,
		Timeout:            time.Duration(cfg.TimeoutSeconds) * time.Second,
	}

	client, err := goftp.DialConfig(ftpCfg, cfg.Host)
	if err != nil {
		panic(err)
	}

	rconn, err := amqp.Dial(cfg.RabbitConnAddr)
	if err != nil {
		panic(err)
	}

	channel, err := rconn.Channel()
	if err != nil {
		panic(err)
	}

	pub := publisher.NewPublisher(channel, cfg.Exchange)
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	db, err := postgres.Connect(ctx, cfg.PGDSN)
	if err != nil {
		panic(err)
	}

	indexer := index.NewIndexer(5*time.Hour, time.Second*time.Duration(cfg.SyncerIntervalSeconds), client, pub, db, cfg.Path)

	go func() {
		err = indexer.Run(ctx)
		if err != nil {
			panic(err)
		}
	}()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
