package di

import (
	"github.com/google/wire"
	"message-push/app/pusher/service/internal/server"
)

var (
	ServerProviderSet = wire.NewSet(
		server.NewKafkaServer,
	)

	ServiceProviderSet = wire.NewSet()
)
