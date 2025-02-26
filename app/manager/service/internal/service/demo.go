package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/tx7do/kratos-transport/broker"
	v1 "message-push/app/manager/service/api/server/v1"
	"message-push/app/manager/service/internal/conf"
	"time"
)

type Demo struct {
	log        *log.Helper
	c          *conf.Bootstrap
	userClient v1.DemoClient
}

func NewDemo(logger log.Logger, conf *conf.Bootstrap) *Demo {
	conn, err := transgrpc.DialInsecure(
		context.Background(),
		transgrpc.WithEndpoint(conf.Client.User.Endpoint),
		transgrpc.WithTimeout(time.Duration(conf.Client.User.Timeout)*time.Millisecond),
	)
	if err != nil {
		panic(err)
	}
	return &Demo{
		log:        log.NewHelper(logger),
		c:          conf,
		userClient: v1.NewDemoClient(conn),
	}
}

func (s *Demo) Consume(ctx context.Context, _ string, _ broker.Headers, msg *v1.DemoMessage) error {
	fmt.Println(msg)
	_, _ = s.userClient.SendMessage(ctx, &v1.SendMessageRequest{
		Message: msg.Message,
	})

	return nil
}
