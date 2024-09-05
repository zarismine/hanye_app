package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"hanye/service/address/model"
	"hanye/service/address/rpc/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	AddressModel model.AddressBookModel
	RedisClient  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:       c,
		AddressModel: model.NewAddressBookModel(sqlConn, c.Cache),
		RedisClient:  redis.MustNewRedis(c.RedisConf),
	}
}
