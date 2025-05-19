package server

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"message-push/app/pusher/service/internal/conf"
	"message-push/app/pusher/service/internal/handler"
	"os"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	c *conf.Bootstrap,
	handler *handler.WebSocketHandler,
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
		opts = append(opts, http.Address(fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))))
	}
	if server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(server.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	srv.Handle("/ws", handler)
	return srv
}
