package main

import (
	"database/sql"
	"log"
	"xrplatform/arworld/backend/middleware/mysql"
	"xrplatform/arworld/backend/middleware/redis_cli"
	"xrplatform/arworld/backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	// create new web engine
	webEngine := gin.Default()

	// connect db
	db := mysql.ConnectDB(webEngine)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	redisClient := redis_cli.ConnectRedis(webEngine)
	defer func(redis *redis.Client) {
		_ = redis.Close()
	}(redisClient)

	// define all router for backend
	routes.DefineRoutes(webEngine)

	// run web service and log error
	if err := webEngine.Run(); err != nil {
		log.Println("start server failed")
	}
}
