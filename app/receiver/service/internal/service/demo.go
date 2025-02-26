package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"message-push/app/receiver/common/model/entity"
	v1 "message-push/app/receiver/service/api/server/v1"
	"message-push/app/receiver/service/internal/usecase"
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
	s.log.Infof("create demo, name: %v", in)

	param := &entity.CreateNameParam{
		Name: in.Name,
	}
	result, err := s.demo.CreateName(ctx, param)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}

	return &v1.DemoCreateNameResponse{
		Message: result.Message,
	}, nil
}

func (s *Demo) DemoSearchName(ctx context.Context, in *v1.DemoSearchNameRequest) (*v1.DemoSearchNameResponse, error) {
	s.log.Infof("search name %v", in)

	result, err := s.demo.SearchName(ctx, &entity.SearchNameParam{
		Id: in.Id,
	})
	if err != nil {
		return nil, err
	}
	return &v1.DemoSearchNameResponse{
		Exist: result.Exist,
		Name:  result.Name,
	}, nil
}

func (s *Demo) DemoMessage(ctx context.Context, in *v1.DemoMessageRequest) (*v1.DemoMessageResponse, error) {
	s.log.Infof("send message %v", in)
	result, err := s.demo.SendMessage(ctx, &entity.SearchNameParam{
		Id: in.Id,
	})
	if err != nil {
		return nil, err
	}
	return &v1.DemoMessageResponse{
		Message: result.Name,
	}, nil
}
