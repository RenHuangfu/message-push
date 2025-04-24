package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"message-push/app/receiver/common/model/entity"
	v1 "message-push/app/receiver/service/api/server/v1"
	"message-push/app/receiver/service/internal/usecase"
)

type BusinessService struct {
	v1.BusinessServer
	log *log.Helper
	buc *usecase.BusinessUsecase
}

func NewBusinessService(logger log.Logger, businessUsecase *usecase.BusinessUsecase) *BusinessService {
	return &BusinessService{
		log: log.NewHelper(logger),
		buc: businessUsecase,
	}
}

func (bs *BusinessService) CreateApp(ctx context.Context, req *v1.CreateAPPRequest) (*v1.CreateAPPResponse, error) {
	bs.log.Infof("create app, name: %s, key: %s", req.AppName, req.AppKey)
	result, err := bs.buc.CreateApp(ctx, &entity.CreateAPPParam{
		AppName: req.AppName,
		AppKey:  req.AppKey,
	})
	if err != nil {
		bs.log.Errorf("create app error: %v", err)
		return nil, err
	}

	return &v1.CreateAPPResponse{
		AppId: result.AppID,
	}, nil
}

func (bs *BusinessService) GetAPP(ctx context.Context, req *v1.GetAPPRequest) (*v1.GetAPPResponse, error) {
	bs.log.Infof("get app, id: %d", req.AppId)
	result, err := bs.buc.GetAPP(ctx, &entity.GetAPPParam{
		AppID: req.AppId,
	})
	if err != nil {
		bs.log.Errorf("get app error: %v", err)
		return nil, err
	}

	return &v1.GetAPPResponse{
		AppId:   result.AppID,
		AppName: result.AppName,
	}, nil
}

func (bs *BusinessService) CreateClient(ctx context.Context, req *v1.CreateClientRequest) (*v1.CreateClientResponse, error) {
	bs.log.Infof("create client, name: %s, key: %s, app_id: %d", req.ClientName, req.ClientKey, req.AppId)
	result, err := bs.buc.CreateClient(ctx, &entity.CreateClientParam{
		ClientName: req.ClientName,
		ClientKey:  req.ClientKey,
		AppID:      req.AppId,
	})
	if err != nil {
		bs.log.Errorf("create client error: %v", err)
		return nil, err
	}

	return &v1.CreateClientResponse{
		ClientId: result.ClientID,
	}, nil
}

func (bs *BusinessService) GetClient(ctx context.Context, req *v1.GetClientRequest) (*v1.GetClientResponse, error) {
	bs.log.Infof("get client, id: %d", req.ClientId)
	result, err := bs.buc.GetClient(ctx, &entity.GetClientParam{
		ClientID: req.ClientId,
	})
	if err != nil {
		bs.log.Errorf("get client error: %v", err)
		return nil, err
	}

	return &v1.GetClientResponse{
		Client: &v1.ClientInfo{
			ClientId:   result.ClientID,
			ClientName: result.ClientName,
		},
	}, nil
}

func (bs *BusinessService) GetClientList(ctx context.Context, req *v1.GetClientListRequest) (*v1.GetClientListResponse, error) {
	bs.log.Infof("get client list, app_id: %d", req.AppId)
	result, err := bs.buc.GetClientList(ctx, &entity.GetClientListParam{
		AppID: req.AppId,
	})
	if err != nil {
		bs.log.Errorf("get client list error: %v", err)
		return nil, err
	}

	var clients []*v1.ClientInfo
	for _, client := range result.Clients {
		clients = append(clients, &v1.ClientInfo{
			ClientId:   client.ClientID,
			ClientName: client.ClientName,
		})
	}

	return &v1.GetClientListResponse{
		Clients: clients,
	}, nil
}
