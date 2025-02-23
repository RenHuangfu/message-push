package usecase

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"message-push/app/access/common/model/entity"
	"message-push/app/access/common/repo"
)

type DemoUsecase struct {
	log  *log.Helper
	demo repo.Demo
}

func NewDemo(logger log.Logger, demo repo.Demo) *DemoUsecase {
	return &DemoUsecase{
		log:  log.NewHelper(logger),
		demo: demo,
	}
}

func (d *DemoUsecase) CreateName(ctx context.Context, p *entity.CreateNameParam) (*entity.CreateNameResult, error) {
	log.Info("usecase")
	return d.demo.CreateName(ctx, p)
}
