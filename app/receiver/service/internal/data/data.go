package data

import (
	"context"
	"github.com/RenHuangfu/tools/redis"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "github.com/go-sql-driver/mysql"
	"message-push/app/receiver/common/model/po/ent"
	v1 "message-push/app/receiver/service/api/server/v1"
	"message-push/app/receiver/service/internal/conf"
	"time"
)

type Data struct {
	c     *conf.Bootstrap
	db    *ent.Client
	redis *redis.Client
	grpc  v1.TriggerEventClient
}

func NewData(c *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	entClient, err := ent.Open(c.Data.Database.Driver, c.Data.Database.Source)
	if err != nil {
		return nil, nil, err
	}

	rdb := redis.MustNew(&redis.Option{
		Addr: c.Data.Redis.Addr,
	})

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
	if err != nil {
		panic(err)
	}

	return &Data{
		c:     c,
		db:    entClient,
		redis: rdb,
		grpc:  v1.NewTriggerEventClient(conn),
	}, cleanup, nil
}
