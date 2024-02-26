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
	sysMenuFieldNames          = builder.RawFieldNames(&SysMenu{})
	sysMenuRows                = strings.Join(sysMenuFieldNames, ",")
	sysMenuRowsExpectAutoSet   = strings.Join(stringx.Remove(sysMenuFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	sysMenuRowsWithPlaceHolder = strings.Join(stringx.Remove(sysMenuFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheHertzdbSysMenuIdPrefix       = "cache:hertzdb:sysMenu:id:"
	cacheHertzdbSysMenuMenuNamePrefix = "cache:hertzdb:sysMenu:menuName:"
)

type (
	sysMenuModel interface {
		Insert(ctx context.Context, data *SysMenu) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysMenu, error)
		FindOneByMenuName(ctx context.Context, menuName string) (*SysMenu, error)
		Update(ctx context.Context, data *SysMenu) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysMenuModel struct {
		sqlc.CachedConn
		table string
	}

	SysMenu struct {
		Id         int64          `db:"id"`          // 主键
		MenuName   string         `db:"menu_name"`   // 菜单名称
		MenuType   int64          `db:"menu_type"`   // 菜单类型(1：目录   2：菜单   3：按钮)
		StatusId   int64          `db:"status_id"`   // 状态(1:正常，0:禁用)
		Sort       int64          `db:"sort"`        // 排序
		ParentId   int64          `db:"parent_id"`   // 父ID
		MenuUrl    string         `db:"menu_url"`    // 路由路径
		ApiUrl     string         `db:"api_url"`     // 接口URL
		MenuIcon   sql.NullString `db:"menu_icon"`   // 菜单图标
		Remark     sql.NullString `db:"remark"`      // 备注
		CreateTime time.Time      `db:"create_time"` // 创建时间
		UpdateTime sql.NullTime   `db:"update_time"` // 修改时间
	}
)

func newSysMenuModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultSysMenuModel {
	return &defaultSysMenuModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`sys_menu`",
	}
}

func (m *defaultSysMenuModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	hertzdbSysMenuIdKey := fmt.Sprintf("%s%v", cacheHertzdbSysMenuIdPrefix, id)
	hertzdbSysMenuMenuNameKey := fmt.Sprintf("%s%v", cacheHertzdbSysMenuMenuNamePrefix, data.MenuName)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, hertzdbSysMenuIdKey, hertzdbSysMenuMenuNameKey)
	return err
}

func (m *defaultSysMenuModel) FindOne(ctx context.Context, id int64) (*SysMenu, error) {
	hertzdbSysMenuIdKey := fmt.Sprintf("%s%v", cacheHertzdbSysMenuIdPrefix, id)
	var resp SysMenu
	err := m.QueryRowCtx(ctx, &resp, hertzdbSysMenuIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysMenuRows, m.table)
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

func (m *defaultSysMenuModel) FindOneByMenuName(ctx context.Context, menuName string) (*SysMenu, error) {
	hertzdbSysMenuMenuNameKey := fmt.Sprintf("%s%v", cacheHertzdbSysMenuMenuNamePrefix, menuName)
	var resp SysMenu
	err := m.QueryRowIndexCtx(ctx, &resp, hertzdbSysMenuMenuNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `menu_name` = ? limit 1", sysMenuRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, menuName); err != nil {
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

func (m *defaultSysMenuModel) Insert(ctx context.Context, data *SysMenu) (sql.Result, error) {
	hertzdbSysMenuIdKey := fmt.Sprintf("%s%v", cacheHertzdbSysMenuIdPrefix, data.Id)
	hertzdbSysMenuMenuNameKey := fmt.Sprintf("%s%v", cacheHertzdbSysMenuMenuNamePrefix, data.MenuName)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysMenuRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.MenuName, data.MenuType, data.StatusId, data.Sort, data.ParentId, data.MenuUrl, data.ApiUrl, data.MenuIcon, data.Remark)
	}, hertzdbSysMenuIdKey, hertzdbSysMenuMenuNameKey)
	return ret, err
}

func (m *defaultSysMenuModel) Update(ctx context.Context, newData *SysMenu) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	hertzdbSysMenuIdKey := fmt.Sprintf("%s%v", cacheHertzdbSysMenuIdPrefix, data.Id)
	hertzdbSysMenuMenuNameKey := fmt.Sprintf("%s%v", cacheHertzdbSysMenuMenuNamePrefix, data.MenuName)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysMenuRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.MenuName, newData.MenuType, newData.StatusId, newData.Sort, newData.ParentId, newData.MenuUrl, newData.ApiUrl, newData.MenuIcon, newData.Remark, newData.Id)
	}, hertzdbSysMenuIdKey, hertzdbSysMenuMenuNameKey)
	return err
}

func (m *defaultSysMenuModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheHertzdbSysMenuIdPrefix, primary)
}

func (m *defaultSysMenuModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysMenuRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSysMenuModel) tableName() string {
	return m.table
}
