package global

import (
	"simple/internal/types/query"
	"simple/model"
	"simple/pkg/cache"
	"simple/pkg/config"

	"gorm.io/gorm"
)

/*
   @NAME    : global
   @author  : 清风
   @desc    :
   @time    : 2025/3/6 23:42
*/

var (
	// Cfg 配置信息
	Cfg *model.Config
	// Config Viper的管理
	Config *config.Manager
	// DB gorm 的db
	DB    *gorm.DB
	Query *query.Query
	// Cache redis client
	Cache *cache.RedisClient
)
