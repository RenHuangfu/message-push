package usecase

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"message-push/app/receiver/common/model/entity"
	"message-push/app/receiver/common/repo"
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
	d.log.Infof("create demo, name: %s", p.Name)
	return d.demo.CreateName(ctx, p)
}

func (d *DemoUsecase) SearchName(ctx context.Context, p *entity.SearchNameParam) (*entity.SearchNameResult, error) {
	d.log.Infof("SearchName")
	return d.demo.SearchNameWithCache(ctx, p)
}
