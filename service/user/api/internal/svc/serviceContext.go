package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"hanye/middleware/jwt"
	"hanye/service/user/api/internal/config"
	"hanye/service/user/api/internal/middleware"
	"hanye/service/user/rpc/user"
)

type ServiceContext struct {
	Config                 config.Config
	UserAgentMiddleware    rest.Middleware
	UserRpcClient          user.UserZrpcClient
	JwtAuthMiddleware      rest.Middleware
	RedisClient            *redis.Redis
	LoginLogKqPusherClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                 c,
		UserAgentMiddleware:    middleware.NewJwtAuthMiddleware().Handle,
		UserRpcClient:          user.NewUserZrpcClient(zrpc.MustNewClient(c.UserRpcConf)),
		RedisClient:            redis.MustNewRedis(c.RedisConf),
		JwtAuthMiddleware:      jwt.NewJwtAuthMiddleware().Handle,
		LoginLogKqPusherClient: kq.NewPusher(c.LoginLogKqPusherConf.Brokers, c.LoginLogKqPusherConf.Topic),
	}
}
