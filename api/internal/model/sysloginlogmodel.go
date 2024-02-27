package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysLoginLogModel = (*customSysLoginLogModel)(nil)

type (
	// SysLoginLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysLoginLogModel.
	SysLoginLogModel interface {
		sysLoginLogModel
		FindAll(ctx context.Context, Current int64, PageSize int64, userName, ip string) (*[]SysLoginLog, int64, error)
		StatisticsLoginLog(ctx context.Context, flag int64) (int32, error)
	}

	customSysLoginLogModel struct {
		*defaultSysLoginLogModel
	}
)

// NewSysLoginLogModel returns a model for the database table.
func NewSysLoginLogModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysLoginLogModel {
	return &customSysLoginLogModel{
		defaultSysLoginLogModel: newSysLoginLogModel(conn, c, opts...),
	}
}
func (m *defaultSysLoginLogModel) FindAll(ctx context.Context, Current int64, PageSize int64, userName, ip string) (*[]SysLoginLog, int64, error) {

	where := "1=1"
	if len(userName) > 0 {
		where = where + fmt.Sprintf(" AND user_name like '%%%s%%'", userName)
	}
	if len(ip) > 0 {
		where = where + fmt.Sprintf(" AND ip like '%%%s%%'", ip)
	}
	where = where + fmt.Sprint(" ORDER BY login_time DESC")
	query := fmt.Sprintf("select %s from %s where %s limit ?,?", sysLoginLogRows, m.table, where)

	var resp []SysLoginLog
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

func (m *customSysLoginLogModel) StatisticsLoginLog(ctx context.Context, flag int64) (int32, error) {
	query := ""
	if flag == 1 {
		query = "select count(distinct ip) current_day_login_count from sys_login_log where date(login_time) = curdate()"
	}
	if flag == 2 {
		query = "SELECT COUNT(DISTINCT ip) AS current_week_login_count FROM sys_login_log WHERE YEARWEEK(login_time, 1) = YEARWEEK(CURDATE(), 1)"
	}
	if flag == 3 {
		query = "SELECT COUNT(DISTINCT ip) AS current_month_login_count FROM sys_login_log WHERE MONTH(login_time) = MONTH(CURDATE())   AND YEAR(login_time) = YEAR(CURDATE())"
	}
	var count int32
	err := m.QueryRowNoCacheCtx(ctx, &count, query)

	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}
