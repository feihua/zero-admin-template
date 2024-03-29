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
	sysLoginLogFieldNames          = builder.RawFieldNames(&SysLoginLog{})
	sysLoginLogRows                = strings.Join(sysLoginLogFieldNames, ",")
	sysLoginLogRowsExpectAutoSet   = strings.Join(stringx.Remove(sysLoginLogFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	sysLoginLogRowsWithPlaceHolder = strings.Join(stringx.Remove(sysLoginLogFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheHertzdbSysLoginLogIdPrefix = "cache:hertzdb:sysLoginLog:id:"
)

type (
	sysLoginLogModel interface {
		Insert(ctx context.Context, data *SysLoginLog) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysLoginLog, error)
		Update(ctx context.Context, data *SysLoginLog) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysLoginLogModel struct {
		sqlc.CachedConn
		table string
	}

	SysLoginLog struct {
		Id        int64     `db:"id"`         // 编号
		UserName  string    `db:"user_name"`  // 用户名
		Status    string    `db:"status"`     // 登录状态（online:在线，登录初始状态，方便统计在线人数；login:退出登录后将online置为login；logout:退出登录）
		Ip        string    `db:"ip"`         // IP地址
		LoginTime time.Time `db:"login_time"` // 登陆时间
	}
)

func newSysLoginLogModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultSysLoginLogModel {
	return &defaultSysLoginLogModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`sys_login_log`",
	}
}

func (m *defaultSysLoginLogModel) Delete(ctx context.Context, id int64) error {
	hertzdbSysLoginLogIdKey := fmt.Sprintf("%s%v", cacheHertzdbSysLoginLogIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, hertzdbSysLoginLogIdKey)
	return err
}

func (m *defaultSysLoginLogModel) FindOne(ctx context.Context, id int64) (*SysLoginLog, error) {
	hertzdbSysLoginLogIdKey := fmt.Sprintf("%s%v", cacheHertzdbSysLoginLogIdPrefix, id)
	var resp SysLoginLog
	err := m.QueryRowCtx(ctx, &resp, hertzdbSysLoginLogIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysLoginLogRows, m.table)
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

func (m *defaultSysLoginLogModel) Insert(ctx context.Context, data *SysLoginLog) (sql.Result, error) {
	hertzdbSysLoginLogIdKey := fmt.Sprintf("%s%v", cacheHertzdbSysLoginLogIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, sysLoginLogRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserName, data.Status, data.Ip, data.LoginTime)
	}, hertzdbSysLoginLogIdKey)
	return ret, err
}

func (m *defaultSysLoginLogModel) Update(ctx context.Context, data *SysLoginLog) error {
	hertzdbSysLoginLogIdKey := fmt.Sprintf("%s%v", cacheHertzdbSysLoginLogIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysLoginLogRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserName, data.Status, data.Ip, data.LoginTime, data.Id)
	}, hertzdbSysLoginLogIdKey)
	return err
}

func (m *defaultSysLoginLogModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheHertzdbSysLoginLogIdPrefix, primary)
}

func (m *defaultSysLoginLogModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysLoginLogRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSysLoginLogModel) tableName() string {
	return m.table
}
