package mongodb

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"xrplatform/arworld/backend/env"
)

func Connect(appCtx context.Context, webEngine *gin.Engine) *mongo.Database {
	// create context in 10 second
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect in mongodb
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ := mongo.Connect(ctx, clientOptions)

	db := client.Database("sessionData")

	webEngine.Use(func(ctx *gin.Context) {
		ctx.Set("mongo_db", db)
	})
	return db
}

func GetDB(ctx *gin.Context) *mongo.Database {
	// get db from context
	db, exist := ctx.Get("mongo_db")
	if exist {
		return db.(*mongo.Database)
	} else {
		return nil
	}
}

func GetEnv(appCtx context.Context) {
	// database env
	mySQLHost := env.GetAppEnv(env.MongoHost)
	mySQLPort := env.GetAppEnv(env.MongoPort)
	mySQLDBName := env.GetAppEnv(env.MongoDB)

	//set URI mongo
	connection := fmt.Sprintf("mongodb://localhost:27017", mySQLHost, mySQLPort, mySQLDBName)

	env.SetAppKey(appCtx, "mongo_conn", connection)
}
