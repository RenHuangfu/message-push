package data

import (
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"message-push/app/access/common/model/po/ent"
	"message-push/app/access/service/internal/conf"
)

type Data struct {
	c  *conf.Data
	db *ent.Client
}

func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	entClient, err := ent.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		return nil, nil, err
	}
	return &Data{
		c:  c,
		db: entClient,
	}, cleanup, nil
}
