package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"xrplatform/arworld/backend/models"
)

func GetSessionData(context *gin.Context) {
	// response Json for client

	context.JSON(200, gin.H{
		"status": 200,
		"data":   "Error code",
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
	context.JSON(200, gin.H{
		"status": 200,
		"data":   formData,
	})
}
