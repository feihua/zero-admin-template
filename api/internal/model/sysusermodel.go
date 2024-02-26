package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysUserModel = (*customSysUserModel)(nil)

type (
	// SysUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserModel.
	SysUserModel interface {
		sysUserModel
		FindAll(ctx context.Context, Current int64, PageSize int64, mobile string, statusId int64) (*[]SysUser, int64, error)
		FindAllByUserId(ctx context.Context, userId int64) (*[]SysMenu, error)
	}

	customSysUserModel struct {
		*defaultSysUserModel
	}
)

// NewSysUserModel returns a model for the database table.
func NewSysUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysUserModel {
	return &customSysUserModel{
		defaultSysUserModel: newSysUserModel(conn, c, opts...),
	}
}

func (m *defaultSysUserModel) FindAll(ctx context.Context, Current int64, PageSize int64, mobile string, statusId int64) (*[]SysUser, int64, error) {

	where := "1=1"
	if len(mobile) > 0 {
		where = where + fmt.Sprintf(" AND user_name like '%%%s%%'", mobile)
	}
	if len(mobile) > 0 {
		where = where + fmt.Sprintf(" OR mobile like '%%%s%%'", mobile)
	}
	if statusId != 2 {
		where = where + fmt.Sprintf(" AND status_id = %d", statusId)
	}
	where = where + fmt.Sprint(" ORDER BY create_time DESC")
	query := fmt.Sprintf("select %s from %s where %s limit ?,?", sysUserRows, m.table, where)

	var resp []SysUser
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

func (m *defaultSysUserModel) FindAllByUserId(ctx context.Context, userId int64) (*[]SysMenu, error) {

	query := `select sm.*
from sys_user_role sur
         left join sys_role sr on sur.role_id = sr.id
         left join sys_role_menu srm on sr.id = srm.role_id
         left join sys_menu sm on srm.menu_id = sm.id
where sur.user_id = ?
order by sm.id`
	var resp []SysMenu
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
