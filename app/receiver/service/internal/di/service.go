package di

import (
	"github.com/google/wire"
	"message-push/app/receiver/service/internal/data"
	"message-push/app/receiver/service/internal/server"
	"message-push/app/receiver/service/internal/service"
	"message-push/app/receiver/service/internal/usecase"
)

var (
	ServerProviderSet = wire.NewSet(
		server.NewHTTPServer,
	)

	ServiceProviderSet = wire.NewSet(
		service.NewBusinessService,
		service.NewMessageService,
	)

	UsecaseProviderSet = wire.NewSet(
		usecase.NewBusinessUsecase,
		usecase.NewMessageUsecase,
	)

	DataProviderSet = wire.NewSet(
		data.NewBusinessRepo,
		data.NewMessageRepo,
		data.NewData,
	)
)
