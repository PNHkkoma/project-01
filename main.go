package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"xrplatform/arworld/backend/env"
	"xrplatform/arworld/backend/middleware/mysql"
	"xrplatform/arworld/backend/middleware/redis_cli"
	"xrplatform/arworld/backend/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type Application struct {
	dataCollection *mongo.Collection
}

func NewApplication(dataCollection *mongo.Collection) *Application {
	return &Application{
		dataCollection: dataCollection,
	}
}

func main() {
	// create new web engine
	webEngine := gin.Default()

	// get app context
	appCtx := env.GetAppContext()

	// add context to gin
	env.SetContext(appCtx, webEngine)

	// db connection
	mysql.GetEnv(appCtx)

	// connect db
	db := mysql.Connect(appCtx, webEngine)
	defer mysql.Close(db)

	// redis connection
	redis_cli.GetEnv(appCtx)

	// connect redis
	redisClient := redis_cli.Connect(appCtx, webEngine)
	defer redis_cli.Close(redisClient)

	// define all router for backend
	routes.DefineRoutes(webEngine)

	// run web service and log error
	if err := webEngine.Run(); err != nil {
		log.Println("start server failed")
	}
}
