package model

import "time"

type Message struct {
	ID       uint
	Sender   string
	Message  string
	CreateAt time.Time
}
