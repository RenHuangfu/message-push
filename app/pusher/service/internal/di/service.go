package di

import (
	"github.com/google/wire"
	"message-push/app/pusher/service/internal/handler"
	"message-push/app/pusher/service/internal/server"
	"message-push/app/pusher/service/internal/service"
)

var (
	ServerProviderSet = wire.NewSet(
		server.NewKafkaServer,
		server.NewHTTPServer,
	)

	HandlerProviderSet = wire.NewSet(
		handler.NewWebSocketHandler,
	)

	ServiceProviderSet = wire.NewSet(
		service.NewSendService)
)
