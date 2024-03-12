package routes

import (
	"log"
	"time"
	"xrplatform/arworld/backend/env"
	"xrplatform/arworld/backend/middleware/mysql"
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

	// session data variable
	var sessionData string

	// get redis_cli client from ctx
	redisClient := redis_cli.GetClient(ctx)

	if redisClient == nil {
		log.Println("cannot connect to redis_cli")
		ctx.JSON(200, gin.H{
			"status": 500,
			"error":  "cannot connect to redis_cli",
		})
		return
	}

	// get data from sessionID in redis_cli
	sessionData = redis_cli.GetSessionDataFromRedis(appCtx, redisClient, formData.SessionID)

	if sessionData != "" {
		// response Json for client
		ctx.JSON(200, gin.H{
			"status": 200,
			"data":   sessionData,
		})
		return
	}

	// get db client from ctx
	db := mysql.GetDB(ctx)

	if db == nil {
		log.Println("cannot connect to db")
		ctx.JSON(200, gin.H{
			"status": 500,
			"error":  "cannot connect to db",
		})
		return
	}

	//check already exists
	scanCode := db.QueryRow(mysql.GetSessionDataQuery,
		formData.SessionID).Scan(&sessionData)

	if scanCode != nil {
		// response Json for client
		ctx.JSON(200, gin.H{
			"status": 500,
			"error":  "Data error",
		})
	} else {
		// add data to redis_cli
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
	db := mysql.GetDB(ctx)

	if db == nil {
		log.Println("cannot connect to db")
		ctx.JSON(200, gin.H{
			"status": 500,
			"error":  "cannot connect to db",
		})
		return
	}

	// check already exists
	checkExist := db.QueryRow(mysql.GetSessionDataQuery,
		formData.SessionID).Scan(&formData.SessionID)

	if checkExist == nil {
		ctx.JSON(200, gin.H{
			"status": 500,
			"error":  "session ID already exists",
		})
	} else {
		// save data to db
		result, err := db.Exec(mysql.InsertSessionDataQuery,
			formData.SessionID, formData.SessionData)

		if err != nil {
			log.Println(err)
			ctx.JSON(200, gin.H{
				"status": 500,
				"data":   "fail to upload data",
			})
		} else {
			log.Println(result)
			ctx.JSON(200, gin.H{
				"status": 200,
				"data":   "success",
			})
		}
	}
}
