package postgres

import (
	"context"

	"github.com/daniilty/hackaton-keysystems-ftp-proxy/internal/model"
)

func (d *db) UpsertSum(ctx context.Context, sum *model.MD5Sum) error {
	const q = "insert into sums(name, sum) values(:name, :sum) on conflict(name) do update set sum=:sum"

	_, err := d.db.NamedExecContext(ctx, q, sum)
	return err
}

func (d *db) GetSums(ctx context.Context) ([]*model.MD5Sum, error) {
	const q = "select * from sums"

	res := []*model.MD5Sum{}

	err := d.db.SelectContext(ctx, &res, q)
	if err != nil {
		return nil, err
	}

	return res, err
}
