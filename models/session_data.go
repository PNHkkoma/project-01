package models

type SessionUploadData struct {
	SessionID   string `form:"id" json:"id"`
	SessionData string `form:"data" json:"data"`
}

type SessionGetData struct {
	SessionID string `form:"id" json:"id"`
}
