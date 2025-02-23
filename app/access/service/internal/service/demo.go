package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"message-push/app/access/common/model/entity"
	v1 "message-push/app/access/service/api/server/v1"
	"message-push/app/access/service/internal/usecase"
)

type Demo struct {
	log  *log.Helper
	demo *usecase.DemoUsecase
}

func NewDemo(logger log.Logger, demo *usecase.DemoUsecase) *Demo {
	return &Demo{
		log:  log.NewHelper(logger),
		demo: demo,
	}
}

func (s *Demo) DemoCreateName(ctx context.Context, in *v1.DemoCreateNameRequest) (*v1.DemoCreateNameResponse, error) {
	log.Info("CreateName")
	param := &entity.CreateNameParam{
		Name: in.Name,
	}
	result, err := s.demo.CreateName(ctx, param)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &v1.DemoCreateNameResponse{
		Message: result.Message,
	}, nil
}
