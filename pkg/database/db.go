package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"simple/model"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

// 链路追踪相关常量
const (
	tracerName = "simple/pkg/database"
	opCreate   = "gorm.Create"
	opQuery    = "gorm.Query"
	opUpdate   = "gorm.Update"
	opDelete   = "gorm.Delete"
	opRawSQL   = "gorm.RawSQL"
)

// 上下文键类型定义
type contextKey string

// 上下文键常量
const spanKey contextKey = "db_span"

// Init 初始化数据库连接
func Init(config *model.DatabaseConfig) (*gorm.DB, error) {
	if config.Write.DSN == "" {
		return nil, fmt.Errorf("主数据库DSN不能为空")
	}

	// 设置日志
	logConfig := logger.Config{
		SlowThreshold:             time.Duration(config.Logger.SlowThreshold * float64(time.Second)),
		LogLevel:                  parseLogLevel(config.Logger.Level),
		IgnoreRecordNotFoundError: config.Logger.IgnoreRecordNotFound,
		Colorful:                  config.Logger.Colorful,
	}

	// 配置日志输出目标
	var logWriter = os.Stdout
	if config.Logger.LogFilePath != "" {
		file, err := os.OpenFile(config.Logger.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			logWriter = file
		} else {
			log.Printf("无法打开日志文件 %s: %v, 使用标准输出代替", config.Logger.LogFilePath, err)
		}
	}

	// 创建日志记录器
	dbLogger := logger.New(log.New(logWriter, "\r\n", log.LstdFlags), logConfig)

	// 创建GORM配置
	gormConfig := &gorm.Config{
		Logger: dbLogger,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	}

	// 连接主数据库
	db, err := gorm.Open(mysql.Open(config.Write.DSN), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("连接主数据库失败: %w", err)
	}

	// 设置连接池参数
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("获取数据库连接失败: %w", err)
	}

	sqlDB.SetMaxIdleConns(config.Write.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.Write.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(config.Write.ConnMaxLifetime)

	// 配置读写分离
	if len(config.Read) > 0 {
		resolverConfig := dbresolver.Config{}

		// 配置主库
		resolverConfig.Sources = []gorm.Dialector{mysql.Open(config.Write.DSN)}

		// 配置从库
		var replicas []gorm.Dialector
		for _, read := range config.Read {
			replicas = append(replicas, mysql.Open(read.DSN))
		}
		resolverConfig.Replicas = replicas

		// 设置读写分离策略
		if config.Policy.Policy == "random" {
			resolverConfig.Policy = dbresolver.RandomPolicy{}
		}

		// 创建dbresolver并设置连接池参数
		dbResolverPlugin := dbresolver.Register(resolverConfig).
			SetMaxIdleConns(config.Write.MaxIdleConns).
			SetMaxOpenConns(config.Write.MaxOpenConns).
			SetConnMaxLifetime(config.Write.ConnMaxLifetime)

		// 应用读写分离
		if err := db.Use(dbResolverPlugin); err != nil {
			return nil, fmt.Errorf("配置读写分离失败: %w", err)
		}
	}

	// 如果启用了链路追踪
	if config.Tracing.Enabled {
		if err := db.Use(&TracingPlugin{
			tracer:             otel.Tracer(tracerName),
			operationPrefix:    config.Tracing.OperationPrefix,
			recordSQL:          config.Tracing.RecordSQL,
			recordAffectedRows: config.Tracing.RecordAffectedRows,
		}); err != nil {
			return nil, fmt.Errorf("配置链路追踪失败: %w", err)
		}
	}

	return db, nil
}

// 解析日志级别
func parseLogLevel(level string) logger.LogLevel {
	switch level {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Info
	}
}

// TracingPlugin 链路追踪插件
type TracingPlugin struct {
	tracer             trace.Tracer
	operationPrefix    string
	recordSQL          bool
	recordAffectedRows bool
}

// Name 返回插件名称
func (tp *TracingPlugin) Name() string {
	return "TracingPlugin"
}

// Initialize 初始化并添加回调
func (tp *TracingPlugin) Initialize(db *gorm.DB) error {
	// 为Create操作注册回调
	err := db.Callback().Create().Before("gorm:create").Register("tracing:before_create", tp.beforeCreate)
	if err != nil {
		return fmt.Errorf("注册Create前回调失败: %w", err)
	}
	err = db.Callback().Create().After("gorm:create").Register("tracing:after_create", tp.afterCreate)
	if err != nil {
		return fmt.Errorf("注册Create后回调失败: %w", err)
	}

	// 为Query操作注册回调
	err = db.Callback().Query().Before("gorm:query").Register("tracing:before_query", tp.beforeQuery)
	if err != nil {
		return fmt.Errorf("注册Query前回调失败: %w", err)
	}
	err = db.Callback().Query().After("gorm:query").Register("tracing:after_query", tp.afterQuery)
	if err != nil {
		return fmt.Errorf("注册Query后回调失败: %w", err)
	}

	// 为Update操作注册回调
	err = db.Callback().Update().Before("gorm:update").Register("tracing:before_update", tp.beforeUpdate)
	if err != nil {
		return fmt.Errorf("注册Update前回调失败: %w", err)
	}
	err = db.Callback().Update().After("gorm:update").Register("tracing:after_update", tp.afterUpdate)
	if err != nil {
		return fmt.Errorf("注册Update后回调失败: %w", err)
	}

	// 为Delete操作注册回调
	err = db.Callback().Delete().Before("gorm:delete").Register("tracing:before_delete", tp.beforeDelete)
	if err != nil {
		return fmt.Errorf("注册Delete前回调失败: %w", err)
	}
	err = db.Callback().Delete().After("gorm:delete").Register("tracing:after_delete", tp.afterDelete)
	if err != nil {
		return fmt.Errorf("注册Delete后回调失败: %w", err)
	}

	// 为Raw操作注册回调
	err = db.Callback().Raw().Before("gorm:raw").Register("tracing:before_raw", tp.beforeRaw)
	if err != nil {
		return fmt.Errorf("注册Raw前回调失败: %w", err)
	}
	err = db.Callback().Raw().After("gorm:raw").Register("tracing:after_raw", tp.afterRaw)
	if err != nil {
		return fmt.Errorf("注册Raw后回调失败: %w", err)
	}

	return nil
}

// startSpan 开始span
func (tp *TracingPlugin) startSpan(ctx context.Context, operation string, db *gorm.DB) (context.Context, trace.Span) {
	if ctx == nil {
		ctx = context.Background()
	}

	opName := operation
	if tp.operationPrefix != "" {
		opName = tp.operationPrefix + operation
	}

	opts := []trace.SpanStartOption{
		trace.WithAttributes(
			attribute.String("db.system", "mysql"),
			attribute.String("db.operation", operation),
		),
		trace.WithSpanKind(trace.SpanKindClient),
	}

	// 记录SQL语句
	if tp.recordSQL && db.Statement != nil && db.Statement.SQL.String() != "" {
		opts = append(opts, trace.WithAttributes(attribute.String("db.statement", db.Statement.SQL.String())))
	}

	// 记录表名
	if db.Statement != nil && db.Statement.Table != "" {
		opts = append(opts, trace.WithAttributes(attribute.String("db.table", db.Statement.Table)))
	}

	return tp.tracer.Start(ctx, opName, opts...)
}

// endSpan 结束span
func (tp *TracingPlugin) endSpan(span trace.Span, db *gorm.DB) {
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		span.SetStatus(codes.Error, db.Error.Error())
		span.RecordError(db.Error)
	} else {
		span.SetStatus(codes.Ok, "")
	}

	// 记录影响的行数
	if tp.recordAffectedRows && db.Statement != nil && db.Statement.SQL.String() != "" {
		span.SetAttributes(attribute.Int64("db.rows_affected", db.Statement.RowsAffected))
	}

	span.End()
}

// 各种操作回调
func (tp *TracingPlugin) beforeCreate(db *gorm.DB) {
	ctx, span := tp.startSpan(db.Statement.Context, opCreate, db)
	db.Statement.Context = ctx
	db.Statement.WithContext(ctx)
	db.Statement.Context = context.WithValue(db.Statement.Context, spanKey, span)
}

func (tp *TracingPlugin) afterCreate(db *gorm.DB) {
	span, ok := db.Statement.Context.Value(spanKey).(trace.Span)
	if ok {
		tp.endSpan(span, db)
	}
}

func (tp *TracingPlugin) beforeQuery(db *gorm.DB) {
	ctx, span := tp.startSpan(db.Statement.Context, opQuery, db)
	db.Statement.Context = ctx
	db.Statement.WithContext(ctx)
	db.Statement.Context = context.WithValue(db.Statement.Context, spanKey, span)
}

func (tp *TracingPlugin) afterQuery(db *gorm.DB) {
	span, ok := db.Statement.Context.Value(spanKey).(trace.Span)
	if ok {
		tp.endSpan(span, db)
	}
}

func (tp *TracingPlugin) beforeUpdate(db *gorm.DB) {
	ctx, span := tp.startSpan(db.Statement.Context, opUpdate, db)
	db.Statement.Context = ctx
	db.Statement.WithContext(ctx)
	db.Statement.Context = context.WithValue(db.Statement.Context, spanKey, span)
}

func (tp *TracingPlugin) afterUpdate(db *gorm.DB) {
	span, ok := db.Statement.Context.Value(spanKey).(trace.Span)
	if ok {
		tp.endSpan(span, db)
	}
}

func (tp *TracingPlugin) beforeDelete(db *gorm.DB) {
	ctx, span := tp.startSpan(db.Statement.Context, opDelete, db)
	db.Statement.Context = ctx
	db.Statement.WithContext(ctx)
	db.Statement.Context = context.WithValue(db.Statement.Context, spanKey, span)
}

func (tp *TracingPlugin) afterDelete(db *gorm.DB) {
	span, ok := db.Statement.Context.Value(spanKey).(trace.Span)
	if ok {
		tp.endSpan(span, db)
	}
}

func (tp *TracingPlugin) beforeRaw(db *gorm.DB) {
	ctx, span := tp.startSpan(db.Statement.Context, opRawSQL, db)
	db.Statement.Context = ctx
	db.Statement.WithContext(ctx)
	db.Statement.Context = context.WithValue(db.Statement.Context, spanKey, span)
}

func (tp *TracingPlugin) afterRaw(db *gorm.DB) {
	span, ok := db.Statement.Context.Value(spanKey).(trace.Span)
	if ok {
		tp.endSpan(span, db)
	}
}

// MasterDB 强制使用主库
func MasterDB(db *gorm.DB) *gorm.DB {
	return db.Clauses(dbresolver.Write)
}

// SlaveDB 强制使用从库
func SlaveDB(db *gorm.DB) *gorm.DB {
	return db.Clauses(dbresolver.Read)
}
