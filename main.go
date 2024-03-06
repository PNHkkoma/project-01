package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xrplatform/arworld/backend/middleware"
	"xrplatform/arworld/backend/router"
)

func main() {

	// create new web engine
	webEngine := gin.Default()

	// define all router for backend
	router.DefineRouters(webEngine)

	middleware.ConnectMySQL(webEngine)

	// run web service and log error
	if err := webEngine.Run(); err != nil {
		log.Println("start server failed")
	}

}
