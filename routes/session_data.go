package routes

import (
	"log"
	"time"
	"xrplatform/arworld/backend/middleware/mysql"
	"xrplatform/arworld/backend/middleware/redis_cli"
	"xrplatform/arworld/backend/models"

	"github.com/gin-gonic/gin"
)

func GetSessionData(context *gin.Context) {
	// declare form data for session
	var formData models.SessionGetData

	// verify data match type of SessionUploadData
	if context.ShouldBind(&formData) != nil {
		// log error here
		log.Println("cannot bind to form data")
		return
	}

	// session data variable
	var sessionData string

	// get redis_cli client from context
	redisClient := redis_cli.GetRedisClient(context)

	if redisClient == nil {
		log.Println("cannot connect to redis_cli")
		context.JSON(200, gin.H{
			"status": 500,
			"error":  "cannot connect to redis_cli",
		})
		return
	}

	// get data from sessionID in redis_cli
	sessionData = redis_cli.GetRedisSessionData(redisClient, formData.SessionID)

	if sessionData != "" {
		// response Json for client
		context.JSON(200, gin.H{
			"status": 200,
			"data":   sessionData,
		})
		return
	}

	// get db client from context
	db := mysql.GetDBFromContext(context)

	if db == nil {
		log.Println("cannot connect to db")
		context.JSON(200, gin.H{
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
		context.JSON(200, gin.H{
			"status": 500,
			"error":  "Data error",
		})
	} else {
		// add data to redis_cli
		redis_cli.SetRedisSessionData(redisClient, formData.SessionID, sessionData, 5*time.Minute)

		// response Json for client
		context.JSON(200, gin.H{
			"status": 200,
			"data":   sessionData,
		})
	}
}

func UploadSessionData(context *gin.Context) {
	// declare form data for session
	var formData models.SessionUploadData

	// verify data match type of SessionUploadData
	if context.ShouldBind(&formData) != nil {
		// log error here
		return
	}

	// get db client from context
	db := mysql.GetDBFromContext(context)

	if db == nil {
		log.Println("cannot connect to db")
		context.JSON(200, gin.H{
			"status": 500,
			"error":  "cannot connect to db",
		})
		return
	}

	// check already exists
	checkExist := db.QueryRow(mysql.GetSessionDataQuery,
		formData.SessionID).Scan(&formData.SessionID)

	if checkExist == nil {
		context.JSON(200, gin.H{
			"status": 500,
			"error":  "session ID already exists",
		})
	} else {
		// save data to db
		result, err := db.Exec(mysql.InsertSessionDataQuery,
			formData.SessionID, formData.SessionData)

		if err != nil {
			log.Println(err)
			context.JSON(200, gin.H{
				"status": 500,
				"data":   "fail to upload data",
			})
		} else {
			log.Println(result)
			context.JSON(200, gin.H{
				"status": 200,
				"data":   "success",
			})
		}
	}
}
