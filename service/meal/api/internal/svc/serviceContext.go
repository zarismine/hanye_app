package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"hanye/service/meal/api/internal/config"
	"hanye/service/meal/rpc/meal"
)

type ServiceContext struct {
	Config        config.Config
	MealRpcClient meal.Meal
	RedisClient   *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		MealRpcClient: meal.NewMeal(zrpc.MustNewClient(c.MealRpcConf)),
		RedisClient:   redis.MustNewRedis(c.RedisConf),
	}
}
