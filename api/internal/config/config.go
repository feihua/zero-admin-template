package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	CacheRedis cache.CacheConf

	Mysql struct {
		Datasource string
	}

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	Redis struct {
		Address string
		Pass    string
	}
}
