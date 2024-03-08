package main

import (
	"database/sql"
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
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	// define all router for backend
	routes.DefineRoutes(webEngine)

	// run web service and log error
	if err := webEngine.Run(); err != nil {
		log.Println("start server failed")
	}
}
