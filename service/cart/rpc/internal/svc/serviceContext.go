package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"hanye/service/cart/model"
	"hanye/service/cart/rpc/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	CartModel   model.CartModel
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:      c,
		CartModel:   model.NewCartModel(sqlConn, c.Cache),
		RedisClient: redis.MustNewRedis(c.RedisConf),
	}
}
