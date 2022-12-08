package dbsupport

import (
	"context"
	"crawler/distributed/config"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
		Password: "",
	})
}

func SaveToRedis(client *redis.Client, ctx context.Context, key string, value bool) {
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Fatalf("[SaveToRedis] error save url to redis: %v", err)
	}
}

func GetFromRedis(client *redis.Client, ctx context.Context, key string) string {
	result, err := client.Get(ctx, key).Result()
	// 当 key 不存在时，redis 会返回 redis:nil 的错误信息
	if err != nil {
		return ""
	}
	return result
}

func GetOrSaveFromRedis(client *redis.Client, ctx context.Context, key string) bool {
	if GetFromRedis(client, ctx, key) != "" {
		return true
	}
	SaveToRedis(client, ctx, key, true)
	return false
}
