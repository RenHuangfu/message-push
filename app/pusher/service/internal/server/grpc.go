package server

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "message-push/app/pusher/service/api/server/v1"
	"message-push/app/pusher/service/internal/conf"
	"message-push/app/pusher/service/internal/service"
	"os"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Bootstrap, tsv *service.SendService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Server.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Server.Grpc.Network))
	}
	if c.Server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(fmt.Sprintf("0.0.0.0:%s", os.Getenv("GRPC_PORT"))))
	}
	if c.Server.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Server.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterDirectSendServiceServer(srv, tsv)
	return srv
}
