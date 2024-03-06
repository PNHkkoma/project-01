package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	webEngine := gin.Default()
	if err := webEngine.Run("localhost:9001"); err != nil {
		log.Println("start server failed")
	}
}
