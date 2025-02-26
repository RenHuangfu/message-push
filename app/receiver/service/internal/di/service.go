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
		service.NewDemo,
	)

	UsecaseProviderSet = wire.NewSet(
		usecase.NewDemo,
	)

	DataProviderSet = wire.NewSet(
		data.NewData,
		data.NewDemo,
	)
)
