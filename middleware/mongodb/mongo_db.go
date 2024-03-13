package mongodb

import (
	"context"
	"fmt"
	"time"
	"xrplatform/arworld/backend/env"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(appCtx context.Context, webEngine *gin.Engine) {
	// create context in 10 second
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// create mongodb connect
	dbConn := env.GetAppKey(appCtx, "mongo_conn").(string)
	clientOptions := options.Client().ApplyURI(dbConn)
	client, _ := mongo.Connect(ctx, clientOptions)

	// get database
	mongoDBName := env.GetAppKey(appCtx, "mongo_dbname").(string)
	db := client.Database(mongoDBName)

	webEngine.Use(func(ctx *gin.Context) {
		ctx.Set("mongo_db", db)
	})
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
	mongoHost := env.GetAppEnv(env.MongoHost)
	mongoPort := env.GetAppEnv(env.MongoPort)
	mongoDBName := env.GetAppEnv(env.MongoDB)

	//set URI mongo
	connection := fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)

	env.SetAppKey(appCtx, "mongo_conn", connection)
	env.SetAppKey(appCtx, "mongo_dbname", mongoDBName)
}
