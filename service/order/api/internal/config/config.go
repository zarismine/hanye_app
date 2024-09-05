package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	CartRpcConf    zrpc.RpcClientConf
	OrderRpcConf   zrpc.RpcClientConf
	AddressRpcConf zrpc.RpcClientConf
	UserRpcConf    zrpc.RpcClientConf
	RedisConf      redis.RedisConf
}
