package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"hanye/service/meal/model"
	"hanye/service/meal/rpc/internal/config"
)

type ServiceContext struct {
	Config           config.Config
	CategoryModel    model.CategoryModel
	SetmealModel     model.SetmealModel
	DishModel        model.DishModel
	DishFlavorModel  model.DishFlavorModel
	SetmealDishModel model.SetmealDishModel
	RedisClient      *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:           c,
		CategoryModel:    model.NewCategoryModel(sqlConn, c.Cache),
		SetmealModel:     model.NewSetmealModel(sqlConn, c.Cache),
		DishModel:        model.NewDishModel(sqlConn, c.Cache),
		DishFlavorModel:  model.NewDishFlavorModel(sqlConn, c.Cache),
		SetmealDishModel: model.NewSetmealDishModel(sqlConn, c.Cache),
		RedisClient:      redis.MustNewRedis(c.RedisConf),
	}
}
