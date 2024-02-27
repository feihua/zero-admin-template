package svc

import (
	"github.com/feihua/zero-admin-template/api/internal/config"
	"github.com/feihua/zero-admin-template/api/internal/middleware"
	"github.com/feihua/zero-admin-template/api/internal/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config

	UserModel       model.SysUserModel
	RoleModel       model.SysRoleModel
	UserRoleModel   model.SysUserRoleModel
	MenuModel       model.SysMenuModel
	RoleMenuModel   model.SysRoleMenuModel
	LoginLogModel   model.SysLoginLogModel
	OperateLogModel model.SysOperateLogModel
	Redis           *redis.Redis
	CheckUrl        rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	newRedis := redis.New(c.Redis.Address, redisConfig(c))
	return &ServiceContext{
		Config: c,

		UserModel:       model.NewSysUserModel(sqlConn, c.CacheRedis),
		RoleModel:       model.NewSysRoleModel(sqlConn, c.CacheRedis),
		UserRoleModel:   model.NewSysUserRoleModel(sqlConn, c.CacheRedis),
		MenuModel:       model.NewSysMenuModel(sqlConn, c.CacheRedis),
		RoleMenuModel:   model.NewSysRoleMenuModel(sqlConn, c.CacheRedis),
		LoginLogModel:   model.NewSysLoginLogModel(sqlConn, c.CacheRedis),
		OperateLogModel: model.NewSysOperateLogModel(sqlConn, c.CacheRedis),
		Redis:           newRedis,
		CheckUrl:        middleware.NewCheckUrlMiddleware(newRedis).Handle,
	}
}

func redisConfig(c config.Config) redis.Option {
	return func(r *redis.Redis) {
		r.Type = redis.NodeType
		r.Pass = c.Redis.Pass
	}
}
