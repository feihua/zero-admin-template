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
	sysRoleMenuFieldNames          = builder.RawFieldNames(&SysRoleMenu{})
	sysRoleMenuRows                = strings.Join(sysRoleMenuFieldNames, ",")
	sysRoleMenuRowsExpectAutoSet   = strings.Join(stringx.Remove(sysRoleMenuFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	sysRoleMenuRowsWithPlaceHolder = strings.Join(stringx.Remove(sysRoleMenuFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheHertzdbSysRoleMenuIdPrefix = "cache:hertzdb:sysRoleMenu:id:"
)

type (
	sysRoleMenuModel interface {
		Insert(ctx context.Context, data *SysRoleMenu) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysRoleMenu, error)
		Update(ctx context.Context, data *SysRoleMenu) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysRoleMenuModel struct {
		sqlc.CachedConn
		table string
	}

	SysRoleMenu struct {
		Id         int64        `db:"id"`          // 主键
		RoleId     int64        `db:"role_id"`     // 角色ID
		MenuId     int64        `db:"menu_id"`     // 菜单ID
		CreateTime time.Time    `db:"create_time"` // 创建时间
		UpdateTime sql.NullTime `db:"update_time"` // 修改时间
	}
)

func newSysRoleMenuModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultSysRoleMenuModel {
	return &defaultSysRoleMenuModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`sys_role_menu`",
	}
}

func (m *defaultSysRoleMenuModel) Delete(ctx context.Context, id int64) error {
	hertzdbSysRoleMenuIdKey := fmt.Sprintf("%s%v", cacheHertzdbSysRoleMenuIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, hertzdbSysRoleMenuIdKey)
	return err
}

func (m *defaultSysRoleMenuModel) FindOne(ctx context.Context, id int64) (*SysRoleMenu, error) {
	hertzdbSysRoleMenuIdKey := fmt.Sprintf("%s%v", cacheHertzdbSysRoleMenuIdPrefix, id)
	var resp SysRoleMenu
	err := m.QueryRowCtx(ctx, &resp, hertzdbSysRoleMenuIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysRoleMenuRows, m.table)
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

func (m *defaultSysRoleMenuModel) Insert(ctx context.Context, data *SysRoleMenu) (sql.Result, error) {
	hertzdbSysRoleMenuIdKey := fmt.Sprintf("%s%v", cacheHertzdbSysRoleMenuIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, sysRoleMenuRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.RoleId, data.MenuId)
	}, hertzdbSysRoleMenuIdKey)
	return ret, err
}

func (m *defaultSysRoleMenuModel) Update(ctx context.Context, data *SysRoleMenu) error {
	hertzdbSysRoleMenuIdKey := fmt.Sprintf("%s%v", cacheHertzdbSysRoleMenuIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysRoleMenuRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.RoleId, data.MenuId, data.Id)
	}, hertzdbSysRoleMenuIdKey)
	return err
}

func (m *defaultSysRoleMenuModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheHertzdbSysRoleMenuIdPrefix, primary)
}

func (m *defaultSysRoleMenuModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysRoleMenuRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSysRoleMenuModel) tableName() string {
	return m.table
}
