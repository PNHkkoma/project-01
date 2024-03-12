package redis_cli

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"xrplatform/arworld/backend/env"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Connect(ctx context.Context, webEngine *gin.Engine) *redis.Client {
	// get redis client
	redisClient := getClient(ctx)

	// check redis connection
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Println("Failed to connect to Redis: ", err)
	}

	// add redis.Client to gin.Engine
	webEngine.Use(func(context *gin.Context) {
		context.Set("redis_cli", redisClient)
	})
	return redisClient
}

func Close(redis *redis.Client) {
	_ = redis.Close()
}

func GetClient(context *gin.Context) *redis.Client {
	// get redis from context
	redisClient, exist := context.Get("redis_cli")
	if exist {
		return redisClient.(*redis.Client)
	} else {
		return nil
	}
}

func GetEnv(ctx context.Context) {
	// define redis connection
	redisHost := env.GetAppEnv(env.RedisHost)
	redisPort := env.GetAppEnv(env.RedisPort)
	redisConn := fmt.Sprintf("%s:%s", redisHost, redisPort)

	// define redis db and password
	redisPwd := env.GetAppEnv(env.RedisPassword)
	redisDB, _ := strconv.Atoi(env.GetAppEnv(env.RedisDB))

	env.SetAppKey(ctx, "redis_conn", redisConn)
	env.SetAppKey(ctx, "redis_pwd", redisPwd)
	env.SetAppKey(ctx, "redis_db", redisDB)
}

func getClient(ctx context.Context) *redis.Client {
	// get data from context
	redisConn := env.GetAppKey(ctx, "redis_conn").(string)
	redisPwd := env.GetAppKey(ctx, "redis_pwd").(string)
	redisDB := env.GetAppKey(ctx, "redis_db").(int)

	// create redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisConn,
		Password: redisPwd,
		DB:       redisDB,
	})
	return redisClient
}
