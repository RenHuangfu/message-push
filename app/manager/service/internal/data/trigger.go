package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"message-push/app/manager/common/model/entity"
	"message-push/app/manager/common/repo"
	"message-push/app/receiver/common/model/po/ent/message"
	"message-push/client/test"
)

type Trigger struct {
	log  *log.Helper
	data *Data
}

func NewTriggerRepo(logger log.Logger, data *Data) repo.TriggerRepo {
	return &Trigger{
		log:  log.NewHelper(logger),
		data: data,
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
		t.log.Infof("message: %v", v)
		msgData, err := json.Marshal(v)
		if err != nil {
			t.log.WithContext(ctx).Infof("error: %s", err)
			return err
		}
		err = t.data.producer.SendMessage(msgData)
		if err != nil {
			t.log.WithContext(ctx).Infof("error: %s", err)
			return err
		}
	}
	test.Test(fmt.Sprintf("manager读取数据库：%v", res))
	return nil
}
