package models

import "time"

type TaskHistory struct {
	Id            uint       `json:"id" bson:"id" gorm:"primary_key"`
	TaskHistoryId string     `json:"taskHistoryId" bson:"taskHistoryId"`
	TaskId        string     `json:"taskId" bson:"taskId"`
	TaskTitle     string     `json:"taskTitle" bson:"taskTitle"`
	Tags          string     `json:"tags" bson:"tags"`
	OldStatus     string     `json:"OldStatus" bson:"OldStatus"`
	NewStatus     string     `json:"NewStatus" bson:"NewStatus"`
	BoardId       string     `json:"boardId" bson:"boardId"`
	Week          int        `json:"week" bson:"week"`
	Priority      *int       `json:"priority" bson:"priority"`
	Dificulty     *int       `json:"dificulty" bson:"dificulty"`
	PercentCompleted *float64  `json:"percentCompleted" bson:"percentCompleted"`
	Type          *string    `json:"type" bson:"type"`
	CreatedAt     *time.Time `json:"createdAt" bson:"createdAt"`
	EditedAt      *time.Time `json:"editedAt" bson:"editedAt"`
	DueDate       *time.Time `json:"dueDate" bson:"dueDate"`
	CreatedBy     string     `json:"createdBy" bson:"createdBy"`
}

type CreateTaskHistory struct {
	TaskHistoryId string     `json:"taskHistoryId" bson:"taskHistoryId"`
	TaskId        string     `json:"taskId" bson:"taskId"`
	TaskTitle     string     `json:"taskTitle" bson:"taskTitle"`
	Tags          string     `json:"tags" bson:"tags"`
	OldStatus     string     `json:"OldStatus" bson:"OldStatus"`
	NewStatus     string     `json:"NewStatus" bson:"NewStatus"`
	BoardId       string     `json:"boardId" bson:"boardId"`
	Priority      *int       `json:"priority" bson:"priority"`
	Dificulty     *int       `json:"dificulty" bson:"dificulty"`
	PercentCompleted *float64  `json:"percentCompleted" bson:"percentCompleted"`
	Type          *string    `json:"type" bson:"type"`
	DueDate       *time.Time `json:"dueDate" bson:"dueDate"`
	CreatedBy     string     `json:"createdBy" bson:"createdBy"`
}

type TaskHistoryResponse struct {
	message string        `json:"message" bson:"message"`
	Data    []TaskHistory `json:"data" bson:"data"`
}
