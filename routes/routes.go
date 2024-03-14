package routes

import "github.com/gin-gonic/gin"

func DefineRoutes(webEngine *gin.Engine) {
	// route for session data
	webEngine.POST("/ar-world/v1/session-data/upload", UploadSessionData)
	webEngine.POST("/ar-world/v1/session-data/get", GetSessionData)

	// route for stores data
	webEngine.POST("/ar-world/v1/store-data/search", SearchStoreByStoreName)
	webEngine.POST("/ar-world/v1/store-data/category/get", GetCategoryData)
	webEngine.POST("/ar-world/v1/store-data/category/get-stores", GetStoreByCategoryData)
}
