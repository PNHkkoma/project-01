package mysql

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(webEngine *gin.Engine) *sql.DB {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/sessiondata?parseTime=true")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	// check db connection
	err = db.Ping()
	if err != nil {
		log.Println("Failed to connect to db: ", err)
	}

	// add sql.DB to gin.Engine
	webEngine.Use(func(context *gin.Context) {
		context.Set("mysql_db", db)
	})
	return db
}

func GetDBFromContext(context *gin.Context) *sql.DB {
	// get db from context
	db, exist := context.Get("mysql_db")
	if !exist {
		return nil
	} else {
		return db.(*sql.DB)
	}
}
