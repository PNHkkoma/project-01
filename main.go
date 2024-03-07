package main

import (
	"log"
	"xrplatform/arworld/backend/middleware/mysql"
	"xrplatform/arworld/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// create new web engine
	webEngine := gin.Default()

	// connect db
	db := mysql.ConnectDB(webEngine)
	defer db.Close()

	// define all router for backend
	routes.DefineRouters(webEngine)

	// run web service and log error
	if err := webEngine.Run(); err != nil {
		log.Println("start server failed")
	}
}
