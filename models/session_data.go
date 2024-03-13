package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SessionUploadData struct {
	SessionID   string `form:"id" json:"id"`
	SessionData string `form:"data" json:"data"`
}

type SessionGetData struct {
	SessionID string `form:"id" json:"id"`
}

type Data struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}
