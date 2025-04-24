package server

import (
	v1 "message-push/app/receiver/service/api/server/v1"
	"message-push/app/receiver/service/internal/conf"
	"message-push/app/receiver/service/internal/service"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	c *conf.Bootstrap,
	business *service.BusinessService,
	message *service.MessageService,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	server := c.Server
	if server.Http.Network != "" {
		opts = append(opts, http.Network(server.Http.Network))
	}
	if server.Http.Addr != "" {
		opts = append(opts, http.Address(server.Http.Addr))
	}
	if server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(server.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterBusinessHTTPServer(srv, business)
	v1.RegisterMessageHTTPServer(srv, message)
	return srv
}
