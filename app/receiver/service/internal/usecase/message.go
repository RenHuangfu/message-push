package usecase

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"message-push/app/receiver/common/model/entity"
	"message-push/app/receiver/common/repo"
)

type MessageUsecase struct {
	log  *log.Helper
	repo repo.MessageRepo
}

func NewMessageUsecase(logger log.Logger, repo repo.MessageRepo) *MessageUsecase {
	l := log.NewHelper(logger)
	return &MessageUsecase{
		log:  l,
		repo: repo,
	}
}

func (mu *MessageUsecase) SendMessage(ctx context.Context, p *entity.SendMessageParam) (*entity.SendMessageResult, error) {
	mu.log.Infof("send message, app_id: %d, client_ids: %v, content: %s", p.AppID, p.ClientIDs, p.Content)
	err := mu.repo.SaveMessageToDB(ctx, p)
	if err != nil {
		return nil, err
	}
	err = mu.repo.TriggerEvent(ctx, p)
	if err != nil {
		return nil, err
	}
	return &entity.SendMessageResult{}, nil
}
