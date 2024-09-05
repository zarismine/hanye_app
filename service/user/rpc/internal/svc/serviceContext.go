package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"hanye/service/user/model"
	"hanye/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	UserinfoModel model.UserModel
	//FollowsModel         model.FollowsModel
	RedisClient *redis.Redis
	//ContentRpcClient     content.Content
	DefaultPic string
	//DefaultAvatar        []string
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:        c,
		UserinfoModel: model.NewUserModel(sqlConn, c.Cache),
		//FollowsModel:         model.NewFollowsModel(sqlConn, c.Cache),
		RedisClient: redis.MustNewRedis(c.RedisConf),
		//ContentRpcClient:     content.NewContent(zrpc.MustNewClient(c.ContentRpcConf)),
		DefaultPic: c.DefaultPic,
	}
}
