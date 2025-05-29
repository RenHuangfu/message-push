package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "github.com/go-sql-driver/mysql"
	"message-push/app/manager/common/model/po/ent"
	v1 "message-push/app/manager/service/api/server/v1"
	"message-push/app/manager/service/internal/conf"
	"time"
)

type Data struct {
	c    *conf.Bootstrap
	db   *ent.Client
	grpc v1.DirectSendServiceClient
}

func NewData(c *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	entClient, err := ent.Open(c.Data.Database.Driver, c.Data.Database.Source)
	if err != nil {
		return nil, nil, err
	}

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(c.Data.Grpc.Endpoint), // 从配置读取地址
		grpc.WithMiddleware( // 中间件链
			tracing.Client(),
			logging.Client(logger),
		),
		grpc.WithDiscovery(nil),         // 使用服务发现（如配置）
		grpc.WithTimeout(5*time.Second), // 超时控制
	)
	return &Data{
		c:    c,
		db:   entClient,
		grpc: v1.NewDirectSendServiceClient(conn),
	}, cleanup, nil
}
