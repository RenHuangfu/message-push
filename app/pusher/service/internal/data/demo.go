package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"message-push/app/receiver/common/model/entity"
	"message-push/app/receiver/common/model/po/ent"
	"message-push/app/receiver/common/model/po/ent/demo"
	"message-push/app/receiver/common/repo"
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

func (d *Demo) SearchNameWithCache(ctx context.Context, p *entity.SearchNameParam) (*entity.SearchNameResult, error) {
	d.log.Info("SearchNameWithCache")

	name, _ := d.data.redis.HGet(ctx, demo_name, fmt.Sprintf("%d", p.Id))
	if name != "" {
		return &entity.SearchNameResult{
			Name:  name,
			Exist: true,
		}, nil
	}

	result, err := d.data.db.Demo.
		Query().
		Where(demo.IDEQ(int(p.Id))).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return &entity.SearchNameResult{
				Exist: false,
			}, nil
		}
		return nil, err
	}

	err = d.data.redis.HSet(ctx, demo_name, p.Id, result.Name)
	if err != nil {
		return nil, err
	}

	return &entity.SearchNameResult{
		Name:  result.Name,
		Exist: true,
	}, nil
}

func (d *Demo) SendMessage(ctx context.Context, p *entity.SearchNameParam) (*entity.SearchNameResult, error) {
	res, err := d.SearchNameWithCache(ctx, p)
	if err != nil {
		return nil, err
	}

	if !res.Exist {
		return res, nil
	}

	err = d.data.producer.SendMessage([]byte(res.Name))
	if err != nil {
		return nil, err
	}

	return res, nil
}
