package models

type SessionData struct {
	SessionID   string `form:"id" json:"id"`
	SessionData string `form:"data" json:"data"`
}

type SessionGetData struct {
	SessionID string `form:"id" json:"id"`
}
