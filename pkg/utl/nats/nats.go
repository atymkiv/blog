package nats

import (
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/config"
	"github.com/nats-io/go-nats"
)

func New(cfg *config.Nats) (*nats.Conn, error) {
	natsClient, err := nats.Connect(cfg.Host)

	return natsClient, err
}
