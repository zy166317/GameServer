package redis

import (
	"GameServer/pkg/db"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

// InitRedis 初始化redis
func InitRedis(re *db.Redis) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", re.Host, re.Port),
		DB:       re.DBIndex,
		PoolSize: re.PoolSize,
	})
}

// GetRedisClient 获取redis客户端
func GetRedisClient() *redis.Client {
	return redisClient
}
