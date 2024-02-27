package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysOperateLogModel = (*customSysOperateLogModel)(nil)

type (
	// SysOperateLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysOperateLogModel.
	SysOperateLogModel interface {
		sysOperateLogModel
		FindAll(ctx context.Context, Current int64, PageSize int64, UserName, Method, Ip string) (*[]SysOperateLog, int64, error)
	}

	customSysOperateLogModel struct {
		*defaultSysOperateLogModel
	}
)

// NewSysOperateLogModel returns a model for the database table.
func NewSysOperateLogModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysOperateLogModel {
	return &customSysOperateLogModel{
		defaultSysOperateLogModel: newSysOperateLogModel(conn, c, opts...),
	}
}
func (m *defaultSysOperateLogModel) FindAll(ctx context.Context, Current int64, PageSize int64, UserName, Method, Ip string) (*[]SysOperateLog, int64, error) {

	where := "1=1"
	if len(UserName) > 0 {
		where = where + fmt.Sprintf(" AND user_name like '%%%s%%'", UserName)
	}
	if len(Method) > 0 {
		where = where + fmt.Sprintf(" AND method like '%%%s%%'", Method)
	}
	if len(Ip) > 0 {
		where = where + fmt.Sprintf(" AND ip like '%%%s%%'", Ip)
	}
	where = where + fmt.Sprint(" ORDER BY operation_time DESC")
	query := fmt.Sprintf("select %s from %s where %s limit ?,?", sysOperateLogRows, m.table, where)

	var resp []SysOperateLog
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
