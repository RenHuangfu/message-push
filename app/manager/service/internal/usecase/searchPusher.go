package usecase

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"message-push/app/manager/common/model/entity"
	"message-push/app/manager/common/repo"
)

type SearchPusherUsecase struct {
	repo repo.LoadBalancer
	log  *log.Helper
}

func NewSearchPusherUsecase(repo repo.LoadBalancer, logger log.Logger) *SearchPusherUsecase {
	return &SearchPusherUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (s *SearchPusherUsecase) SearchPusher(ctx context.Context, p *entity.SearchPusherParam) (string, error) {
	s.log.Infof("SearchPusher: %v", p)
	return s.repo.SearchPusher(ctx, p)
}
