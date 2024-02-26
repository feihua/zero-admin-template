package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysUserRoleModel = (*customSysUserRoleModel)(nil)

type (
	// SysUserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserRoleModel.
	SysUserRoleModel interface {
		sysUserRoleModel
		FindAllByUserId(ctx context.Context, userId int64) (*SysUserRole, error)
		FindAllRoleIdsByUserId(ctx context.Context, userId int64) ([]int64, error)
		DeleteByUserId(ctx context.Context, id int64) error
	}

	customSysUserRoleModel struct {
		*defaultSysUserRoleModel
	}
)

// NewSysUserRoleModel returns a model for the database table.
func NewSysUserRoleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysUserRoleModel {
	return &customSysUserRoleModel{
		defaultSysUserRoleModel: newSysUserRoleModel(conn, c, opts...),
	}
}

func (m *defaultSysUserRoleModel) FindAllByUserId(ctx context.Context, userId int64) (*SysUserRole, error) {
	query := fmt.Sprintf("select %s from %s where user_id = ? and role_id = 1", sysUserRoleRows, m.table)

	var resp SysUserRole
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultSysUserRoleModel) FindAllRoleIdsByUserId(ctx context.Context, userId int64) ([]int64, error) {
	query := fmt.Sprintf("select role_id from %s where user_id = ?", m.table)

	var ids []int64
	err := m.QueryRowsNoCacheCtx(ctx, &ids, query, userId)
	switch err {
	case nil:
		return ids, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysUserRoleModel) DeleteByUserId(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where user_id = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, id)
	return err
}
