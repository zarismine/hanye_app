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
	dishFieldNames          = builder.RawFieldNames(&Dish{})
	dishRows                = strings.Join(dishFieldNames, ",")
	dishRowsExpectAutoSet   = strings.Join(stringx.Remove(dishFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	dishRowsWithPlaceHolder = strings.Join(stringx.Remove(dishFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheHanyeDishIdPrefix   = "cache:hanye:dish:id:"
	cacheHanyeDishNamePrefix = "cache:hanye:dish:name:"
	cacheHanyeDishPrefix     = "cache:hanye:dish:"
)

type (
	dishModel interface {
		Insert(ctx context.Context, data *Dish) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Dish, error)
		FindOneByName(ctx context.Context, name string) (*Dish, error)
		FindDishesByCategoryId(ctx context.Context,categoryId int64) (*[]*Dish, error)
		FindDishesBySetmealId(ctx context.Context,setmealId int64) (*[]*Dish, error)
		FindDetailAndPicById(ctx context.Context,id int64) (string, string, error)
		Update(ctx context.Context, data *Dish) error
		Delete(ctx context.Context, id int64) error
	}

	defaultDishModel struct {
		sqlc.CachedConn
		table string
	}

	Dish struct {
		Id         int64     `db:"id"`
		Name       string    `db:"name"`
		Pic        string    `db:"pic"`
		Detail     string    `db:"detail"`
		Price      float64   `db:"price"`
		Status     int64     `db:"status"`
		CategoryId int64     `db:"category_id"`
		CreateUser int64     `db:"create_user"`
		UpdateUser int64     `db:"update_user"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newDishModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultDishModel {
	return &defaultDishModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`dish`",
	}
}

func (m *defaultDishModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	hanyeDishIdKey := fmt.Sprintf("%s%v", cacheHanyeDishIdPrefix, id)
	hanyeDishNameKey := fmt.Sprintf("%s%v", cacheHanyeDishNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, hanyeDishIdKey, hanyeDishNameKey)
	return err
}

func (m *defaultDishModel) FindOne(ctx context.Context, id int64) (*Dish, error) {
	hanyeDishIdKey := fmt.Sprintf("%s%v", cacheHanyeDishIdPrefix, id)
	var resp Dish
	err := m.QueryRowCtx(ctx, &resp, hanyeDishIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", dishRows, m.table)
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

func (m *defaultDishModel) FindOneByName(ctx context.Context, name string) (*Dish, error) {
	hanyeDishNameKey := fmt.Sprintf("%s%v", cacheHanyeDishNamePrefix, name)
	var resp Dish
	err := m.QueryRowIndexCtx(ctx, &resp, hanyeDishNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", dishRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, name); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDishModel) FindDishesByCategoryId(ctx context.Context,categoryId int64) (*[]*Dish, error) {
	cacheHanyeDishCategoryId := fmt.Sprintf("%s%v", cacheHanyeDishPrefix, categoryId)
	var resp []*Dish
	m.GetCacheCtx(ctx, cacheHanyeDishCategoryId, &resp)
	if len(resp) != 0 {
		return &resp, nil
	}
	query := fmt.Sprintf("select * from %s where `category_id` = ?",  m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, categoryId)
	switch err {
	case nil:
		m.SetCacheCtx(ctx, cacheHanyeDishCategoryId, resp)
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDishModel) FindDishesBySetmealId(ctx context.Context,setmealId int64) (*[]*Dish, error) {
	var resp []*Dish
	query := fmt.Sprintf("select * from %s where `category_id` = ?",  m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, setmealId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDishModel) FindDetailAndPicById(ctx context.Context,id int64) (string, string, error) {
	dish, err := m.FindOne(ctx, id)
	if err != nil {
		return "", "", err
	}
	return dish.Detail, dish.Pic, nil
}

func (m *defaultDishModel) Insert(ctx context.Context, data *Dish) (sql.Result, error) {
	hanyeDishIdKey := fmt.Sprintf("%s%v", cacheHanyeDishIdPrefix, data.Id)
	hanyeDishNameKey := fmt.Sprintf("%s%v", cacheHanyeDishNamePrefix, data.Name)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, dishRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.Pic, data.Detail, data.Price, data.Status, data.CategoryId, data.CreateUser, data.UpdateUser)
	}, hanyeDishIdKey, hanyeDishNameKey)
	return ret, err
}

func (m *defaultDishModel) Update(ctx context.Context, newData *Dish) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	hanyeDishIdKey := fmt.Sprintf("%s%v", cacheHanyeDishIdPrefix, data.Id)
	hanyeDishNameKey := fmt.Sprintf("%s%v", cacheHanyeDishNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, dishRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Name, newData.Pic, newData.Detail, newData.Price, newData.Status, newData.CategoryId, newData.CreateUser, newData.UpdateUser, newData.Id)
	}, hanyeDishIdKey, hanyeDishNameKey)
	return err
}

func (m *defaultDishModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheHanyeDishIdPrefix, primary)
}

func (m *defaultDishModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", dishRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultDishModel) tableName() string {
	return m.table
}
