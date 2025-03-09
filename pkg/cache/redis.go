package cache

import (
	"context"
	"fmt"
	"simple/model"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisClient Redis客户端接口
type RedisClient interface {
	// 基本操作
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, keys ...string) error
	Exists(ctx context.Context, key string) (bool, error)
	Expire(ctx context.Context, key string, expiration time.Duration) error

	// 列表操作
	LPush(ctx context.Context, key string, values ...interface{}) error
	RPush(ctx context.Context, key string, values ...interface{}) error
	LPop(ctx context.Context, key string) (string, error)
	RPop(ctx context.Context, key string) (string, error)
	LRange(ctx context.Context, key string, start, stop int64) ([]string, error)

	// 集合操作
	SAdd(ctx context.Context, key string, members ...interface{}) error
	SMembers(ctx context.Context, key string) ([]string, error)
	SRem(ctx context.Context, key string, members ...interface{}) error

	// 有序集合操作
	ZAdd(ctx context.Context, key string, score float64, member interface{}) error
	ZRange(ctx context.Context, key string, start, stop int64) ([]string, error)

	// 哈希表操作
	HSet(ctx context.Context, key string, field string, value interface{}) error
	HGet(ctx context.Context, key string, field string) (string, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HDel(ctx context.Context, key string, fields ...string) error

	// 发布订阅
	Subscribe(ctx context.Context, channels ...string) *redis.PubSub
	Publish(ctx context.Context, channel string, message interface{}) error

	// 关闭连接
	Close() error
}

// redisClient Redis客户端实现
type redisClient struct {
	client redis.UniversalClient
}

// NewRedisClient 创建Redis客户端
func NewRedisClient(config *model.RedisConfig) (RedisClient, error) {
	var client redis.UniversalClient

	opts := &redis.UniversalOptions{
		PoolSize:        config.Pool.MaxActive,
		MinIdleConns:    config.Pool.MaxIdle,
		ConnMaxIdleTime: config.Pool.IdleTimeout,

		DialTimeout:  config.Pool.ConnectTimeout,
		ReadTimeout:  config.Pool.ReadTimeout,
		WriteTimeout: config.Pool.WriteTimeout,
	}

	switch config.Mode {
	case "single":
		// 单机模式
		opts.Addrs = []string{
			fmt.Sprintf("%s:%d", config.Single.Host, config.Single.Port),
		}
		opts.DB = config.Single.DB
		opts.Password = config.Single.Password

	case "cluster":
		// 集群模式
		addrs := make([]string, len(config.Cluster.Nodes))
		for i, node := range config.Cluster.Nodes {
			addrs[i] = fmt.Sprintf("%s:%d", node.Host, node.Port)
		}
		opts.Addrs = addrs
		opts.Password = config.Cluster.Password
		opts.RouteRandomly = config.Cluster.EnableFollowRedirect

	case "sentinel":
		// 哨兵模式
		addrs := make([]string, len(config.Sentinel.Nodes))
		for i, node := range config.Sentinel.Nodes {
			addrs[i] = fmt.Sprintf("%s:%d", node.Host, node.Port)
		}
		opts.Addrs = addrs
		opts.MasterName = config.Sentinel.MasterName
		opts.Password = config.Sentinel.Password
		opts.DB = config.Sentinel.DB

	default:
		return nil, fmt.Errorf("unsupported redis mode: %s", config.Mode)
	}

	// 创建通用客户端
	client = redis.NewUniversalClient(opts)

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis connection test failed: %v", err)
	}

	return &redisClient{
		client: client,
	}, nil
}

// Set 设置键值对
func (r *redisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

// Get 获取值
func (r *redisClient) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

// Del 删除键
func (r *redisClient) Del(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}

// Exists 检查键是否存在
func (r *redisClient) Exists(ctx context.Context, key string) (bool, error) {
	n, err := r.client.Exists(ctx, key).Result()
	return n > 0, err
}

// Expire 设置过期时间
func (r *redisClient) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return r.client.Expire(ctx, key, expiration).Err()
}

// LPush 左推入列表
func (r *redisClient) LPush(ctx context.Context, key string, values ...interface{}) error {
	return r.client.LPush(ctx, key, values...).Err()
}

// RPush 右推入列表
func (r *redisClient) RPush(ctx context.Context, key string, values ...interface{}) error {
	return r.client.RPush(ctx, key, values...).Err()
}

// LPop 左弹出列表
func (r *redisClient) LPop(ctx context.Context, key string) (string, error) {
	return r.client.LPop(ctx, key).Result()
}

// RPop 右弹出列表
func (r *redisClient) RPop(ctx context.Context, key string) (string, error) {
	return r.client.RPop(ctx, key).Result()
}

// LRange 获取列表范围
func (r *redisClient) LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return r.client.LRange(ctx, key, start, stop).Result()
}

// SAdd 添加集合成员
func (r *redisClient) SAdd(ctx context.Context, key string, members ...interface{}) error {
	return r.client.SAdd(ctx, key, members...).Err()
}

// SMembers 获取集合所有成员
func (r *redisClient) SMembers(ctx context.Context, key string) ([]string, error) {
	return r.client.SMembers(ctx, key).Result()
}

// SRem 删除集合成员
func (r *redisClient) SRem(ctx context.Context, key string, members ...interface{}) error {
	return r.client.SRem(ctx, key, members...).Err()
}

// ZAdd 添加有序集合成员
func (r *redisClient) ZAdd(ctx context.Context, key string, score float64, member interface{}) error {
	return r.client.ZAdd(ctx, key, redis.Z{Score: score, Member: member}).Err()
}

// ZRange 获取有序集合范围
func (r *redisClient) ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return r.client.ZRange(ctx, key, start, stop).Result()
}

// HSet 设置哈希表字段
func (r *redisClient) HSet(ctx context.Context, key string, field string, value interface{}) error {
	return r.client.HSet(ctx, key, field, value).Err()
}

// HGet 获取哈希表字段
func (r *redisClient) HGet(ctx context.Context, key string, field string) (string, error) {
	return r.client.HGet(ctx, key, field).Result()
}

// HGetAll 获取哈希表所有字段
func (r *redisClient) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.client.HGetAll(ctx, key).Result()
}

// HDel 删除哈希表字段
func (r *redisClient) HDel(ctx context.Context, key string, fields ...string) error {
	return r.client.HDel(ctx, key, fields...).Err()
}

// Subscribe 订阅频道
func (r *redisClient) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
	return r.client.Subscribe(ctx, channels...)
}

// Publish 发布消息
func (r *redisClient) Publish(ctx context.Context, channel string, message interface{}) error {
	return r.client.Publish(ctx, channel, message).Err()
}

// Close 关闭连接
func (r *redisClient) Close() error {
	return r.client.Close()
}
