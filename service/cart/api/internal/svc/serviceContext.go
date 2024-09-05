package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"hanye/middleware/jwt"
	"hanye/service/cart/api/internal/config"
	"hanye/service/cart/api/internal/middleware"
	"hanye/service/cart/rpc/cart"
	"hanye/service/meal/rpc/meal"
)

type ServiceContext struct {
	Config              config.Config
	UserAgentMiddleware rest.Middleware
	CartRpcClient       cart.CartZrpcClient
	MealRpcClient       meal.Meal
	JwtAuthMiddleware   rest.Middleware
	RedisClient         *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		UserAgentMiddleware: middleware.NewJwtAuthMiddleware().Handle,
		CartRpcClient:       cart.NewCartZrpcClient(zrpc.MustNewClient(c.CartRpcConf)),
		MealRpcClient:       meal.NewMeal(zrpc.MustNewClient(c.MealRpcConf)),
		RedisClient:         redis.MustNewRedis(c.RedisConf),
		JwtAuthMiddleware:   jwt.NewJwtAuthMiddleware().Handle,
	}
}
