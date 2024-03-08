package models

type SessionUploadData struct {
	SessionID   string `form:"id"`
	SessionData string `form:"data"`
}

type SessionGetData struct {
	SessionID string `form:"id"`
}
