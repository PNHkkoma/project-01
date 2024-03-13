package routes

import (
	"log"
	"time"
	"xrplatform/arworld/backend/env"
	"xrplatform/arworld/backend/middleware/mongodb"
	"xrplatform/arworld/backend/middleware/redis_cli"
	"xrplatform/arworld/backend/models"

	"github.com/gin-gonic/gin"
)

func GetSessionData(ctx *gin.Context) {
	appCtx := env.GetContext(ctx)

	// declare form data for session
	var formData models.SessionGetData

	// verify data match type of SessionUploadData
	if ctx.ShouldBind(&formData) != nil {
		// log error here
		log.Println("cannot bind to form data")
		return
	}

	// get redis client from ctx
	redisClient := redis_cli.GetClient(ctx)

	if redisClient == nil {
		log.Println("cannot connect to redis")
		ctx.JSON(200, gin.H{
			"status": 500,
			"error":  "cannot connect to redis",
		})
		return
	}

	// get data from sessionID in redis
	cacheData := redis_cli.GetSessionDataFromRedis(appCtx, redisClient, formData.SessionID)

	if cacheData != "" {
		// response Json for client
		ctx.JSON(200, gin.H{
			"status": 200,
			"data":   cacheData,
		})
		return
	}

	// get db client from ctx
	db := mongodb.GetDB(ctx)

	if db == nil {
		log.Println("cannot connect to db")
		ctx.JSON(200, gin.H{
			"status": 500,
			"error":  "cannot connect to db",
		})
		return
	}

	//check already exists
	sessionData, scanCode := mongodb.QueryGetSessionData(db, formData.SessionID)

	if scanCode != nil {
		// response Json for client
		ctx.JSON(200, gin.H{
			"status": 500,
			"error":  "Data error",
		})
	} else {
		// add data to redis
		redis_cli.SetSessionDataToRedis(appCtx, redisClient, formData.SessionID,
			sessionData, 5*time.Minute)

		// response Json for client
		ctx.JSON(200, gin.H{
			"status": 200,
			"data":   sessionData,
		})
	}
}

func UploadSessionData(ctx *gin.Context) {
	// declare form data for session
	var formData models.SessionUploadData

	// verify data match type of SessionUploadData
	if ctx.ShouldBind(&formData) != nil {
		// log error here
		return
	}

	// get db client from ctx
	db := mongodb.GetDB(ctx)

	if db == nil {
		log.Println("cannot connect to db")
		ctx.JSON(200, gin.H{
			"status": 500,
			"error":  "cannot connect to db",
		})
		return
	}

	// save data to db
	err := mongodb.QueryUploadSessionData(db, formData.SessionID, formData.SessionData)

	if err != nil {
		log.Println(err)
		ctx.JSON(200, gin.H{
			"status": 500,
			"data":   "fail to upload data",
		})
	} else {
		ctx.JSON(200, gin.H{
			"status": 200,
			"data":   "success",
		})
	}
}
