package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DishModel = (*customDishModel)(nil)

type (
	// DishModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDishModel.
	DishModel interface {
		dishModel
	}

	customDishModel struct {
		*defaultDishModel
	}
)

// NewDishModel returns a model for the database table.
func NewDishModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DishModel {
	return &customDishModel{
		defaultDishModel: newDishModel(conn, c, opts...),
	}
}
