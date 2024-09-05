package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SetmealModel = (*customSetmealModel)(nil)

type (
	// SetmealModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSetmealModel.
	SetmealModel interface {
		setmealModel
	}

	customSetmealModel struct {
		*defaultSetmealModel
	}
)

// NewSetmealModel returns a model for the database table.
func NewSetmealModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SetmealModel {
	return &customSetmealModel{
		defaultSetmealModel: newSetmealModel(conn, c, opts...),
	}
}
