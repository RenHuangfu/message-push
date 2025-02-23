package di

import (
	"github.com/google/wire"
	"message-push/app/access/service/internal/data"
	"message-push/app/access/service/internal/server"
	"message-push/app/access/service/internal/service"
	"message-push/app/access/service/internal/usecase"
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
