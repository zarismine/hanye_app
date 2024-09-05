package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"hanye/middleware/jwt"
	"hanye/service/address/api/internal/config"
	"hanye/service/address/api/internal/middleware"
	"hanye/service/address/rpc/address"
)

type ServiceContext struct {
	Config              config.Config
	UserAgentMiddleware rest.Middleware
	AddressRpcClient    address.AddressZrpcClient
	JwtAuthMiddleware   rest.Middleware
	RedisClient         *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		UserAgentMiddleware: middleware.NewJwtAuthMiddleware().Handle,
		AddressRpcClient:    address.NewAddressZrpcClient(zrpc.MustNewClient(c.AddressRpcConf)),
		RedisClient:         redis.MustNewRedis(c.RedisConf),
		JwtAuthMiddleware:   jwt.NewJwtAuthMiddleware().Handle,
	}
}
