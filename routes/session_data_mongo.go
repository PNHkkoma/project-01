package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
	"xrplatform/arworld/backend/middleware/mongodb"
	"xrplatform/arworld/backend/models"
)

func GetSessionDataMongo(ctx *gin.Context) {
	// declare form data for session
	var formData models.Data
	ctx.Header("Content-Type", "application/json")
	// verify data match type of SessionUploadData
	if err := ctx.ShouldBind(&formData); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid JSON data"})
		return
	}

	dbConnect := mongodb.DBSet()
	db := mongodb.SessionData(dbConnect, "sessionData")
	ctx2, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var result models.Data

	err := db.FindOne(ctx2, bson.M{"_id": formData.ID}).Decode(result)
	if err != nil {
		log.Println("Error finding document:", formData.ID)
		return
	}
	ctx.JSON(200, gin.H{
		"data": result,
	})
	////check already exists
	//scanCode := db.QueryRow(mysql.GetSessionDataQuery,
	//	formData.SessionID).Scan(&sessionData)
	//
	//if scanCode != nil {
	//	// response Json for client
	//	context.JSON(200, gin.H{
	//		"status": 500,
	//		"error":  "Data error",
	//	})
	//} else {
	//	// add data to redis_cli
	//	redis_cli.SetRedisSessionData(redisClient, formData.SessionID, sessionData, 5*time.Minute)
	//
	//	// response Json for client
	//	context.JSON(200, gin.H{
	//		"status": 200,
	//		"data":   sessionData,
	//	})
	//}
}

//func UploadSessionDataMongo(context *gin.Context) {
//	// declare form data for session
//	var formData models.SessionUploadData
//
//	// verify data match type of SessionUploadData
//	if context.ShouldBind(&formData) != nil {
//		// log error here
//		return
//	}
//
//	// get db client from context
//	db := mysql.GetDBFromContext(context)
//
//	if db == nil {
//		log.Println("cannot connect to db")
//		context.JSON(200, gin.H{
//			"status": 500,
//			"error":  "cannot connect to db",
//		})
//		return
//	}
//
//	// check already exists
//	checkExist := db.QueryRow(mysql.GetSessionDataQuery,
//		formData.SessionID).Scan(&formData.SessionID)
//
//	if checkExist == nil {
//		context.JSON(200, gin.H{
//			"status": 500,
//			"error":  "session ID already exists",
//		})
//	} else {
//		// save data to db
//		result, err := db.Exec(mysql.InsertSessionDataQuery,
//			formData.SessionID, formData.SessionData)
//
//		if err != nil {
//			log.Println(err)
//			context.JSON(200, gin.H{
//				"status": 500,
//				"data":   "fail to upload data",
//			})
//		} else {
//			log.Println(result)
//			context.JSON(200, gin.H{
//				"status": 200,
//				"data":   "success",
//			})
//		}
//	}
//}
