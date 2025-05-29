package data

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"message-push/app/manager/common/model/entity"
	"message-push/app/manager/common/repo"
	v1 "message-push/app/manager/service/api/server/v1"
	"message-push/app/receiver/common/model/po/ent/message"
)

type Trigger struct {
	log          *log.Helper
	data         *Data
	loadBalancer repo.LoadBalancer
}

func NewTriggerRepo(logger log.Logger, data *Data, balancer repo.LoadBalancer) repo.TriggerRepo {
	return &Trigger{
		log:          log.NewHelper(logger),
		data:         data,
		loadBalancer: balancer,
	}
}

func (t *Trigger) ProcessMessage(ctx context.Context, param *entity.TriggerParam) error {
	t.log.Infof("trigger process message: %v", param)
	res, err := t.data.db.Message.
		Query().
		Where(
			message.AppID(int(param.AppID)),
			message.SendTimeGTE(param.SendTime),
			message.SendCountGT(0),
			message.IsDel(0),
		).
		All(ctx)
	if err != nil {
		t.log.WithContext(ctx).Infof("error: %s", err)
		return err
	}
	for _, v := range res {
		clientIds := make([]int32, len(v.ClientIds))
		for i, v := range v.ClientIds {
			clientIds[i] = int32(v)
		}
		_, _ = t.data.grpc.DirectSend(ctx, &v1.DirectSendRequest{
			AppId:    int32(v.AppID),
			ClientId: clientIds,
			Content:  v.Content,
		})

		t.log.Infof("message: %v", v)
		msgData, err := json.Marshal(v)
		if err != nil {
			t.log.WithContext(ctx).Infof("error: %s", err)
			return err
		}

		producer, err := t.loadBalancer.SelectProducer(ctx, v)
		if err != nil {
			t.log.WithContext(ctx).Infof("error: %s", err)
			return err
		}

		err = producer.SendMessage(msgData)
		if err != nil {
			t.log.WithContext(ctx).Infof("error: %s", err)
			return err
		}

		err = t.data.db.Message.
			Update().
			Where(
				message.ID(v.ID),
				message.SendCountGT(0),
			).
			SetIsDel(1).
			Exec(ctx)
		if err != nil {
			t.log.WithContext(ctx).Infof("error: %s", err)
			return err
		}
	}
	return nil
}
