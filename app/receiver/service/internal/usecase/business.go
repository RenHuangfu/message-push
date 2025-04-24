package usecase

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"message-push/app/receiver/common/model/entity"
	"message-push/app/receiver/common/repo"
)

type BusinessUsecase struct {
	log  *log.Helper
	repo repo.BusinessRepo
}

func NewBusinessUsecase(logger log.Logger, repo repo.BusinessRepo) *BusinessUsecase {
	l := log.NewHelper(logger)
	return &BusinessUsecase{
		log:  l,
		repo: repo,
	}
}

func (bu *BusinessUsecase) CreateApp(ctx context.Context, p *entity.CreateAPPParam) (*entity.CreateAPPResult, error) {
	bu.log.Infof("create app, name: %s, key: %s", p.AppName, p.AppKey)
	return bu.repo.CreateAPP(ctx, p)
}

func (bu *BusinessUsecase) GetAPP(ctx context.Context, p *entity.GetAPPParam) (*entity.GetAPPResult, error) {
	bu.log.Infof("get app, id: %d", p.AppID)
	return bu.repo.GetAPP(ctx, p)
}

func (bu *BusinessUsecase) CreateClient(ctx context.Context, p *entity.CreateClientParam) (*entity.CreateClientResult, error) {
	bu.log.Infof("create client, name: %s, key: %s, app_id: %d", p.ClientName, p.ClientKey, p.AppID)
	return bu.repo.CreateClient(ctx, p)
}

func (bu *BusinessUsecase) GetClient(ctx context.Context, p *entity.GetClientParam) (*entity.GetClientResult, error) {
	bu.log.Infof("get client, id: %d", p.ClientID)
	return bu.repo.GetClient(ctx, p)
}

func (bu *BusinessUsecase) GetClientList(ctx context.Context, p *entity.GetClientListParam) (*entity.GetClientListResult, error) {
	bu.log.Infof("get client list, app_id: %d", p.AppID)
	return bu.repo.GetClientList(ctx, p)
}
