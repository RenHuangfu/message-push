package di

import (
	"github.com/google/wire"
	"message-push/app/manager/service/internal/server"
	"message-push/app/manager/service/internal/service"
)

var (
	ServerProviderSet = wire.NewSet(
		server.NewKafkaServer,
	)

	ServiceProviderSet = wire.NewSet(
		service.NewDemo,
	)

	UsecaseProviderSet = wire.NewSet()

	DataProviderSet = wire.NewSet()
)
