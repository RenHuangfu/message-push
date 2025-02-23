package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"message-push/app/access/common/model/entity"
	"message-push/app/access/common/repo"
)

type Demo struct {
	data *Data
	log  *log.Helper
}

func NewDemo(data *Data, logger log.Logger) repo.Demo {
	return &Demo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (d *Demo) CreateName(ctx context.Context, p *entity.CreateNameParam) (*entity.CreateNameResult, error) {
	d.log.Infof("create demo, name: %s", p.Name)
	err := d.data.db.Demo.Create().
		SetName(p.Name).
		Exec(ctx)
	if err != nil {
		d.log.Errorf("create demo failed, err: %v", err)
		return nil, err
	}
	return &entity.CreateNameResult{
		Message: "success",
	}, nil
}
