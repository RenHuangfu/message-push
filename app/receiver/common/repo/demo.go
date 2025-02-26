package repo

import (
	"context"
	"message-push/app/receiver/common/model/entity"
)

type Demo interface {
	CreateName(ctx context.Context, p *entity.CreateNameParam) (*entity.CreateNameResult, error)
	SearchNameWithCache(ctx context.Context, p *entity.SearchNameParam) (*entity.SearchNameResult, error)
	SendMessage(ctx context.Context, p *entity.SearchNameParam) (*entity.SearchNameResult, error)
}
