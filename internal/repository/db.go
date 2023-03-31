package repository

import (
	"context"

	"github.com/daniilty/hackaton-keysystems-ftp-proxy/internal/model"
)

type DB interface {
	UpsertSum(context.Context, *model.MD5Sum) error
	GetSums(context.Context) ([]*model.MD5Sum, error)
}
