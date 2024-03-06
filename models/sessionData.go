package models

type SessionFormData struct {
	SessionID   string `form:"id"`
	SessionData string `form:"data"`
}
