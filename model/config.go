package model

import "time"

/*
   @NAME    : model
   @author  : 清风
   @desc    :
   @time    : 2025/3/6 22:39
*/

// Config 总配置结构
type Config struct {
	Server    ServerConfig    `yaml:"server" mapstructure:"server"`
	Database  DatabaseConfig  `yaml:"database" mapstructure:"database"`
	Redis     RedisConfig     `yaml:"redis" mapstructure:"redis"`
	JWT       JWTConfig       `yaml:"jwt" mapstructure:"jwt"`
	Telemetry TelemetryConfig `yaml:"telemetry" mapstructure:"telemetry"`
	Log       LogConfig       `yaml:"log" mapstructure:"log"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port         int           `yaml:"port" mapstructure:"port"`
	Mode         string        `yaml:"mode" mapstructure:"mode"`
	Static       string        `yaml:"static" mapstructure:"static"`
	ReadTimeout  time.Duration `yaml:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout" mapstructure:"write_timeout"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Write   DBConnConfig    `yaml:"write" mapstructure:"write"`
	Read    []DBConnConfig  `yaml:"read" mapstructure:"read"`
	Policy  DBPolicyConfig  `yaml:"policy" mapstructure:"policy"`
	Logger  DBLoggerConfig  `yaml:"logger" mapstructure:"logger"`
	Tracing DBTracingConfig `yaml:"tracing" mapstructure:"tracing"`
}

// DBConnConfig 数据库连接配置
type DBConnConfig struct {
	DSN             string        `yaml:"dsn" mapstructure:"dsn"`
	MaxIdleConns    int           `yaml:"max_idle_connections" mapstructure:"max_idle_connections"`
	MaxOpenConns    int           `yaml:"max_open_connections" mapstructure:"max_open_connections"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime" mapstructure:"conn_max_lifetime"`
}

// DBPolicyConfig 数据库策略配置
type DBPolicyConfig struct {
	Sources  []string `yaml:"sources" mapstructure:"sources"`
	Replicas []string `yaml:"replicas" mapstructure:"replicas"`
	Policy   string   `yaml:"policy" mapstructure:"policy"`
	Tables   []string `yaml:"tables" mapstructure:"tables"`
}

// DBLoggerConfig 数据库日志配置
type DBLoggerConfig struct {
	Level                string  `yaml:"level" mapstructure:"level"`
	SlowThreshold        float64 `yaml:"slow_threshold" mapstructure:"slow_threshold"`
	IgnoreRecordNotFound bool    `yaml:"ignore_record_not_found" mapstructure:"ignore_record_not_found"`
	Colorful             bool    `yaml:"colorful" mapstructure:"colorful"`
	LogFilePath          string  `yaml:"log_file_path" mapstructure:"log_file_path"`
	TraceFields          bool    `yaml:"trace_fields" mapstructure:"trace_fields"`
	ContextFields        bool    `yaml:"context_fields" mapstructure:"context_fields"`
}

// DBTracingConfig 数据库链路追踪配置
type DBTracingConfig struct {
	Enabled            bool   `yaml:"enabled" mapstructure:"enabled"`
	OperationPrefix    string `yaml:"operation_prefix" mapstructure:"operation_prefix"`
	RecordSQL          bool   `yaml:"record_sql" mapstructure:"record_sql"`
	RecordAffectedRows bool   `yaml:"record_affected_rows" mapstructure:"record_affected_rows"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Mode     string        `yaml:"mode" mapstructure:"mode"`
	Single   RedisSingle   `yaml:"single" mapstructure:"single"`
	Cluster  RedisCluster  `yaml:"cluster" mapstructure:"cluster"`
	Sentinel RedisSentinel `yaml:"sentinel" mapstructure:"sentinel"`
	Pool     RedisPool     `yaml:"pool" mapstructure:"pool"`
	Options  RedisOptions  `yaml:"options" mapstructure:"options"`
}

// RedisSingle 单机配置
type RedisSingle struct {
	Host     string `yaml:"host" mapstructure:"host"`
	Port     int    `yaml:"port" mapstructure:"port"`
	Password string `yaml:"password" mapstructure:"password"`
	DB       int    `yaml:"db" mapstructure:"db"`
}

// RedisCluster 集群配置
type RedisCluster struct {
	Nodes                []RedisNode `yaml:"nodes" mapstructure:"nodes"`
	Password             string      `yaml:"password" mapstructure:"password"`
	EnableFollowRedirect bool        `yaml:"enable_follow_redirect" mapstructure:"enable_follow_redirect"`
}

// RedisNode Redis节点配置
type RedisNode struct {
	Host string `yaml:"host" mapstructure:"host"`
	Port int    `yaml:"port" mapstructure:"port"`
}

// RedisSentinel 哨兵配置
type RedisSentinel struct {
	MasterName string      `yaml:"master_name" mapstructure:"master_name"`
	Nodes      []RedisNode `yaml:"nodes" mapstructure:"nodes"`
	Password   string      `yaml:"password" mapstructure:"password"`
	DB         int         `yaml:"db" mapstructure:"db"`
}

// RedisPool 连接池配置
type RedisPool struct {
	MaxIdle        int           `yaml:"max_idle" mapstructure:"max_idle"`
	MaxActive      int           `yaml:"max_active" mapstructure:"max_active"`
	IdleTimeout    time.Duration `yaml:"idle_timeout" mapstructure:"idle_timeout"`
	ConnectTimeout time.Duration `yaml:"connect_timeout" mapstructure:"connect_timeout"`
	ReadTimeout    time.Duration `yaml:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout   time.Duration `yaml:"write_timeout" mapstructure:"write_timeout"`
}

// RedisOptions Redis其他选项
type RedisOptions struct {
	Prefix            string `yaml:"prefix" mapstructure:"prefix"`
	EnableCompression bool   `yaml:"enable_compression" mapstructure:"enable_compression"`
	MinCompressLen    int    `yaml:"min_compress_len" mapstructure:"min_compress_len"`
	EnableTLS         bool   `yaml:"enable_tls" mapstructure:"enable_tls"`
	SkipVerify        bool   `yaml:"skip_verify" mapstructure:"skip_verify"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	SigningMethod string        `yaml:"signing_method" mapstructure:"signing_method"`
	SigningKey    string        `yaml:"signing_key" mapstructure:"signing_key"`
	Expiration    JWTExpiration `yaml:"expiration" mapstructure:"expiration"`
	Issuer        string        `yaml:"issuer" mapstructure:"issuer"`
	Subject       string        `yaml:"subject" mapstructure:"subject"`
	Audience      []string      `yaml:"audience" mapstructure:"audience"`
	Blacklist     JWTBlacklist  `yaml:"blacklist" mapstructure:"blacklist"`
	Options       JWTOptions    `yaml:"options" mapstructure:"options"`
	Refresh       JWTRefresh    `yaml:"refresh" mapstructure:"refresh"`
}

// JWTExpiration JWT过期时间配置
type JWTExpiration struct {
	AccessToken  time.Duration `yaml:"access_token" mapstructure:"access_token"`
	RefreshToken time.Duration `yaml:"refresh_token" mapstructure:"refresh_token"`
}

// JWTBlacklist JWT黑名单配置
type JWTBlacklist struct {
	Enabled     bool          `yaml:"enabled" mapstructure:"enabled"`
	Prefix      string        `yaml:"prefix" mapstructure:"prefix"`
	GracePeriod time.Duration `yaml:"grace_period" mapstructure:"grace_period"`
}

// JWTOptions JWT选项配置
type JWTOptions struct {
	VerifyExpiry    bool `yaml:"verify_expiry" mapstructure:"verify_expiry"`
	VerifyIssuedAt  bool `yaml:"verify_issued_at" mapstructure:"verify_issued_at"`
	VerifyIssuer    bool `yaml:"verify_issuer" mapstructure:"verify_issuer"`
	VerifySubject   bool `yaml:"verify_subject" mapstructure:"verify_subject"`
	VerifyAudience  bool `yaml:"verify_audience" mapstructure:"verify_audience"`
	VerifyNotBefore bool `yaml:"verify_not_before" mapstructure:"verify_not_before"`
}

// JWTRefresh JWT刷新配置
type JWTRefresh struct {
	AutoRefresh  bool          `yaml:"auto_refresh" mapstructure:"auto_refresh"`
	BeforeExpiry time.Duration `yaml:"before_expiry" mapstructure:"before_expiry"`
	Reuse        bool          `yaml:"reuse" mapstructure:"reuse"`
}

// TelemetryConfig 遥测配置
type TelemetryConfig struct {
	ServiceName  string        `yaml:"service_name" mapstructure:"service_name"`
	SamplingRate float64       `yaml:"sampling_rate" mapstructure:"sampling_rate"`
	OTLP         OTLPConfig    `yaml:"otlp" mapstructure:"otlp"`
	Trace        TraceConfig   `yaml:"trace" mapstructure:"trace"`
	Metrics      MetricsConfig `yaml:"metrics" mapstructure:"metrics"`
	Logs         LogsConfig    `yaml:"logs" mapstructure:"logs"`
}

// OTLPConfig OTLP配置
type OTLPConfig struct {
	Endpoint string        `yaml:"endpoint" mapstructure:"endpoint"`
	Insecure bool          `yaml:"insecure" mapstructure:"insecure"`
	Timeout  time.Duration `yaml:"timeout" mapstructure:"timeout"`
}

// TraceConfig 追踪配置
type TraceConfig struct {
	Enabled    bool                   `yaml:"enabled" mapstructure:"enabled"`
	Attributes map[string]interface{} `yaml:"attributes" mapstructure:"attributes"`
}

// MetricsConfig 指标配置
type MetricsConfig struct {
	Enabled  bool          `yaml:"enabled" mapstructure:"enabled"`
	Interval time.Duration `yaml:"interval" mapstructure:"interval"`
}

// LogsConfig 日志配置
type LogsConfig struct {
	Enabled bool   `yaml:"enabled" mapstructure:"enabled"`
	Level   string `yaml:"level" mapstructure:"level"`
}

// LogConfig Zap日志配置
type LogConfig struct {
	Level       string      `yaml:"level" mapstructure:"level"`
	Format      string      `yaml:"format" mapstructure:"format"`
	Output      LogOutput   `yaml:"output" mapstructure:"output"`
	Rotate      LogRotate   `yaml:"rotate" mapstructure:"rotate"`
	Caller      LogCaller   `yaml:"caller" mapstructure:"caller"`
	Development bool        `yaml:"development" mapstructure:"development"`
	Sampling    LogSampling `yaml:"sampling" mapstructure:"sampling"`
	Fields      LogFields   `yaml:"fields" mapstructure:"fields"`
}

// LogOutput 日志输出配置
type LogOutput struct {
	Console bool    `yaml:"console" mapstructure:"console"`
	File    LogFile `yaml:"file" mapstructure:"file"`
}

// LogFile 日志文件配置
type LogFile struct {
	Enabled bool   `yaml:"enabled" mapstructure:"enabled"`
	Path    string `yaml:"path" mapstructure:"path"`
}

// LogRotate 日志轮转配置
type LogRotate struct {
	Enabled    bool `yaml:"enabled" mapstructure:"enabled"`
	MaxSize    int  `yaml:"max_size" mapstructure:"max_size"`
	MaxBackups int  `yaml:"max_backups" mapstructure:"max_backups"`
	MaxAge     int  `yaml:"max_age" mapstructure:"max_age"`
	Compress   bool `yaml:"compress" mapstructure:"compress"`
}

// LogCaller 日志调用者配置
type LogCaller struct {
	Enabled bool `yaml:"enabled" mapstructure:"enabled"`
	Skip    int  `yaml:"skip" mapstructure:"skip"`
}

// LogSampling 日志采样配置
type LogSampling struct {
	Enabled    bool `yaml:"enabled" mapstructure:"enabled"`
	Initial    int  `yaml:"initial" mapstructure:"initial"`
	Thereafter int  `yaml:"thereafter" mapstructure:"thereafter"`
}

// LogFields 日志字段配置
type LogFields struct {
	Service string `yaml:"service" mapstructure:"service"`
	Env     string `yaml:"env" mapstructure:"env"`
}
