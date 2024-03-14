package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func QueryGetSessionData(db *mongo.Database, sessionID string) (string, error) {
	var collection = db.Collection("Sessions")
	// create variable to save bson value
	var result bson.M
	err := collection.FindOne(context.Background(), bson.M{"SessionID": sessionID}).Decode(&result)

	data, found := result["SessionData"].(string)
	if !found {
		log.Println("Session data not found")
	}
	return data, err
}

func QueryUploadSessionData(db *mongo.Database, sessionID string, sessionData string) error {
	var collection = db.Collection("Sessions")
	newSessionData := bson.D{
		{Key: "SessionID", Value: sessionID},
		{Key: "SessionData", Value: sessionData},
	}
	_, err := collection.InsertOne(context.Background(), newSessionData)
	return err
}
