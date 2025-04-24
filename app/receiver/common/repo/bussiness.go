package repo

import (
	"context"
	"message-push/app/receiver/common/model/entity"
)

type BusinessRepo interface {
	CreateAPP(ctx context.Context, p *entity.CreateAPPParam) (*entity.CreateAPPResult, error)
	GetAPP(ctx context.Context, p *entity.GetAPPParam) (*entity.GetAPPResult, error)
	CreateClient(ctx context.Context, p *entity.CreateClientParam) (*entity.CreateClientResult, error)
	GetClient(ctx context.Context, p *entity.GetClientParam) (*entity.GetClientResult, error)
	GetClientList(ctx context.Context, p *entity.GetClientListParam) (*entity.GetClientListResult, error)
}
