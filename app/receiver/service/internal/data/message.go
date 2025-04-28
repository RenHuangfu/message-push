package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"message-push/app/receiver/common/model/entity"
	"message-push/app/receiver/common/repo"
	v1 "message-push/app/receiver/service/api/server/v1"
)

type MessageData struct {
	log  *log.Helper
	data *Data
}

func NewMessageRepo(
	logger log.Logger,
	data *Data,
) repo.MessageRepo {
	return &MessageData{
		log:  log.NewHelper(logger),
		data: data,
	}
}

func (md *MessageData) SaveMessageToDB(ctx context.Context, p *entity.SendMessageParam) error {
	md.log.WithContext(ctx).Infof("receiver存数据到数据库：%v", p)
	clientIds := make([]int, len(p.ClientIDs))
	for i, v := range p.ClientIDs {
		clientIds[i] = int(v)
	}
	err := md.data.db.Message.
		Create().
		SetAppID(int(p.AppID)).
		SetClientIds(clientIds).
		SetContent(p.Content).
		SetDelay(int(p.Delay)).
		SetSendTime(p.SendTime).
		SetSendCount(int(p.SendCount)).
		SetStatus(0).
		SetIsDel(0).
		Exec(ctx)
	if err != nil {
		md.log.WithContext(ctx).Infof("error %s", err)
		return err
	}
	return nil
}

func (md *MessageData) TriggerEvent(ctx context.Context, p *entity.SendMessageParam) error {
	md.log.WithContext(ctx).Infof("receiver触发事件：%v", p)
	_, err := md.data.grpc.ProcessMessage(ctx, &v1.ProcessMessageRequest{
		AppId:     p.AppID,
		SendTime:  p.SendTime.UnixMilli(),
		SendCount: p.SendCount,
		Delay:     p.Delay,
	})
	if err != nil {
		md.log.WithContext(ctx).Infof("error: %s", err)
		return err
	}
	return nil
}
