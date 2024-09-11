package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var (
	redisClient *redis.Client          // Redis client
	ctx         = context.Background() // Redis 操作的 context
)

// InitRedis 初始化 Redis 连接
func InitRedis() error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 地址
		Password: "",               // 没有密码
		DB:       0,                // 默认DB
	})

	// 测试连接
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}

	log.Println("Redis connection established")
	return nil
}

// GetRedisClient 获取 Redis 客户端
func GetRedisClient() *redis.Client {
	if redisClient == nil {
		log.Fatal("Redis client is not initialized. Call InitRedis first.")
	}
	return redisClient
}
