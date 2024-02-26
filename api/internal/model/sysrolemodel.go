package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysRoleModel = (*customSysRoleModel)(nil)

type (
	// SysRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysRoleModel.
	SysRoleModel interface {
		sysRoleModel
		FindAll(ctx context.Context, Current int64, PageSize int64, roleName string, statusId int64) (*[]SysRole, int64, error)
		QueryRoleList(ctx context.Context) (*[]SysRole, error)
	}

	customSysRoleModel struct {
		*defaultSysRoleModel
	}
)

// NewSysRoleModel returns a model for the database table.
func NewSysRoleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysRoleModel {
	return &customSysRoleModel{
		defaultSysRoleModel: newSysRoleModel(conn, c, opts...),
	}
}

func (m *defaultSysRoleModel) FindAll(ctx context.Context, Current int64, PageSize int64, roleName string, statusId int64) (*[]SysRole, int64, error) {

	where := "1=1"
	if len(roleName) > 0 {
		where = where + fmt.Sprintf(" AND role_name like '%%%s%%'", roleName)
	}
	if statusId != 2 {
		where = where + fmt.Sprintf(" AND status_id = %d", statusId)
	}
	where = where + fmt.Sprint(" ORDER BY create_time DESC")
	query := fmt.Sprintf("select %s from %s where %s limit ?,?", sysRoleRows, m.table, where)

	var resp []SysRole
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, (Current-1)*PageSize, PageSize)
	if err != nil {
		return nil, 0, err
	}
	query = fmt.Sprintf("select count(*) as count from %s where %s", m.table, where)

	var count int64
	err = m.QueryRowNoCacheCtx(ctx, &count, query)

	if err != nil {
		return nil, 0, err
	}

	return &resp, count, nil
}

func (m *defaultSysRoleModel) QueryRoleList(ctx context.Context) (*[]SysRole, error) {

	query := fmt.Sprintf("select %s from %s ORDER BY create_time DESC", sysRoleRows, m.table)

	var resp []SysRole
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
