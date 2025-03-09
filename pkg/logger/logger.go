package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"simple/model"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	// Log 全局日志实例
	Log *zap.Logger
	// Sugar 全局Sugar日志实例，支持printf风格的API
	Sugar *zap.SugaredLogger
)

// 初始化日志
func Init(config *model.LogConfig) error {
	var err error
	Log, err = NewLogger(config)
	if err != nil {
		return err
	}
	Sugar = Log.Sugar()
	return nil
}

// NewLogger 创建一个新的日志实例
func NewLogger(config *model.LogConfig) (*zap.Logger, error) {
	// 解析日志级别
	level, err := parseLevel(config.Level)
	if err != nil {
		return nil, err
	}

	// 创建Core
	var cores []zapcore.Core

	// 控制台输出 - 始终使用控制台格式
	if config.Output.Console {
		// 控制台编码器配置
		consoleEncoderConfig := zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}

		// 开发模式下使用彩色级别
		if config.Development {
			consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		}

		consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)
		consoleCore := zapcore.NewCore(
			consoleEncoder,
			zapcore.Lock(os.Stdout),
			level,
		)
		cores = append(cores, consoleCore)
	}

	// 文件输出 - 始终使用JSON格式
	if config.Output.File.Enabled && config.Output.File.Path != "" {
		// JSON编码器配置 - 不使用颜色编码
		jsonEncoderConfig := zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder, // 不使用彩色编码
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}

		jsonEncoder := zapcore.NewJSONEncoder(jsonEncoderConfig)
		writer := getLogWriter(config)
		fileCore := zapcore.NewCore(
			jsonEncoder,
			zapcore.AddSync(writer),
			level,
		)
		cores = append(cores, fileCore)
	}

	// 合并cores
	core := zapcore.NewTee(cores...)

	// 创建Logger
	var zapOptions []zap.Option

	// 添加调用者信息
	if config.Caller.Enabled {
		zapOptions = append(zapOptions, zap.AddCaller())
		zapOptions = append(zapOptions, zap.AddCallerSkip(config.Caller.Skip))
	}

	// 开发模式
	if config.Development {
		zapOptions = append(zapOptions, zap.Development())
	}

	// 添加堆栈跟踪
	zapOptions = append(zapOptions, zap.AddStacktrace(zap.ErrorLevel))

	// 添加自定义字段
	if config.Fields.Service != "" {
		zapOptions = append(zapOptions, zap.Fields(zap.String("service", config.Fields.Service)))
	}
	if config.Fields.Env != "" {
		zapOptions = append(zapOptions, zap.Fields(zap.String("env", config.Fields.Env)))
	}

	// 创建Logger
	logger := zap.New(core, zapOptions...)

	// 配置采样
	if config.Sampling.Enabled {
		logger = logger.WithOptions(
			zap.WrapCore(func(core zapcore.Core) zapcore.Core {
				return zapcore.NewSamplerWithOptions(
					core,
					time.Second,
					config.Sampling.Initial,
					config.Sampling.Thereafter,
				)
			}),
		)
	}

	return logger, nil
}

// 解析日志级别
func parseLevel(level string) (zapcore.Level, error) {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel, nil
	case "info":
		return zapcore.InfoLevel, nil
	case "warn":
		return zapcore.WarnLevel, nil
	case "error":
		return zapcore.ErrorLevel, nil
	case "dpanic":
		return zapcore.DPanicLevel, nil
	case "panic":
		return zapcore.PanicLevel, nil
	case "fatal":
		return zapcore.FatalLevel, nil
	default:
		return zapcore.InfoLevel, fmt.Errorf("未知的日志级别: %s", level)
	}
}

// 获取日志写入器（带日志分割）
func getLogWriter(config *model.LogConfig) *lumberjack.Logger {
	logDir := filepath.Dir(config.Output.File.Path)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		fmt.Printf("创建日志目录失败: %v\n", err)
	}

	// 默认值设置
	maxSize := 100   // 默认100MB
	maxBackups := 30 // 默认保留30个备份
	maxAge := 30     // 默认保留30天
	compress := true // 默认启用压缩

	// 应用配置的值
	if config.Rotate.Enabled {
		if config.Rotate.MaxSize > 0 {
			maxSize = config.Rotate.MaxSize
		}
		if config.Rotate.MaxBackups > 0 {
			maxBackups = config.Rotate.MaxBackups
		}
		if config.Rotate.MaxAge > 0 {
			maxAge = config.Rotate.MaxAge
		}
		compress = config.Rotate.Compress
	}

	return &lumberjack.Logger{
		Filename:   config.Output.File.Path,
		MaxSize:    maxSize,    // 单位：MB
		MaxBackups: maxBackups, // 最大保留备份数
		MaxAge:     maxAge,     // 最大保留天数
		Compress:   compress,   // 是否压缩
	}
}

// Debug 输出debug级别日志
func Debug(msg string, fields ...zap.Field) {
	Log.Debug(msg, fields...)
}

// Info 输出info级别日志
func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

// Warn 输出warn级别日志
func Warn(msg string, fields ...zap.Field) {
	Log.Warn(msg, fields...)
}

// Error 输出error级别日志
func Error(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}

// DPanic 输出dpanic级别日志
func DPanic(msg string, fields ...zap.Field) {
	Log.DPanic(msg, fields...)
}

// Panic 输出panic级别日志
func Panic(msg string, fields ...zap.Field) {
	Log.Panic(msg, fields...)
}

// Fatal 输出fatal级别日志
func Fatal(msg string, fields ...zap.Field) {
	Log.Fatal(msg, fields...)
}

// Debugf 带格式的debug日志
func Debugf(format string, args ...interface{}) {
	Sugar.Debugf(format, args...)
}

// Infof 带格式的info日志
func Infof(format string, args ...interface{}) {
	Sugar.Infof(format, args...)
}

// Warnf 带格式的warn日志
func Warnf(format string, args ...interface{}) {
	Sugar.Warnf(format, args...)
}

// Errorf 带格式的error日志
func Errorf(format string, args ...interface{}) {
	Sugar.Errorf(format, args...)
}

// DPanicf 带格式的dpanic日志
func DPanicf(format string, args ...interface{}) {
	Sugar.DPanicf(format, args...)
}

// Panicf 带格式的panic日志
func Panicf(format string, args ...interface{}) {
	Sugar.Panicf(format, args...)
}

// Fatalf 带格式的fatal日志
func Fatalf(format string, args ...interface{}) {
	Sugar.Fatalf(format, args...)
}

// DebugWithCtx 带上下文字段的debug日志
func DebugWithCtx(ctx map[string]interface{}, msg string, fields ...zap.Field) {
	contextFields := getContextFields(ctx)
	Log.Debug(msg, append(contextFields, fields...)...)
}

// InfoWithCtx 带上下文字段的info日志
func InfoWithCtx(ctx map[string]interface{}, msg string, fields ...zap.Field) {
	contextFields := getContextFields(ctx)
	Log.Info(msg, append(contextFields, fields...)...)
}

// WarnWithCtx 带上下文字段的warn日志
func WarnWithCtx(ctx map[string]interface{}, msg string, fields ...zap.Field) {
	contextFields := getContextFields(ctx)
	Log.Warn(msg, append(contextFields, fields...)...)
}

// ErrorWithCtx 带上下文字段的error日志
func ErrorWithCtx(ctx map[string]interface{}, msg string, fields ...zap.Field) {
	contextFields := getContextFields(ctx)
	Log.Error(msg, append(contextFields, fields...)...)
}

// 从上下文中获取zap字段
func getContextFields(ctx map[string]interface{}) []zap.Field {
	fields := make([]zap.Field, 0, len(ctx))
	for k, v := range ctx {
		fields = append(fields, zap.Any(k, v))
	}
	return fields
}

// Sync 刷新并关闭日志
func Sync() {
	// 忽略sync错误，在某些平台上标准输出的sync操作会返回错误
	_ = Log.Sync()
}
