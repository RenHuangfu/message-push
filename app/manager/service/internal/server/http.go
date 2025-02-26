package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"message-push/app/manager/service/internal/conf"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	c *conf.Bootstrap,
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
	return srv
}
