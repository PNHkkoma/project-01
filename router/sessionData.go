package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"xrplatform/arworld/backend/models"
)

func GetSessionData(context *gin.Context) {
	//var SessionID string
	//var sessionData string
	//
	//if context.ShouldBind(&SessionID) == nil {
	//	// log content
	//	log.Println(SessionID)
	//}
	//DB, exist := context.Get("db")
	//db := DB.(*sql.DB)
	//log.Println(exist)
	//
	////check already exists
	//querryData := `SELECT SessionData FROM sessiondata WHERE SessionID = ?`
	//checkData := db.QueryRow(querryData, context.Request.Body).Scan(&sessionData)
	//
	//log.Println(checkData, "gì đó", SessionID)
	//
	//if checkData != nil {
	//	// response Json for client
	//	context.JSON(409, gin.H{
	//		"status": 409,
	//		"error":  "Data error",
	//	})
	//} else {
	//	// response Json for client
	//	context.JSON(200, gin.H{
	//		"status": 200,
	//		"data":   checkData,
	//	})
	//}
	context.JSON(200, gin.H{
		"status": 200,
		"data":   "data error",
	})

}

func UploadSessionData(context *gin.Context) {
	// declare form data for session
	var formData models.SessionFormData

	// verify data match type of SessionFormData
	if context.ShouldBind(&formData) == nil {
		// log content
		log.Println(formData.SessionData)
		log.Println(formData.SessionID)
	}

	DB, exist := context.Get("db")
	db := DB.(*sql.DB)
	log.Println(exist)

	// check already exists
	querryExists := `SELECT SessionID FROM sessiondata WHERE SessionID = ?`
	checkExist := db.QueryRow(querryExists, formData.SessionID).Scan(&formData.SessionID)
	log.Println(checkExist)

	if checkExist == nil {
		context.JSON(409, gin.H{
			"error": "SessionID already exists",
		})
	} else {
		// save data
		result, err := db.Exec(`INSERT INTO sessiondata (SessionID, SessionData) VALUES (?, ?)`, formData.SessionID, formData.SessionData)
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
