package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AddressBookModel = (*customAddressBookModel)(nil)

type (
	// AddressBookModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAddressBookModel.
	AddressBookModel interface {
		addressBookModel
	}

	customAddressBookModel struct {
		*defaultAddressBookModel
	}
)

// NewAddressBookModel returns a model for the database table.
func NewAddressBookModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AddressBookModel {
	return &customAddressBookModel{
		defaultAddressBookModel: newAddressBookModel(conn, c, opts...),
	}
}
