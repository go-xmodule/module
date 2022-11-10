/**
* Created by GoLand
* @file redis.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date 2022/5/10 10:27 上午
* @desc redis.go
 */

package dirver

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisClient struct {
	client       *redis.Client
	backupClient *redis.Client
	context      context.Context
}
type SubscribeData struct {
	Channel      string      `json:"Channel"`
	Pattern      string      `json:"Pattern"`
	Payload      string      `json:"Payload"`
	PayloadSlice interface{} `json:"PayloadSlice"`
}

var Redis RedisClient

// NewRedis 实例化
func NewRedis() *RedisClient {
	redisClient := new(RedisClient)
	redisClient.context = context.Background()
	return redisClient
}

// Connect 链接服务
func (r *RedisClient) Connect(host string, port int, password string, db int) (*RedisClient, error) {
	address := fmt.Sprintf("%s:%d", host, port)
	r.client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set
		DB:       db,       // use default DB

		// 连接池容量及闲置连接数量
		PoolSize:     100, // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU
		MinIdleConns: 40,  // 在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。

		// 超时
		DialTimeout:  5 * time.Second, // 连接建立超时时间，默认5秒。
		ReadTimeout:  3 * time.Second, // 读超时，默认3秒， -1表示取消读超时
		WriteTimeout: 3 * time.Second, // 写超时，默认等于读超时
		PoolTimeout:  4 * time.Second, // 当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。

		// 闲置连接检查包括IdleTimeout，MaxConnAge
		IdleCheckFrequency: 60 * time.Second, // 闲置连接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理。
		IdleTimeout:        5 * time.Minute,  // 闲置超时，默认5分钟，-1表示取消闲置超时检查
		MaxConnAge:         0 * time.Second,  // 连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接

		// 命令执行失败时的重试策略
		MaxRetries:      2,                      // 命令执行失败时，最多重试多少次，默认为0即不重试
		MinRetryBackoff: 8 * time.Millisecond,   // 每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
		MaxRetryBackoff: 512 * time.Millisecond, // 每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔
	})
	// 判断是否能够链接到数据库
	_, err := r.client.Ping(context.Background()).Result()
	r.backupClient = r.client
	return r, err
}

// Set 字符串设置
func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	// 执行命令
	err := r.client.Set(r.context, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) HSetNX(key, field string, value interface{}) (bool, error) {
	// 执行命令
	result := r.client.HSetNX(r.context, key, field, value)
	if err := result.Err(); err != nil {
		return false, err
	}
	return result.Result()
}

// HSet hash-set
func (r *RedisClient) HSet(key string, values map[string]string) (int64, error) {
	// 执行命令
	result := r.client.HSet(r.context, key, values)
	if err := result.Err(); err != nil {
		return 0, err
	}
	return result.Result()
}

func (r *RedisClient) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	// 执行命令
	// r.client.Conn(r.context).Select(r.context,1)
	result := r.client.SetNX(r.context, key, value, expiration)
	if err := result.Err(); err != nil {
		return false, err
	}
	return result.Result()
}

func (r *RedisClient) HGetAll(key string) (map[string]string, error) {
	// r.checkConnect()
	// 执行命令
	result := r.client.HGetAll(r.context, key)
	if err := result.Err(); err != nil {
		return nil, err
	}
	return result.Result()
}

func (r *RedisClient) Get(key string) (bool, string, error) {
	// r.checkConnect()
	result, err := r.client.Get(r.context, key).Result()
	if err == redis.Nil {
		return false, "", nil
	} else if err != nil {
		return false, "", err
	} else {
		return true, result, nil
	}
}

func (r *RedisClient) Delete(key ...string) error {
	// r.checkConnect()
	err := r.client.Del(r.context, key...).Err()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (r *RedisClient) LPop(key string) (string, error) {
	result := r.client.LPop(r.context, key)
	if err := result.Err(); err != nil {
		return "", err
	} else {
		return result.Result()
	}
}

func (r *RedisClient) LPush(key string, value interface{}) (int64, error) {
	result := r.client.LPush(r.context, key, value)
	if err := result.Err(); err != nil {
		return 0, err
	} else {
		return result.Result()
	}
}

func (r *RedisClient) BLPop(key string, timeout time.Duration) ([]string, error) {
	result := r.client.BLPop(r.context, timeout, key)
	if err := result.Err(); err != nil {
		return []string{}, err
	} else {
		return result.Result()
	}
}

func (r *RedisClient) Keys(pattern string) ([]string, error) {
	result := r.client.Keys(r.context, pattern)
	if err := result.Err(); err != nil {
		return nil, err
	} else {
		return result.Result()
	}
}

// Publish 发布
func (r *RedisClient) Publish(channel string, message interface{}) (int64, error) {
	result := r.client.Publish(r.context, channel, message)
	if err := result.Err(); err != nil {
		return 0, err
	} else {
		return result.Result()
	}
}

// Subscribe 订阅
func (r *RedisClient) Subscribe(channel string) <-chan *redis.Message {
	result := r.client.Subscribe(r.context, channel)
	return result.Channel()
}
