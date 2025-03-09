package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"simple/model"
	"testing"
	"time"
)

// ExampleConfig 示例配置
func exampleConfig() *model.RedisConfig {
	return &model.RedisConfig{
		Mode: "single",
		Single: model.RedisSingle{
			Host:     "localhost",
			Port:     6379,
			Password: "",
			DB:       0,
		},
		Pool: model.RedisPool{
			MaxIdle:        10,
			MaxActive:      100,
			IdleTimeout:    time.Minute * 5,
			ConnectTimeout: time.Second * 5,
			ReadTimeout:    time.Second * 2,
			WriteTimeout:   time.Second * 2,
		},
	}
}

// Example 展示Redis基本操作示例
func Example() {
	// 初始化客户端
	client, err := NewRedisClient(exampleConfig())
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	// 设置字符串
	err = client.Set(ctx, "greeting", "Hello Redis!", time.Hour)
	if err != nil {
		panic(err)
	}

	// 获取字符串
	val, err := client.Get(ctx, "greeting")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get greeting: %s\n", val)

	// 检查键是否存在
	exists, err := client.Exists(ctx, "greeting")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Key exists: %v\n", exists)

	// Output:
	// Get greeting: Hello Redis!
	// Key exists: true
}

// ExampleListOperations 列表操作示例
func TestExampleListOperations(t *testing.T) {
	client, err := NewRedisClient(exampleConfig())
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	// 左推入列表
	err = client.LPush(ctx, "fruits", "apple", "banana", "orange")
	if err != nil {
		panic(err)
	}

	// 获取列表范围
	fruits, err := client.LRange(ctx, "fruits", 0, -1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Fruits list: %v\n", fruits)

	// 右弹出元素
	fruit, err := client.RPop(ctx, "fruits")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Popped fruit: %s\n", fruit)

	// Output:
	// Fruits list: [orange banana apple]
	// Popped fruit: apple
}

// User 用户结构体示例
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Example_HashOperations 哈希表操作示例
func TestExampleHashOperations(t *testing.T) {
	client, err := NewRedisClient(exampleConfig())
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	// 创建用户
	user := User{
		ID:       1,
		Username: "john_doe",
		Email:    "john@example.com",
	}

	// 将结构体转换为JSON
	userJSON, _ := json.Marshal(user)

	// 存储用户信息到哈希表
	err = client.HSet(ctx, "users:1", "data", string(userJSON))
	if err != nil {
		panic(err)
	}

	// 获取用户信息
	userStr, err := client.HGet(ctx, "users:1", "data")
	if err != nil {
		panic(err)
	}

	// 解析JSON
	var retrievedUser User
	_ = json.Unmarshal([]byte(userStr), &retrievedUser)

	fmt.Printf("Retrieved user: %+v\n", retrievedUser)

	// Output:
	// Retrieved user: {ID:1 Username:john_doe Email:john@example.com}
}

// TestExampleSetOperations 集合操作示例测试
func TestExampleSetOperations(t *testing.T) {
	client, err := NewRedisClient(exampleConfig())
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	// 添加标签到集合
	err = client.SAdd(ctx, "tags", "golang", "redis", "cache", "database")
	if err != nil {
		panic(err)
	}

	// 获取所有标签
	tags, err := client.SMembers(ctx, "tags")
	if err != nil {
		panic(err)
	}
	fmt.Printf("All tags: %v\n", tags)

	// 删除标签
	err = client.SRem(ctx, "tags", "database")
	if err != nil {
		panic(err)
	}

	// 再次获取标签
	tags, _ = client.SMembers(ctx, "tags")
	fmt.Printf("Tags after removal: %v\n", tags)

	// Output:
	// All tags: [golang redis cache database]
	// Tags after removal: [golang redis cache]
}

// ExamplePubSub 发布订阅示例
func ExamplePubSub() {
	client, err := NewRedisClient(exampleConfig())
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	// 订阅频道
	pubsub := client.Subscribe(ctx, "news")
	defer pubsub.Close()

	// 在goroutine中处理消息
	go func() {
		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				return
			}
			fmt.Printf("Received message: %s\n", msg.Payload)
		}
	}()

	// 发布消息
	err = client.Publish(ctx, "news", "Breaking news: Redis is awesome!")
	if err != nil {
		panic(err)
	}

	// 等待消息处理
	time.Sleep(time.Second)

	// Output:
	// Received message: Breaking news: Redis is awesome!
}

// TestRedisClient 测试Redis客户端
func TestRedisClient(t *testing.T) {
	// 跳过测试如果没有Redis服务器
	client, err := NewRedisClient(exampleConfig())
	if err != nil {
		t.Skip("Redis server is not available")
	}
	defer client.Close()

	ctx := context.Background()

	// 测试基本操作
	t.Run("Basic Operations", func(t *testing.T) {
		err := client.Set(ctx, "test_key", "test_value", time.Minute)
		if err != nil {
			t.Errorf("Set failed: %v", err)
		}

		val, err := client.Get(ctx, "test_key")
		if err != nil {
			t.Errorf("Get failed: %v", err)
		}
		if val != "test_value" {
			t.Errorf("Expected 'test_value', got '%s'", val)
		}
	})

	// 测试列表操作
	t.Run("List Operations", func(t *testing.T) {
		err := client.LPush(ctx, "test_list", "item1", "item2")
		if err != nil {
			t.Errorf("LPush failed: %v", err)
		}

		items, err := client.LRange(ctx, "test_list", 0, -1)
		if err != nil {
			t.Errorf("LRange failed: %v", err)
		}
		if len(items) != 2 {
			t.Errorf("Expected 2 items, got %d", len(items))
		}
	})

	// 清理测试数据
	client.Del(ctx, "test_key", "test_list")
}
