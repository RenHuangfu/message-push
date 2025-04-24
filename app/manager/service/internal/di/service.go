package di

import (
	"github.com/google/wire"
	"message-push/app/manager/service/internal/data"
	"message-push/app/manager/service/internal/server"
	"message-push/app/manager/service/internal/service"
	"message-push/app/manager/service/internal/usecase"
)

var (
	ServerProviderSet = wire.NewSet(
		server.NewGRPCServer,
	)

	ServiceProviderSet = wire.NewSet(
		service.NewTriggerService,
	)

	UsecaseProviderSet = wire.NewSet(
		usecase.NewTriggerUsecase,
	)

	DataProviderSet = wire.NewSet(
		data.NewTriggerRepo,
		data.NewData)
)
