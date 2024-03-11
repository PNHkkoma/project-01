package redis_cli

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func ConnectRedis(webEngine *gin.Engine) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// check redis connection
	ctx := context.Background()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Println("Failed to connect to Redis: ", err)
	}

	// add redis_cli.Client to gin.Engine
	webEngine.Use(func(context *gin.Context) {
		context.Set("redis_cli", redisClient)
	})
	return redisClient
}

func GetRedisClient(context *gin.Context) *redis.Client {
	// get redis_cli from context
	redisClient, exist := context.Get("redis_cli")
	if !exist {
		return nil
	} else {
		return redisClient.(*redis.Client)
	}
}

func SetRedisSessionData(redisDB *redis.Client, key string, value string, duration time.Duration) {
	// create empty context
	ctx := context.Background()

	// set data to redis_cli
	err := redisDB.Set(ctx, fmt.Sprintf("sessiondata:%s", key), value, duration).Err()
	if err == nil {
		log.Printf("Set data to redis_cli: key=%s value=%s", key, value)
	} else {
		log.Printf("Cannot set key=%s to redis_cli", key)
	}
}

func GetRedisSessionData(redisDB *redis.Client, key string) string {
	// create empty context
	ctx := context.Background()

	// set data to redis_cli
	value := redisDB.Get(ctx, fmt.Sprintf("sessiondata:%s", key)).Val()
	return value
}
