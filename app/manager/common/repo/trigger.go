package repo

import (
	"context"
	"message-push/app/manager/common/model/entity"
)

type TriggerRepo interface {
	ProcessMessage(ctx context.Context, param *entity.TriggerParam) error
}
