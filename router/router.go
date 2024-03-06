package router

import "github.com/gin-gonic/gin"

func DefineRouters(webEngine *gin.Engine) {

	// route for session data
	webEngine.POST("/ar-world/v1/session-data/upload", UploadSessionData)
	webEngine.GET("/ar-world/v1/session-data/get", GetSessionData)

}
