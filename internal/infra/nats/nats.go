package nats

import (
	"fmt"
	"lantak/config"

	"github.com/nats-io/nats.go"
)

func NewNats(cfg config.Config) *nats.Conn {
	nc, err := nats.Connect(fmt.Sprintf("%s:%d", cfg.NatsServer.Address, cfg.NatsServer.Port))
	if err != nil {
		panic("cannot create nats server: " + err.Error())
	}
	return nc
}
