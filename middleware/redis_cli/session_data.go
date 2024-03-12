package redis_cli

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func SetSessionDataToRedis(appCtx context.Context, redisDB *redis.Client, key string, value string, duration time.Duration) {
	// set data to redis_cli
	err := redisDB.Set(appCtx, fmt.Sprintf("sessiondata:%s", key), value, duration).Err()
	if err == nil {
		log.Printf("Set data to redis_cli: key=%s value=%s", key, value)
	} else {
		log.Printf("Cannot set key=%s to redis_cli", key)
	}
}

func GetSessionDataFromRedis(appCtx context.Context, redisDB *redis.Client, key string) string {
	// set data to redis_cli
	value := redisDB.Get(appCtx, fmt.Sprintf("sessiondata:%s", key)).Val()
	return value
}
