//go:build wireinject
// +build wireinject

package wire

import (
	"lantak/config"
	"lantak/internal/controller"
	"lantak/internal/infra/nats"
	"lantak/internal/service"

	"github.com/google/wire"
)

func InitChatController(cfg config.Config) *controller.ChatController {
	wire.Build(
		nats.NewNats,
		service.NewMessageService,
		controller.NewChatController,
	)
	return nil
}
