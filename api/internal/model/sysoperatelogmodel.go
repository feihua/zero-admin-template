package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysOperateLogModel = (*customSysOperateLogModel)(nil)

type (
	// SysOperateLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysOperateLogModel.
	SysOperateLogModel interface {
		sysOperateLogModel
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
