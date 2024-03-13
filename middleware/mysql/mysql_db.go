package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"xrplatform/arworld/backend/env"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Connect(appCtx context.Context, webEngine *gin.Engine) *sql.DB {
	db := getDB(appCtx)

	// check db connection
	err := db.Ping()
	if err != nil {
		log.Println("Failed to connect to db: ", err)
	}

	// add sql.DB to gin.Engine
	webEngine.Use(func(ctx *gin.Context) {
		ctx.Set("mysql_db", db)
	})
	return db
}

func Close(db *sql.DB) {
	_ = db.Close()
}

func GetDB(ctx *gin.Context) *sql.DB {
	// get db from context
	db, exist := ctx.Get("mysql_db")
	if exist {
		return db.(*sql.DB)
	} else {
		return nil
	}
}

func GetEnv(appCtx context.Context) {
	// database env
	mySQLHost := env.GetAppEnv(env.MySqlHost)
	mySQLPort := env.GetAppEnv(env.MySqlPort)
	mySQLDBName := env.GetAppEnv(env.MySqlDB)
	connection := fmt.Sprintf("root:@(%s:%s)/%s?parseTime=true", mySQLHost, mySQLPort, mySQLDBName)

	env.SetAppKey(appCtx, "mysql_conn", connection)
}

func getDB(appCtx context.Context) *sql.DB {
	dbConn := env.GetAppKey(appCtx, "mysql_conn").(string)
	log.Println(dbConn)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	return db
}
