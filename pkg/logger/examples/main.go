package main

import (
	"simple/model"
	"simple/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	// 配置日志
	config := &model.LogConfig{
		Level:  "debug",
		Format: "console", // console或json
		Output: model.LogOutput{
			Console: true,
			File: model.LogFile{
				Enabled: true,
				Path:    "./logs/app.log",
			},
		},
		Rotate: model.LogRotate{
			Enabled:    true,
			MaxSize:    10,   // 单位MB
			MaxBackups: 5,    // 保留5个备份
			MaxAge:     7,    // 保留7天
			Compress:   true, // 压缩旧日志
		},
		Caller: model.LogCaller{
			Enabled: true,
			Skip:    0,
		},
		Development: true,
		Fields: model.LogFields{
			Service: "simple-app",
			Env:     "dev",
		},
	}

	// 初始化日志
	if err := logger.Init(config); err != nil {
		panic(err)
	}
	defer logger.Sync()

	// 基本日志记录
	logger.Debug("这是一条调试日志")
	logger.Info("这是一条信息日志")
	logger.Warn("这是一条警告日志")
	logger.Error("这是一条错误日志", zap.Error(ErrExample()))

	// 格式化日志记录
	logger.Debugf("格式化日志: %s, %d", "测试", 123)
	logger.Infof("当前用户: %s", "admin")

	// 带上下文的日志
	ctx := map[string]interface{}{
		"request_id": "req-123456",
		"user_id":    10001,
		"ip":         "192.168.1.1",
	}
	logger.InfoWithCtx(ctx, "用户登录成功")
	logger.ErrorWithCtx(ctx, "操作失败", zap.String("action", "delete"))
}

// ErrExample 返回一个示例错误
func ErrExample() error {
	return &customError{message: "这是一个示例错误"}
}

// 自定义错误类型
type customError struct {
	message string
}

func (e *customError) Error() string {
	return e.message
}
