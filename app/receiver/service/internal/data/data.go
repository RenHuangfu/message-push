package data

import (
	"github.com/RenHuangfu/tools/kafka"
	"github.com/RenHuangfu/tools/redis"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"message-push/app/receiver/common/model/po/ent"
	"message-push/app/receiver/service/internal/conf"
)

type Data struct {
	c        *conf.Bootstrap
	db       *ent.Client
	redis    *redis.Client
	producer kafka.Producer
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

	producer := kafka.NewKafkaProducer(
		kafka.WithBrokers(c.Data.Kafka.Brokers),
		kafka.WithTopic(c.Data.Kafka.Topic),
		kafka.WithAck(0),
		kafka.WithAsync())
	if producer == nil {
		panic("nil producer")
	}

	return &Data{
		c:        c,
		db:       entClient,
		redis:    rdb,
		producer: producer,
	}, cleanup, nil
}
