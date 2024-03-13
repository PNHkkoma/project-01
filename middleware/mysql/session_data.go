package mysql

import (
	"database/sql"
)

func QueryGetSessionData(db *sql.DB, sessionID string) (string, error) {
	var data string
	scanCode := db.QueryRow(GetSessionDataQuery, sessionID).Scan(&data)
	return data, scanCode
}

func QueryUploadSessionData(db *sql.DB, sessionID string, sessionData string) (sql.Result, error) {
	result, err := db.Exec(InsertSessionDataQuery, sessionID, sessionData)
	return result, err
}
