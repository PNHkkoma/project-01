package routes

import "github.com/gin-gonic/gin"

func DefineRoutes(webEngine *gin.Engine) {
	// route for session data
	webEngine.POST("/ar-world/v1/session-data/upload", UploadSessionData)
	webEngine.POST("/ar-world/v1/session-data/get", GetSessionData)
}
