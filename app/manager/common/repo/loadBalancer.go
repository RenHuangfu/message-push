package repo

import (
	"context"
	"github.com/RenHuangfu/tools/kafka"
	"message-push/app/manager/common/model/entity"
	"message-push/app/manager/common/model/po/ent"
)

type LoadBalancer interface {
	SelectProducer(ctx context.Context, p *ent.Message) (kafka.Producer, error)
	SearchPusher(ctx context.Context, p *entity.SearchPusherParam) (string, error)
}
