package middleware

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func ConnectMySQL(webEngine *gin.Engine) {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/sessiondata?parseTime=true")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	// defer db.Close()

	webEngine.Use(func(c *gin.Context) { c.Set("db", db) })
}

func GetDBFromContext(context *gin.Context) *sql.DB {
	db, exist := context.Get("db")

	if !exist {
		return nil
	} else {
		return db.(*sql.DB)
	}
}
