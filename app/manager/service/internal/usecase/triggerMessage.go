package usecase

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"message-push/app/manager/common/model/entity"
	"message-push/app/manager/common/repo"
)

type TriggerUsecase struct {
	log  *log.Helper
	repo repo.TriggerRepo
}

func NewTriggerUsecase(logger log.Logger, repo repo.TriggerRepo) *TriggerUsecase {
	l := log.NewHelper(logger)
	return &TriggerUsecase{
		log:  l,
		repo: repo,
	}
}

func (tu *TriggerUsecase) TriggerMessage(ctx context.Context, p *entity.TriggerParam) error {
	tu.log.Infof("trigger message, %v", p)
	err := tu.repo.ProcessMessage(ctx, p)
	if err != nil {
		return err
	}
	return nil
}
