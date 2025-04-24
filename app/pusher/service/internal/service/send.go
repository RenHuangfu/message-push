package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"message-push/app/pusher/service/internal/conf"
)

type SendService struct {
	log *log.Helper
	c   *conf.Bootstrap
}
