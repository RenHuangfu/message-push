package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"message-push/app/manager/common/model/entity"
	v1 "message-push/app/manager/service/api/server/v1"
	"message-push/app/manager/service/internal/usecase"
	"time"
)

type TriggerService struct {
	v1.TriggerEventServer
	log *log.Helper
	tu  *usecase.TriggerUsecase
}

func NewTriggerService(logger log.Logger, tu *usecase.TriggerUsecase) *TriggerService {
	l := log.NewHelper(logger)
	return &TriggerService{
		log: l,
		tu:  tu,
	}
}

func (ts *TriggerService) ProcessMessage(ctx context.Context, req *v1.ProcessMessageRequest) (*v1.ProcessMessageResponse, error) {
	ts.log.Infof("trigger event, %v", req)
	err := ts.tu.TriggerMessage(ctx, &entity.TriggerParam{
		AppID:     req.AppId,
		SendTime:  time.UnixMilli(req.SendTime),
		SendCount: req.SendCount,
		Delay:     time.Duration(req.Delay) * time.Millisecond,
	})
	if err != nil {
		ts.log.Errorf("trigger event error: %v", err)
		return nil, err
	}
	return &v1.ProcessMessageResponse{}, nil
}
