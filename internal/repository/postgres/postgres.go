package postgres

import (
	"context"

	"github.com/daniilty/hackaton-keysystems-ftp-proxy/internal/repository"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func Connect(ctx context.Context, addr string) (repository.DB, error) {
	d, err := sqlx.ConnectContext(ctx, "pgx", addr)
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
