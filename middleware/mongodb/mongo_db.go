package mongodb

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func Connect(appCtx context.Context, webEngine *gin.Engine) *mongo.Database {
	// create context in 10 second
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect in mongodb
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ := mongo.Connect(ctx, clientOptions)

	db := client.Database("sessionData")
	return db
}

func SessionData(client *mongo.Client, CollectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("sessionData").Collection("sessionData")
	return collection

}
