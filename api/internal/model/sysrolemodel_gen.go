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
	sysRoleFieldNames          = builder.RawFieldNames(&SysRole{})
	sysRoleRows                = strings.Join(sysRoleFieldNames, ",")
	sysRoleRowsExpectAutoSet   = strings.Join(stringx.Remove(sysRoleFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	sysRoleRowsWithPlaceHolder = strings.Join(stringx.Remove(sysRoleFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheHertzdbSysRoleIdPrefix       = "cache:hertzdb:sysRole:id:"
	cacheHertzdbSysRoleRoleNamePrefix = "cache:hertzdb:sysRole:roleName:"
)

type (
	sysRoleModel interface {
		Insert(ctx context.Context, data *SysRole) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysRole, error)
		FindOneByRoleName(ctx context.Context, roleName string) (*SysRole, error)
		Update(ctx context.Context, data *SysRole) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysRoleModel struct {
		sqlc.CachedConn
		table string
	}

	SysRole struct {
		Id         int64        `db:"id"`          // 主键
		RoleName   string       `db:"role_name"`   // 名称
		StatusId   int64        `db:"status_id"`   // 状态(1:正常，0:禁用)
		Sort       int64        `db:"sort"`        // 排序
		Remark     string       `db:"remark"`      // 备注
		CreateTime time.Time    `db:"create_time"` // 创建时间
		UpdateTime sql.NullTime `db:"update_time"` // 修改时间
	}
)

func newSysRoleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultSysRoleModel {
	return &defaultSysRoleModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`sys_role`",
	}
}

func (m *defaultSysRoleModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	hertzdbSysRoleIdKey := fmt.Sprintf("%s%v", cacheHertzdbSysRoleIdPrefix, id)
	hertzdbSysRoleRoleNameKey := fmt.Sprintf("%s%v", cacheHertzdbSysRoleRoleNamePrefix, data.RoleName)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, hertzdbSysRoleIdKey, hertzdbSysRoleRoleNameKey)
	return err
}

func (m *defaultSysRoleModel) FindOne(ctx context.Context, id int64) (*SysRole, error) {
	hertzdbSysRoleIdKey := fmt.Sprintf("%s%v", cacheHertzdbSysRoleIdPrefix, id)
	var resp SysRole
	err := m.QueryRowCtx(ctx, &resp, hertzdbSysRoleIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysRoleRows, m.table)
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

func (m *defaultSysRoleModel) FindOneByRoleName(ctx context.Context, roleName string) (*SysRole, error) {
	hertzdbSysRoleRoleNameKey := fmt.Sprintf("%s%v", cacheHertzdbSysRoleRoleNamePrefix, roleName)
	var resp SysRole
	err := m.QueryRowIndexCtx(ctx, &resp, hertzdbSysRoleRoleNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `role_name` = ? limit 1", sysRoleRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, roleName); err != nil {
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

func (m *defaultSysRoleModel) Insert(ctx context.Context, data *SysRole) (sql.Result, error) {
	hertzdbSysRoleIdKey := fmt.Sprintf("%s%v", cacheHertzdbSysRoleIdPrefix, data.Id)
	hertzdbSysRoleRoleNameKey := fmt.Sprintf("%s%v", cacheHertzdbSysRoleRoleNamePrefix, data.RoleName)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, sysRoleRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.RoleName, data.StatusId, data.Sort, data.Remark)
	}, hertzdbSysRoleIdKey, hertzdbSysRoleRoleNameKey)
	return ret, err
}

func (m *defaultSysRoleModel) Update(ctx context.Context, newData *SysRole) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	hertzdbSysRoleIdKey := fmt.Sprintf("%s%v", cacheHertzdbSysRoleIdPrefix, data.Id)
	hertzdbSysRoleRoleNameKey := fmt.Sprintf("%s%v", cacheHertzdbSysRoleRoleNamePrefix, data.RoleName)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysRoleRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.RoleName, newData.StatusId, newData.Sort, newData.Remark, newData.Id)
	}, hertzdbSysRoleIdKey, hertzdbSysRoleRoleNameKey)
	return err
}

func (m *defaultSysRoleModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheHertzdbSysRoleIdPrefix, primary)
}

func (m *defaultSysRoleModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysRoleRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSysRoleModel) tableName() string {
	return m.table
}
