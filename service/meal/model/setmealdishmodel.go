package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SetmealDishModel = (*customSetmealDishModel)(nil)

type (
	// SetmealDishModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSetmealDishModel.
	SetmealDishModel interface {
		setmealDishModel
	}

	customSetmealDishModel struct {
		*defaultSetmealDishModel
	}
)

// NewSetmealDishModel returns a model for the database table.
func NewSetmealDishModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SetmealDishModel {
	return &customSetmealDishModel{
		defaultSetmealDishModel: newSetmealDishModel(conn, c, opts...),
	}
}
