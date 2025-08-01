package service

import (
	"fmt"
	"lantak/internal/model"

	"github.com/nats-io/nats.go"
)

type MessageService interface {
	Subscribe(subj string)
	Publish(subj, msg string, user model.User)
	Disconnect()
}

type messageService struct {
	nc *nats.Conn
	// chat []model.Message
}

func NewMessageService(nc *nats.Conn) MessageService {
	return &messageService{
		nc: nc,
	}
}

func (s *messageService) Subscribe(subj string) {
	s.nc.Subscribe(subj, func(msg *nats.Msg) {
		fmt.Printf("%s\n", msg.Data)
	})
}

func (s *messageService) Publish(subj, msg string, user model.User) {
	s.nc.Publish(subj, fmt.Appendf(nil, "%s > %s", user.Name, msg))
}

func (s *messageService) Disconnect() {
	s.nc.Flush()
	s.nc.Close()
}
