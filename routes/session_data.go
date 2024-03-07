package routes

import (
	"log"
	"xrplatform/arworld/backend/middleware/mysql"
	"xrplatform/arworld/backend/models"

	"github.com/gin-gonic/gin"
)

func GetSessionData(context *gin.Context) {
	// declare form data for session
	var formData models.SessionGetData

	// verify data match type of SessionUploadData
	if context.ShouldBind(&formData) != nil {
		// log error here
		return
	}

	db := mysql.GetDBFromContext(context)
	// check db == nil

	var sessionData string
	//check already exists

	scanCode := db.QueryRow(mysql.GetSessionDataQuery,
		formData.SessionID).Scan(&sessionData)
	log.Println(sessionData)

	if scanCode != nil {
		// response Json for client
		context.JSON(409, gin.H{
			"status": 409,
			"error":  "Data error",
		})
	} else {
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

	db := mysql.GetDBFromContext(context)
	// check db == nil

	// check already exists
	checkExist := db.QueryRow(mysql.GetSessionDataQuery,
		formData.SessionID).Scan(&formData.SessionID)
	log.Println(checkExist)

	if checkExist == nil {
		context.JSON(409, gin.H{
			"error": "SessionID already exists",
		})
	} else {
		// save data
		result, err := db.Exec(mysql.InsertSessionDataQuery,
			formData.SessionID, formData.SessionData)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(result)
		}
		context.JSON(200, gin.H{
			"status": 200,
			"data":   formData,
		})
	}
}
