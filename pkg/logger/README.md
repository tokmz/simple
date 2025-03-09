# 日志封装包

基于 [Zap](https://github.com/uber-go/zap) 和 [Lumberjack](https://github.com/natefinch/lumberjack) 的日志封装，提供了易用的 API 和日志分割功能。

## 功能特点

- 支持多种日志级别：Debug、Info、Warn、Error、DPanic、Panic、Fatal
- 支持多种输出格式：Console（控制台）和 JSON
- 支持多种输出目标：控制台和文件
- 支持日志文件分割（按大小分割）
- 支持日志文件保留策略（按时间和数量）
- 支持带格式化的日志记录方法
- 支持带上下文字段的日志记录
- 支持开发模式（彩色日志）
- 支持调用者信息记录
- 支持日志采样
- 支持自定义字段（服务名、环境等）

## 安装

```bash
go get go.uber.org/zap
go get gopkg.in/natefinch/lumberjack.v2
```

## 使用方法

### 1. 配置并初始化日志

```go
import (
    "simple/internal/model"
    "simple/pkg/logger"
)

// 配置日志
config := model.LogConfig{
    Level:  "debug",         // 日志级别：debug, info, warn, error, dpanic, panic, fatal
    Format: "console",       // 日志格式：console, json
    Output: model.LogOutput{
        Console: true,       // 是否输出到控制台
        File: model.LogFile{
            Enabled: true,   // 是否输出到文件
            Path:    "./logs/app.log", // 日志文件路径
        },
    },
    Rotate: model.LogRotate{
        Enabled:    true,    // 是否启用日志分割
        MaxSize:    10,      // 单个日志文件最大大小（MB）
        MaxBackups: 5,       // 最大保留日志文件数
        MaxAge:     7,       // 最大保留天数
        Compress:   true,    // 是否压缩旧日志
    },
    Caller: model.LogCaller{
        Enabled: true,       // 是否记录调用者信息
        Skip:    0,          // 调用层级跳过数
    },
    Development: true,       // 开发模式（启用彩色日志等）
    Fields: model.LogFields{
        Service: "my-service", // 服务名
        Env:     "dev",        // 环境名
    },
}

// 初始化日志
if err := logger.Init(config); err != nil {
    panic(err)
}
defer logger.Sync() // 确保所有日志都被刷新
```

### 2. 记录日志

```go
// 基本日志记录
logger.Debug("这是调试信息")
logger.Info("这是一般信息")
logger.Warn("这是警告信息")
logger.Error("这是错误信息", zap.String("error_code", "E001"))
logger.DPanic("这是开发模式下的panic信息")
logger.Panic("这会导致panic")
logger.Fatal("这会导致程序退出")

// 带格式化参数的日志记录
logger.Debugf("用户 %s 已登录，ID: %d", username, userId)
logger.Infof("处理了 %d 条记录", count)
logger.Warnf("操作耗时: %.2f 秒", elapsed)
logger.Errorf("请求失败: %v", err)

// 带上下文的日志记录
ctx := map[string]interface{}{
    "request_id": "req-123456",
    "user_id":    10001,
    "ip":         "192.168.1.1",
}
logger.InfoWithCtx(ctx, "用户登录成功")
logger.ErrorWithCtx(ctx, "操作失败", zap.String("action", "delete"))
```

### 3. 使用自定义字段

```go
// 使用 zap.Field 添加结构化字段
logger.Info("处理请求",
    zap.String("method", "GET"),
    zap.String("path", "/api/users"),
    zap.Int("status", 200),
    zap.Duration("elapsed", time.Millisecond*247),
)

// 记录错误
logger.Error("数据库连接失败",
    zap.Error(err),
    zap.String("db", "mysql"),
    zap.String("host", "localhost"),
)
```

## 示例代码

完整示例可以参考 [examples/main.go](examples/main.go)。
