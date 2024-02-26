package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysMenuModel = (*customSysMenuModel)(nil)

type (
	// SysMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysMenuModel.
	SysMenuModel interface {
		sysMenuModel
		FindAll(ctx context.Context, menuName string) (*[]SysMenu, error)
	}

	customSysMenuModel struct {
		*defaultSysMenuModel
	}
)

// NewSysMenuModel returns a model for the database table.
func NewSysMenuModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysMenuModel {
	return &customSysMenuModel{
		defaultSysMenuModel: newSysMenuModel(conn, c, opts...),
	}
}

func (m *defaultSysMenuModel) FindAll(ctx context.Context, menuName string) (*[]SysMenu, error) {

	query := fmt.Sprintf("select %s from %s order by id", sysMenuRows, m.table)
	var resp []SysMenu
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
