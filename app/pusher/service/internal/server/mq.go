package server

import (
	"context"
	"github.com/tx7do/kratos-transport/transport/kafka"
	"message-push/app/pusher/service/internal/conf"
	"message-push/app/pusher/service/internal/service"
	"os"
)

func NewKafkaServer(c *conf.Bootstrap, svc *service.SendService) *kafka.Server {
	ctx := context.Background()

	srv := kafka.NewServer(
		kafka.WithAddress(c.Data.Kafka.Brokers),
		kafka.WithGlobalTracerProvider(),
		kafka.WithGlobalPropagator(),
		kafka.WithCodec("json"),
	)

	_ = kafka.RegisterSubscriber(srv, ctx,
		os.Getenv("TOPIC"), c.Data.Kafka.Group, false,
		svc.Consume,
	)

	return srv
}
