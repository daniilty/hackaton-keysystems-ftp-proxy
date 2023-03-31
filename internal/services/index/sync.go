package index

import (
	"context"
	"sync"
	"time"

	"github.com/daniilty/hackaton-keysystems-ftp-proxy/internal/model"
	"github.com/daniilty/hackaton-keysystems-ftp-proxy/internal/repository"
)

type syncer struct {
	interval time.Duration
	mux      *sync.Mutex
	ops      []*model.MD5Sum

	db repository.DB
}

func newSyncer(interval time.Duration, db repository.DB) *syncer {
	return &syncer{
		interval: interval,
		mux:      &sync.Mutex{},
		ops:      []*model.MD5Sum{},
		db:       db,
	}
}

func (s *syncer) run(ctx context.Context) error {
	ticker := time.NewTicker(s.interval)

	for {
		err := s.writeOps(context.Background())
		if err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			err := s.writeOps(context.Background())
			if err != nil {
				return err
			}

			return nil
		case <-ticker.C:
		}
	}
}

func (s *syncer) writeOps(ctx context.Context) error {
	for {
		op, ok := s.popOp()
		if !ok {
			return nil
		}

		s.db.UpsertSum(ctx, op)
	}
}

func (s *syncer) popOp() (*model.MD5Sum, bool) {
	s.mux.Lock()
	n := len(s.ops)
	if n == 0 {
		s.mux.Unlock()
		return nil, false
	}

	op := s.ops[len(s.ops)-1]
	s.ops = s.ops[:len(s.ops)-1]
	s.mux.Unlock()

	return op, true
}

func (s *syncer) pushOp(op *model.MD5Sum) {
	s.mux.Lock()
	s.ops = append(s.ops, op)
	s.mux.Unlock()
}
