package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	CartRpcConf zrpc.RpcClientConf
	MealRpcConf zrpc.RpcClientConf
	RedisConf   redis.RedisConf
}
