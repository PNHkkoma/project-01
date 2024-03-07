package mysql

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(webEngine *gin.Engine) *sql.DB {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/sessiondata?parseTime=true")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	// add sql.DB to gin.Engine
	webEngine.Use(func(context *gin.Context) {
		context.Set("db", db)
	})
	return db
}

func GetDBFromContext(context *gin.Context) *sql.DB {
	// get db from context
	db, exist := context.Get("db")

	if !exist {
		return nil
	} else {
		return db.(*sql.DB)
	}
}
