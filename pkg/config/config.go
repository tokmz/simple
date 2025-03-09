package config

import (
	"errors"
	"fmt"
	"simple/model"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Manager 配置管理器
type Manager struct {
	mutex sync.RWMutex
	viper *viper.Viper
	path  string
}

// NewManager 创建配置管理器
func NewManager() *Manager {
	v := viper.New()

	// 设置配置文件搜索路径
	v.AddConfigPath("./resource/config")
	v.AddConfigPath("./config")
	// v.AddConfigPath("./configs")
	v.AddConfigPath(".")

	// 设置配置文件类型和名称
	v.SetConfigType("yaml")
	v.SetConfigName("config")

	// 大小写不敏感
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()

	// 其他常用配置
	v.AllowEmptyEnv(true)

	return &Manager{
		viper: v,
	}
}

var (
	ErrReadConfig = errors.New("读取配置文件失败")
	ErrUnmarshal  = errors.New("解析配置到结构体失败")
	ErrNotFound   = errors.New("配置项不存在")
)

// LoadFile 从文件加载配置
func (m *Manager) LoadFile(cfg *model.Config, path ...string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// 如果提供了路径，则使用指定路径
	if len(path) > 0 && path[0] != "" {
		m.path = path[0]
		m.viper.SetConfigFile(path[0])
	}

	// 读取配置文件
	if err := m.viper.ReadInConfig(); err != nil {
		return fmt.Errorf("%w: %s - %v", ErrReadConfig, m.viper.ConfigFileUsed(), err)
	}

	// 将配置解析到结构体
	if err := m.viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("%w: %v", ErrUnmarshal, err)
	}

	// 记录使用的配置文件路径
	if m.path == "" {
		m.path = m.viper.ConfigFileUsed()
	}

	// 监听配置文件变化
	m.viper.WatchConfig()
	m.viper.OnConfigChange(func(e fsnotify.Event) {
		if e.Op == fsnotify.Write || e.Op == fsnotify.Create {
			m.mutex.Lock()
			defer m.mutex.Unlock()

			fmt.Printf("配置文件发生变化: %s, 操作: %s\n", e.Name, e.Op.String())
			// 重新加载配置
			if err := m.viper.ReadInConfig(); err != nil {
				fmt.Printf("重新加载配置失败: %v\n", err)
				return
			}
			// 重新解析配置到结构体
			if err := m.viper.Unmarshal(cfg); err != nil {
				fmt.Printf("重新解析配置失败: %v\n", err)
			}
		}
	})

	return nil
}

// Get 获取配置
func (m *Manager) Get(key string) interface{} {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	if !m.viper.IsSet(key) {
		return nil
	}
	return m.viper.Get(key)
}

// GetString 获取字符串配置
func (m *Manager) GetString(key string) string {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.viper.GetString(key)
}

// GetStringSlice 获取字符串切片配置
func (m *Manager) GetStringSlice(key string) []string {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.viper.GetStringSlice(key)
}

// GetStringMap 获取字符串映射配置
func (m *Manager) GetStringMap(key string) map[string]interface{} {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.viper.GetStringMap(key)
}

// GetInt 获取整数配置
func (m *Manager) GetInt(key string) int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.viper.GetInt(key)
}

// GetBool 获取布尔配置
func (m *Manager) GetBool(key string) bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.viper.GetBool(key)
}

// GetDuration 获取时间间隔配置
func (m *Manager) GetDuration(key string) time.Duration {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.viper.GetDuration(key)
}

// UnmarshalKey 将指定key的值解析到结构体
func (m *Manager) UnmarshalKey(key string, val interface{}) error {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if !m.viper.IsSet(key) {
		return fmt.Errorf("%w: %s", ErrNotFound, key)
	}
	return m.viper.UnmarshalKey(key, val)
}

// Unmarshal 将配置解析到结构体
func (m *Manager) Unmarshal(val interface{}) error {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.viper.Unmarshal(val)
}

// Set 设置配置
func (m *Manager) Set(key string, value interface{}) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.viper.Set(key, value)
}

// WriteConfig 写入配置到文件
func (m *Manager) WriteConfig() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.viper.WriteConfig()
}

// SafeWriteConfig 安全写入配置到文件(不覆盖已存在的文件)
func (m *Manager) SafeWriteConfig() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.viper.SafeWriteConfig()
}

// LoadConfig 加载配置到model.Config结构体
func (m *Manager) LoadConfig() (*model.Config, error) {
	var cfg model.Config
	if err := m.LoadFile(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// GetServerConfig 获取服务器配置
func (m *Manager) GetServerConfig() (*model.ServerConfig, error) {
	var cfg model.ServerConfig
	if err := m.UnmarshalKey("server", &cfg); err != nil {
		return nil, fmt.Errorf("解析服务器配置失败: %w", err)
	}
	return &cfg, nil
}

// GetDatabaseConfig 获取数据库配置
func (m *Manager) GetDatabaseConfig() (*model.DatabaseConfig, error) {
	var cfg model.DatabaseConfig
	if err := m.UnmarshalKey("database", &cfg); err != nil {
		return nil, fmt.Errorf("解析数据库配置失败: %w", err)
	}
	return &cfg, nil
}

// GetRedisConfig 获取Redis配置
func (m *Manager) GetRedisConfig() (*model.RedisConfig, error) {
	var cfg model.RedisConfig
	if err := m.UnmarshalKey("redis", &cfg); err != nil {
		return nil, fmt.Errorf("解析Redis配置失败: %w", err)
	}
	return &cfg, nil
}

// GetJWTConfig 获取JWT配置
func (m *Manager) GetJWTConfig() (*model.JWTConfig, error) {
	var cfg model.JWTConfig
	if err := m.UnmarshalKey("jwt", &cfg); err != nil {
		return nil, fmt.Errorf("解析JWT配置失败: %w", err)
	}
	return &cfg, nil
}

// GetTelemetryConfig 获取链路追踪配置
func (m *Manager) GetTelemetryConfig() (*model.TelemetryConfig, error) {
	var cfg model.TelemetryConfig
	if err := m.UnmarshalKey("telemetry", &cfg); err != nil {
		return nil, fmt.Errorf("解析遥测配置失败: %w", err)
	}
	return &cfg, nil
}

// GetLogConfig 获取日志配置
func (m *Manager) GetLogConfig() (*model.LogConfig, error) {
	var cfg model.LogConfig
	if err := m.UnmarshalKey("log", &cfg); err != nil {
		return nil, fmt.Errorf("解析日志配置失败: %w", err)
	}
	return &cfg, nil
}

// GetFloat64 获取浮点数配置
func (m *Manager) GetFloat64(key string) float64 {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.viper.GetFloat64(key)
}

// GetIntSlice 获取整数切片配置
func (m *Manager) GetIntSlice(key string) []int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.viper.GetIntSlice(key)
}

// GetStringMapString 获取字符串映射字符串配置
func (m *Manager) GetStringMapString(key string) map[string]string {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.viper.GetStringMapString(key)
}

// GetStringMapStringSlice 获取字符串映射字符串切片配置
func (m *Manager) GetStringMapStringSlice(key string) map[string][]string {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.viper.GetStringMapStringSlice(key)
}

// IsSet 检查配置项是否存在
func (m *Manager) IsSet(key string) bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.viper.IsSet(key)
}
