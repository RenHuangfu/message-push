package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/kafka"
	"message-push/app/pusher/service/internal/conf"
	"message-push/app/pusher/service/internal/service"
)

func NewKafkaServer(c *conf.Bootstrap, _ log.Logger, svc *service.Demo) *kafka.Server {
	ctx := context.Background()

	srv := kafka.NewServer(
		kafka.WithAddress(c.Data.Kafka.Brokers),
		kafka.WithGlobalTracerProvider(),
		kafka.WithGlobalPropagator(),
		kafka.WithCodec("json"),
	)

	_ = kafka.RegisterSubscriber(srv, ctx,
		c.Data.Kafka.Topic, "", false,
		svc.Consume,
	)

	return srv
}
