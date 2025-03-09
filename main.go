package main

import (
	"fmt"
	"simple/internal/global"
	"simple/internal/types/query"
	"simple/model"
	"simple/pkg/cache"
	"simple/pkg/config"
	"simple/pkg/database"
	"simple/pkg/logger"

	"go.uber.org/zap"
)

/*
   @NAME    : main.go
   @author  : 清风
   @desc    :
   @time    : 2025/3/6 23:42
*/

func main() {
	global.Cfg = &model.Config{}

	var err error

	global.Config = config.NewManager()
	if global.Cfg, err = global.Config.LoadConfig(); err != nil {
		panic(err)
	} else {
		fmt.Printf("初始化配置成功\n")
	}

	if err = logger.Init(&global.Cfg.Log); err != nil {
		panic(err)
	} else {
		logger.Info("日志配置成功")
	}
	defer logger.Sync()

	if global.DB, err = database.Init(&global.Cfg.Database); err != nil {
		logger.Error("数据类连接失败", zap.Error(err))
		panic(err)
	} else {
		global.Query = query.Use(global.DB)
		logger.Info("数据类连接成功")
	}

	if err = cache.Setup(&global.Cfg.Redis); err != nil {
		logger.Error("redis 缓存连接失败", zap.Error(err))
		panic(err)
	} else {
		logger.Info("redis 缓存连接成功")
	}
	defer Close()
}

func Close() {
	if db, err := global.DB.DB(); err != nil {
		panic(err)
	} else {
		if err = db.Close(); err != nil {
			panic(err)
		} else {
			logger.Info("数据库连接关闭成功")
		}
	}

	if err := cache.Close(); err != nil {
		panic(err)
	} else {
		logger.Info("redis 缓存连接关闭成功")
	}
}
