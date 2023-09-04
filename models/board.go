package models

import "time"

type Board struct {
	Id         uint       `json:"id" bson:"id" gorm:"primary_key"`
	BoardId    string     `json:"boardId"    bson:"boardId"`
	BoardTitle string     `json:"boardTitle" bson:"boardTitle"`
	BoardBody  string     `json:"boardDescription"  bson:"boardDescription"`
	Tags       string     `json:"tags" bson:"tags"`
	Status     string     `json:"status" bson:"status"`
	CreatedAt  *time.Time `json:"createdAt" bson:"createdAt"`
	EditedAt   *time.Time `json:"editedAt" bson:"editedAt"`
	CreatedBy  string     `json:"createdBy" bson:"createdBy"`
}

type CreateBoard struct {
	BoardId    string     `json:"boardId"    bson:"boardId"`
	BoardTitle string     `json:"boardTitle" bson:"boardTitle"`
	BoardBody  string     `json:"boardDescription"  bson:"boardDescription"`
	Tags       string     `json:"tags" bson:"tags"`
	Status     string     `json:"status" bson:"status"`
	CreatedAt  *time.Time `json:"createdAt" bson:"createdAt"`
	EditedAt   *time.Time `json:"editedAt" bson:"editedAt"`
	CreatedBy  string     `json:"createdBy" bson:"createdBy"`
}

type BoardResponse struct {
	message string `json:"message" bson:"message"`
	Data    []Board `json:"data" bson:"data"`
}
