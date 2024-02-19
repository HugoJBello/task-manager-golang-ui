
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
	Priority  *int       `json:"priority" bson:"priority"`
	Dificulty *int       `json:"dificulty" bson:"dificulty"`
	PercentCompleted *float64       `json:"percentCompleted" bson:"percentCompleted"`
	Type      *string    `json:"type" bson:"type"`
	Archived  *string    `json:"archived" bson:"archived"`
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
	Priority  *int       `json:"priority" bson:"priority"`
	Dificulty *int       `json:"dificulty" bson:"dificulty"`
	PercentCompleted *float64       `json:"percentCompleted" bson:"percentCompleted"`
	Type      *string    `json:"type" bson:"type"`
	DueDate   *time.Time `json:"dueDate" bson:"dueDate"`
	CreatedBy string     `json:"createdBy" bson:"createdBy"`
	Archived  *string    `json:"archived" bson:"archived"`
}

type TaskResponse struct {
	message string `json:"message" bson:"message"`
	Data    []Task `json:"data" bson:"data"`
}
