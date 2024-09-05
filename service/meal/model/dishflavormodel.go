package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DishFlavorModel = (*customDishFlavorModel)(nil)

type (
	// DishFlavorModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDishFlavorModel.
	DishFlavorModel interface {
		dishFlavorModel
	}

	customDishFlavorModel struct {
		*defaultDishFlavorModel
	}
)

// NewDishFlavorModel returns a model for the database table.
func NewDishFlavorModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DishFlavorModel {
	return &customDishFlavorModel{
		defaultDishFlavorModel: newDishFlavorModel(conn, c, opts...),
	}
}
