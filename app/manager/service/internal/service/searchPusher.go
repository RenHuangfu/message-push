package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"message-push/app/manager/common/model/entity"
	v1 "message-push/app/manager/service/api/server/v1"
	"message-push/app/manager/service/internal/usecase"
)

type SearchPusherService struct {
	v1.SearchPusherServer
	uc  *usecase.SearchPusherUsecase
	log *log.Helper
}

func NewSearchPusherService(logger log.Logger, uc *usecase.SearchPusherUsecase) *SearchPusherService {
	return &SearchPusherService{
		log: log.NewHelper(logger),
		uc:  uc,
	}
}

func (s *SearchPusherService) SearchPusher(ctx context.Context, req *v1.PushSearchRequest) (*v1.PushSearchResponse, error) {
	s.log.Infof("search pusher, app_id: %s, client_id: %s", req.AppId, req.ClientId)
	result, err := s.uc.SearchPusher(ctx, &entity.SearchPusherParam{
		AppId:    req.AppId,
		ClientId: req.ClientId,
	})
	if err != nil {
		s.log.Errorf("search pusher error: %v", err)
		return nil, err
	}
	return &v1.PushSearchResponse{
		Url: result,
	}, nil
}
