package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"hanye/service/order/model"
	"hanye/service/order/rpc/internal/config"
)

type ServiceContext struct {
	Config           config.Config
	OrderModel       model.OrdersModel
	OrderDetailModel model.OrderDetailModel
	RedisClient      *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:           c,
		OrderModel:       model.NewOrdersModel(sqlConn, c.Cache),
		OrderDetailModel: model.NewOrderDetailModel(sqlConn, c.Cache),
		RedisClient:      redis.MustNewRedis(c.RedisConf),
	}
}
