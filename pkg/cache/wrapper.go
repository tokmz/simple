package cache

import (
	"simple/model"
)

var (
	// DefaultClient 默认Redis客户端
	DefaultClient RedisClient
)

// Setup 设置Redis客户端
func Setup(config *model.RedisConfig) error {
	client, err := NewRedisClient(config)
	if err != nil {
		return err
	}
	DefaultClient = client
	return nil
}

// Client 获取Redis客户端实例
func Client() RedisClient {
	if DefaultClient == nil {
		panic("redis client not initialized, call Setup first")
	}
	return DefaultClient
}

// Close 关闭Redis连接
func Close() error {
	if DefaultClient != nil {
		return DefaultClient.Close()
	}
	return nil
}
