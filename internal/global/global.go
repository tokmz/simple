package global

import (
	"gorm.io/gorm"
	"simple/model"
	"simple/pkg/cache"
	"simple/pkg/config"
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
	DB *gorm.DB
	// Cache redis client
	Cache *cache.RedisClient
)
