package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"message-push/app/receiver/common/model/entity"
	"message-push/app/receiver/common/model/po/ent/businessclient"
	"message-push/app/receiver/common/repo"
	"strconv"
)

type Business struct {
	log  *log.Helper
	data *Data
}

func NewBusinessRepo(data *Data, logger log.Logger) repo.BusinessRepo {
	return &Business{
		log:  log.NewHelper(logger),
		data: data,
	}
}

func (b *Business) CreateAPP(ctx context.Context, p *entity.CreateAPPParam) (*entity.CreateAPPResult, error) {
	b.log.Infof("create app, name: %s, key: %s", p.AppName, p.AppKey)
	b.data.db.BusinessApp.Create().
		SetName(p.AppName).
		SetAppKey(p.AppKey).
		SaveX(ctx)

	return &entity.CreateAPPResult{}, nil
}

func (b *Business) GetAPP(ctx context.Context, p *entity.GetAPPParam) (*entity.GetAPPResult, error) {
	b.log.Infof("get app, id: %d", p.AppID)
	app := b.data.db.BusinessApp.GetX(ctx, int(p.AppID))

	return &entity.GetAPPResult{
		AppID:   int32(app.ID),
		AppName: app.Name,
	}, nil
}

func (b *Business) CreateClient(ctx context.Context, p *entity.CreateClientParam) (*entity.CreateClientResult, error) {
	b.log.Infof("create client, name: %s, key: %s", p.ClientName, p.ClientKey)
	b.data.db.BusinessClient.Create().
		SetName(p.ClientName).
		SetClientKey(p.ClientKey).
		SetAppID(strconv.Itoa(int(p.AppID))).
		SaveX(ctx)

	return &entity.CreateClientResult{}, nil
}

func (b *Business) GetClient(ctx context.Context, p *entity.GetClientParam) (*entity.GetClientResult, error) {
	b.log.Infof("get client, id: %d", p.ClientID)
	client := b.data.db.BusinessClient.GetX(ctx, int(p.ClientID))

	return &entity.GetClientResult{
		ClientID:   int32(client.ID),
		ClientName: client.Name,
	}, nil
}

func (b *Business) GetClientList(ctx context.Context, p *entity.GetClientListParam) (*entity.GetClientListResult, error) {
	b.log.Infof("get client list, app_id: %d", p.AppID)
	clients := b.data.db.BusinessClient.Query().Where(businessclient.AppID(strconv.Itoa(int(p.AppID)))).AllX(ctx)
	result := &entity.GetClientListResult{
		Clients: make([]*entity.GetClientResult, 0),
	}
	for _, client := range clients {
		result.Clients = append(result.Clients, &entity.GetClientResult{
			ClientID:   int32(client.ID),
			ClientName: client.Name,
		})
	}

	return result, nil
}
