package repo

import (
	"context"
	"message-push/app/access/common/model/entity"
)

type Demo interface {
	CreateName(ctx context.Context, p *entity.CreateNameParam) (*entity.CreateNameResult, error)
}
