package mysql

import (
	"database/sql"
)

func QueryGetSessionData(db *sql.DB, sessionID string, sessionData *string) error {
	scanCode := db.QueryRow(GetSessionDataQuery, sessionID).Scan(sessionData)
	return scanCode
}

func QueryUploadSessionData(db *sql.DB, sessionID string, sessionData string) (sql.Result, error) {
	result, err := db.Exec(InsertSessionDataQuery, sessionID, sessionData)
	return result, err
}
