package postgres

import (
	"context"

	"github.com/daniilty/hackaton-keysystems-ftp-proxy/internal/repository"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(ctx context.Context, addr string) (repository.DB, error) {
	d, err := sqlx.ConnectContext(ctx, "postgres", addr)
	if err != nil {
		return nil, err
	}

	return &db{
		db: d,
	}, nil
}

type db struct {
	db *sqlx.DB
}
