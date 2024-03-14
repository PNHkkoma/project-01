package mongodb

import (
	"context"
	"log"
	"xrplatform/arworld/backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func QueryGetCategoryData(db *mongo.Database) ([]interface{}, error) {
	collection := db.Collection("Stores")
	ctx := context.Background()

	// create variable to save bson value
	results, err := collection.Distinct(ctx, "Category", bson.M{})
	log.Println(results)

	if err != nil {
		log.Println("Cannot get category data")
	}

	return results, err
}

func QueryGetStoreByCategory(db *mongo.Database, category string) ([]models.StoreData, error) {
	collection := db.Collection("Stores")
	ctx := context.Background()

	// create variable to save bson value
	cursor, err := collection.Find(ctx, bson.M{"Category": category})

	if err != nil {
		log.Println("Cannot get store data in collection")
	}

	var results []models.StoreData

	// get result
	err = cursor.All(ctx, &results)

	if err != nil {
		log.Println("Cannot get data from mongo cursor")
	}
	return results, err
}

func QuerySearchStoreByName(db *mongo.Database, queryKey string) ([]models.StoreData, error) {
	collection := db.Collection("Stores")
	ctx := context.Background()

	// create variable to save bson value
	queryOpts := bson.M{"$regex": queryKey, "$options": "i"}
	cursor, err := collection.Find(ctx, bson.M{"StoreName": queryOpts})

	if err != nil {
		log.Println("Cannot get store data in collection")
	}

	var results []models.StoreData

	// get result
	err = cursor.All(ctx, &results)

	if err != nil {
		log.Println("Cannot get data from mongo cursor")
	}
	return results, err
}
