package repo

import (
	"context"
	"message-push/app/receiver/common/model/entity"
)

type MessageRepo interface {
	SaveMessageToDB(context.Context, *entity.SendMessageParam) error
	TriggerEvent(ctx context.Context, p *entity.SendMessageParam) error
}
