package di

import (
	"github.com/google/wire"
	"message-push/app/pusher/service/internal/server"
	"message-push/app/pusher/service/internal/service"
)

var (
	ServerProviderSet = wire.NewSet(
		server.NewGRPCServer,
	)

	ServiceProviderSet = wire.NewSet(
		service.NewDemo,
	)
)
