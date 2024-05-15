package nats

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/wlcmtunknwndth/hackBPA/internal/config"
	"github.com/wlcmtunknwndth/hackBPA/internal/storage"
	"time"
)

type Storage interface {
	GetEvent(context.Context, uint) (*storage.Event, error)
	DeleteEvent(context.Context, uint) error
	CreateEvent(context.Context, *storage.Event) (uint, error)
	PatchEvent(context.Context, *storage.Event) error
}

type Nats struct {
	b  *nats.Conn
	db Storage
}

func New(cfg *config.Nats, db Storage) (*Nats, error) {
	const op = "broker.nats.New"

	natsService, err := nats.Connect(cfg.Address,
		nats.RetryOnFailedConnect(cfg.Retry),
		nats.MaxReconnects(cfg.MaxReconnects),
		nats.ReconnectWait(cfg.ReconnectWait),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err = natsService.Flush(); err != nil {
		return nil, fmt.Errorf("%s: flush: %w", op, err)
	}
	if err = natsService.FlushTimeout(time.Second); err != nil {
		return nil, fmt.Errorf("%s: flush timeout: %w", op, err)
	}
	return &Nats{b: natsService, db: db}, nil
}

func (n *Nats) Close() {
	n.b.Close()
}
