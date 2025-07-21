package domain

import "time"

type Message struct {
	ID     uint
	Name   string
	IP     string
	JoinAT time.Time
	IsSelf bool
}
