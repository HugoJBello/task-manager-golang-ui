package models

import "time"

type Task struct {
	Id        uint       `json:"id" bson:"id" gorm:"primary_key"`
	TaskId    string     `json:"taskId" bson:"taskId"`
	TaskTitle string     `json:"taskTitle" bson:"taskTitle"`
	TaskBody  string     `json:"taskBody" bson:"taskBody"`
	Tags      string     `json:"tags" bson:"tags"`
	Status    string     `json:"status" bson:"status"`
	BoardId   string     `json:"boardId" bson:"boardId"`
	CreatedAt *time.Time `json:"createdAt" bson:"createdAt"`
	EditedAt  *time.Time `json:"editedAt" bson:"editedAt"`
	DueDate   *time.Time `json:"dueDate" bson:"dueDate"`
	CreatedBy string     `json:"createdBy" bson:"createdBy"`
}

type CreateTask struct {
	TaskId    string     `json:"taskId" bson:"taskId"`
	TaskTitle string     `json:"taskTitle" bson:"taskTitle"`
	TaskBody  string     `json:"taskBody" bson:"taskBody"`
	Tags      string     `json:"tags" bson:"tags"`
	Status    string     `json:"status" bson:"status"`
	BoardId   string     `json:"boardId" bson:"boardId"`
	DueDate   *time.Time `json:"dueDate" bson:"dueDate"`
	CreatedBy string     `json:"createdBy" bson:"createdBy"`
}

type TaskResponse struct {
	message string `json:"message" bson:"message"`
	Data    []Task `json:"data" bson:"data"`
}
