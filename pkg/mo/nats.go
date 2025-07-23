package mo

import (
	"fmt"
	"lantak/config"

	"github.com/nats-io/nats.go"
)

type Server struct {
	nc   *nats.Conn
	subj string
	user string
}

func NewNats(user, subj string, cfg config.Config) *Server {
	nc, err := nats.Connect(fmt.Sprintf("%s:%d", cfg.NatsServer.Address, cfg.NatsServer.Port))
	if err != nil {
		panic("cannot create nats server: " + err.Error())
	}
	return &Server{
		nc:   nc,
		subj: subj,
		user: user,
	}
}

func (s *Server) Sub() {
	s.nc.Subscribe(s.subj, func(msg *nats.Msg) {
		fmt.Printf("%s\n", msg.Data)
	})
}

func (s *Server) Pub(msg string) {
	s.nc.Publish(s.subj, fmt.Appendf(nil, "%s > %s", s.user, msg))
}

func (s *Server) Disconnect() {
	s.nc.Flush()
	s.nc.Close()
}
