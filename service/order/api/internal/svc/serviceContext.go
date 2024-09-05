package svc

import (
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"hanye/middleware/jwt"
	"hanye/service/address/rpc/address"
	"hanye/service/cart/rpc/cart"
	"hanye/service/order/api/internal/config"
	"hanye/service/order/api/internal/middleware"
	"hanye/service/order/rpc/order"
	"hanye/service/user/rpc/user"
	"sync"
)

type Node struct {
	UserId        string
	Conn          *websocket.Conn //连接
	Addr          string          //客户端地址
	HeartbeatTime int64           //心跳时间
	LoginTime     int64           //登录时间
	DataQueue     chan []byte     //消息
}

type ServiceContext struct {
	sync.RWMutex
	Config              config.Config
	UserAgentMiddleware rest.Middleware
	CartRpcClient       cart.CartZrpcClient
	OrderRpcClient      order.OrderZrpcClient
	AddressRpcClient    address.AddressZrpcClient
	UserRpcClient       user.UserZrpcClient
	JwtAuthMiddleware   rest.Middleware
	RedisClient         *redis.Redis
	ClientMap           map[string]*Node
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		UserAgentMiddleware: middleware.NewJwtAuthMiddleware().Handle,
		CartRpcClient:       cart.NewCartZrpcClient(zrpc.MustNewClient(c.CartRpcConf)),
		OrderRpcClient:      order.NewOrderZrpcClient(zrpc.MustNewClient(c.OrderRpcConf)),
		AddressRpcClient:    address.NewAddressZrpcClient(zrpc.MustNewClient(c.AddressRpcConf)),
		UserRpcClient:       user.NewUserZrpcClient(zrpc.MustNewClient(c.UserRpcConf)),
		RedisClient:         redis.MustNewRedis(c.RedisConf),
		JwtAuthMiddleware:   jwt.NewJwtAuthMiddleware().Handle,
		ClientMap:           make(map[string]*Node),
	}
}

func (l *ServiceContext) SendMsg(data []byte) {
	for _, v := range l.ClientMap {
		v.DataQueue <- data
	}
}
