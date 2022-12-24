package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
)

type Config struct {
	Url string
}

func NewNatsEventStoreConn(cfg *Config) (*nats.Conn, error) {
	natsConn, err := nats.Connect(cfg.Url)

	if err != nil {
		return nil, errors.Wrap(err, "pkg.nats.NewConn")
	}

	return natsConn, nil
}
