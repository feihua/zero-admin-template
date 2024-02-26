package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysRoleMenuModel = (*customSysRoleMenuModel)(nil)

type (
	// SysRoleMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysRoleMenuModel.
	SysRoleMenuModel interface {
		sysRoleMenuModel
		FindAllByRoleId(ctx context.Context, roleId int64) ([]int64, error)
		DeleteByRoleId(ctx context.Context, id int64) error
	}

	customSysRoleMenuModel struct {
		*defaultSysRoleMenuModel
	}
)

// NewSysRoleMenuModel returns a model for the database table.
func NewSysRoleMenuModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysRoleMenuModel {
	return &customSysRoleMenuModel{
		defaultSysRoleMenuModel: newSysRoleMenuModel(conn, c, opts...),
	}
}

func (m *defaultSysRoleMenuModel) FindAllByRoleId(ctx context.Context, roleId int64) ([]int64, error) {

	query := fmt.Sprintf("select menu_id from %s where role_id = ?", m.table)
	var resp []int64
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, roleId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return resp, err
	}
}

func (m *defaultSysRoleMenuModel) DeleteByRoleId(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where role_id = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, id)
	return err
}
