package routes

import (
	"log"
	"xrplatform/arworld/backend/middleware/mongodb"
	"xrplatform/arworld/backend/models"

	"github.com/gin-gonic/gin"
)

func GetCategoryData(ctx *gin.Context) {
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
	results, err := mongodb.QueryGetCategoryData(db)

	if err != nil {
		// response Json for client
		ctx.JSON(200, gin.H{
			"status": 500,
			"error":  "Data error",
		})
	} else {
		// response Json for client
		ctx.JSON(200, gin.H{
			"status": 200,
			"data":   results,
		})
	}
}

func GetStoreByCategoryData(ctx *gin.Context) {
	// get db client from ctx
	db := mongodb.GetDB(ctx)

	// declare form data for session
	var formData models.StoreDataByCategory

	// verify data match type of SessionData
	if ctx.ShouldBind(&formData) != nil {
		// log error here
		log.Println("cannot bind to form data")
		return
	}

	if db == nil {
		log.Println("cannot connect to db")
		ctx.JSON(200, gin.H{
			"status": 500,
			"error":  "cannot connect to db",
		})
		return
	}

	//check already exists
	results, err := mongodb.QueryGetStoreByCategory(db, formData.Category)

	if err != nil {
		// response Json for client
		ctx.JSON(200, gin.H{
			"status": 500,
			"error":  "Data error",
		})
	} else {
		// response Json for client
		ctx.JSON(200, gin.H{
			"status": 200,
			"data":   results,
		})
	}
}

func SearchStoreByStoreName(ctx *gin.Context) {
	// get db client from ctx
	db := mongodb.GetDB(ctx)

	// declare form data for session
	var formData models.StoreDataSearchKey

	// verify data match type of SessionData
	if ctx.ShouldBind(&formData) != nil {
		// log error here
		log.Println("cannot bind to form data")
		return
	}

	if db == nil {
		log.Println("cannot connect to db")
		ctx.JSON(200, gin.H{
			"status": 500,
			"error":  "cannot connect to db",
		})
		return
	}

	//check already exists
	results, err := mongodb.QuerySearchStoreByName(db, formData.QueryKey)

	if err != nil {
		// response Json for client
		ctx.JSON(200, gin.H{
			"status": 500,
			"error":  "Data error",
		})
	} else {
		// response Json for client
		ctx.JSON(200, gin.H{
			"status": 200,
			"data":   results,
		})
	}
}
