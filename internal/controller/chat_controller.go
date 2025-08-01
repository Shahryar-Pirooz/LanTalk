package controller

import (
	"bufio"
	"fmt"
	"lantak/internal/model"
	"lantak/internal/service"
	"os"
)

type ChatController struct {
	s    service.MessageService
	subj string
	user *model.User
}

func NewChatController(s service.MessageService) *ChatController {
	return &ChatController{
		s:    s,
		user: new(model.User),
	}
}

func (c *ChatController) PromptUserInput() {
	var name string
	fmt.Print("\nWhat is your subject: ")
	fmt.Scan(&c.subj)
	fmt.Print("\nEnter your name: ")
	fmt.Scan(&name)
	c.user.Name = name
	c.s.Subscribe(c.subj)
}

func (c *ChatController) HandleInput() {
	for {
		buf := bufio.NewReader(os.Stdin)
		i, _ := buf.ReadString('\n')
		if i == "/exit\n" {
			c.s.Disconnect()
			break
		}
		c.s.Publish(c.subj, i, *c.user)
	}
}
