package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func QueryGetSessionData(db *mongo.Database, sessionID string, ctx context.Context) (string, error) {
	var collection = db.Collection("sessionData")
	// Tạo một biến để lưu trữ dữ liệu trả về
	var result bson.M
	err := collection.FindOne(ctx, bson.M{"SessionID": sessionID}).Decode(&result)

	data, found := result["SessionData"].(string)
	if !found {
		log.Println("Session data not found")
	}
	return data, err
}

func QueryUploadSessionData(db *mongo.Database, sessionID string, sessionData string) error {
	var collection = db.Collection("sessionData")
	newSessionData := bson.D{
		{"SessionID", sessionID},
		{"SessionData", sessionData},
	}
	_, err := collection.InsertOne(context.Background(), newSessionData)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
