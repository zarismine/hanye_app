// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	cartFieldNames          = builder.RawFieldNames(&Cart{})
	cartRows                = strings.Join(cartFieldNames, ",")
	cartRowsExpectAutoSet   = strings.Join(stringx.Remove(cartFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	cartRowsWithPlaceHolder = strings.Join(stringx.Remove(cartFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheHanyeCartIdPrefix = "cache:hanye:cart:id:"
	//cacheHanyeCartUserIdDishIdPrefix    = "cache:hanye:cart:userId:dishId:"
	//cacheHanyeCartUserIdSetmealIdPrefix = "cache:hanye:cart:userId:setmealId:"
	cacheHanyeCartUserIdPrefix = "cache:hanye:cart:userId:"
)

type (
	cartModel interface {
		Insert(ctx context.Context, data *Cart) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Cart, error)
		FindOneByUserIdDishIdFlavor(ctx context.Context, userId string, dishId int64, dishFlavor string) (*Cart, error)
		FindOneByUserIdSetmealIdFlavor(ctx context.Context, userId string, setmealId int64, dishFlavor string) (*Cart, error)
		FindByUserId(ctx context.Context, userId string) (*[]*Cart, error)
		DeleteByUserId(ctx context.Context, userId string) error
		Update(ctx context.Context, data *Cart) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCartModel struct {
		sqlc.CachedConn
		table string
	}

	Cart struct {
		Id         int64     `db:"id"`          // 主键
		Name       string    `db:"name"`        // 商品名称
		Pic        string    `db:"pic"`         // 图片
		UserId     string    `db:"user_id"`     // 用户id
		DishId     int64     `db:"dish_id"`     // 菜品id
		SetmealId  int64     `db:"setmeal_id"`  // 套餐id
		DishFlavor string    `db:"dish_flavor"` // 口味
		Number     int64     `db:"number"`      // 数量
		Amount     float64   `db:"amount"`      // 金额
		CreateTime time.Time `db:"create_time"` // 创建时间
	}
)

func newCartModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultCartModel {
	return &defaultCartModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`cart`",
	}
}

func (m *defaultCartModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	hanyeCartUserIdKey := fmt.Sprintf("%s%v", cacheHanyeCartUserIdPrefix, data.UserId)
	hanyeCartIdKey := fmt.Sprintf("%s%v", cacheHanyeCartIdPrefix, id)
	//hanyeCartUserIdDishIdKey := fmt.Sprintf("%s%v:%v", cacheHanyeCartUserIdDishIdPrefix, data.UserId, data.DishId)
	//hanyeCartUserIdSetmealIdKey := fmt.Sprintf("%s%v:%v", cacheHanyeCartUserIdSetmealIdPrefix, data.UserId, data.SetmealId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, hanyeCartIdKey)
	m.DelCacheCtx(ctx, hanyeCartUserIdKey)
	return err
}

func (m *defaultCartModel) FindOne(ctx context.Context, id int64) (*Cart, error) {
	hanyeCartIdKey := fmt.Sprintf("%s%v", cacheHanyeCartIdPrefix, id)
	var resp Cart
	err := m.QueryRowCtx(ctx, &resp, hanyeCartIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", cartRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCartModel) FindOneByUserIdDishIdFlavor(ctx context.Context, userId string, dishId int64, dishFlavor string) (*Cart, error) {
	var resp Cart
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `dish_id` = ? and `dish_flavor` = ? limit 1", cartRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, userId, dishId, dishFlavor)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCartModel) FindOneByUserIdSetmealIdFlavor(ctx context.Context, userId string, setmealId int64, dishFlavor string) (*Cart, error) {
	var resp Cart
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `setmeal_id` = ? and `dish_flavor` = ? limit 1", cartRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, userId, setmealId, dishFlavor)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCartModel) FindByUserId(ctx context.Context, userId string) (*[]*Cart, error) {
	hanyeCartUserIdKey := fmt.Sprintf("%s%v", cacheHanyeCartUserIdPrefix, userId)
	var resp []*Cart
	m.GetCacheCtx(ctx, hanyeCartUserIdKey, &resp)
	if len(resp) != 0 {
		return &resp, nil
	}
	query := fmt.Sprintf("select * from %s where `user_id` = ?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		m.SetCacheCtx(ctx, hanyeCartUserIdKey, resp)
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCartModel) DeleteByUserId(ctx context.Context, userId string) error {
	carts, err := m.FindByUserId(ctx, userId)
	if err != nil {
		return err
	}
	for _, v := range *carts {
		err = m.Delete(ctx, v.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *defaultCartModel) Insert(ctx context.Context, data *Cart) (sql.Result, error) {
	hanyeCartUserIdKey := fmt.Sprintf("%s%v", cacheHanyeCartUserIdPrefix, data.UserId)
	hanyeCartIdKey := fmt.Sprintf("%s%v", cacheHanyeCartIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, cartRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.Pic, data.UserId, data.DishId, data.SetmealId, data.DishFlavor, data.Number, data.Amount)
	}, hanyeCartIdKey)
	m.DelCacheCtx(ctx, hanyeCartUserIdKey)
	return ret, err
}

func (m *defaultCartModel) Update(ctx context.Context, newData *Cart) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}
	hanyeCartUserIdKey := fmt.Sprintf("%s%v", cacheHanyeCartUserIdPrefix, data.UserId)
	//var resp []*Cart
	hanyeCartIdKey := fmt.Sprintf("%s%v", cacheHanyeCartIdPrefix, data.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, cartRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Name, newData.Pic, newData.UserId, newData.DishId, newData.SetmealId, newData.DishFlavor, newData.Number, newData.Amount, newData.Id)
	}, hanyeCartIdKey)
	m.DelCacheCtx(ctx, hanyeCartUserIdKey)
	return err
}

func (m *defaultCartModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheHanyeCartIdPrefix, primary)
}

func (m *defaultCartModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", cartRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultCartModel) tableName() string {
	return m.table
}
