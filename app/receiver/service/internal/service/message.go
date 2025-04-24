package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"message-push/app/receiver/common/model/entity"
	v1 "message-push/app/receiver/service/api/server/v1"
	"message-push/app/receiver/service/internal/usecase"
	"time"
)

type MessageService struct {
	v1.MessageServer
	log *log.Helper
	muc *usecase.MessageUsecase
}

func NewMessageService(logger log.Logger, messageUsecase *usecase.MessageUsecase) *MessageService {
	return &MessageService{
		log: log.NewHelper(logger),
		muc: messageUsecase,
	}
}

func (ms *MessageService) SendMessage(ctx context.Context, req *v1.SendMessageRequest) (*v1.SendMessageResponse, error) {
	ms.log.Infof("send message, app_id: %d, client_ids: %v, content: %s", req.AppId, req.ClientIds, req.Content)
	if req.SendCount <= 0 {
		ms.log.Infof("SendCount is less then 0")
		return nil, nil
	}
	result, err := ms.muc.SendMessage(ctx, &entity.SendMessageParam{
		AppID:     req.AppId,
		ClientIDs: req.ClientIds,
		Content:   req.Content,
		Delay:     req.Delay,
		SendTime:  time.UnixMilli(req.SendTime),
		SendCount: req.SendCount,
		IsDel:     0,
	})
	if err != nil {
		ms.log.Errorf("send message error: %v", err)
		return nil, err
	}
	return &v1.SendMessageResponse{
		MessageId: result.MessageID,
	}, nil
}
