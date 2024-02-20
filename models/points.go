package models

import "time"

type PointsReport struct {
	Week               int           `json:"week" bson:"week"`
	Points             float64       `json:"points" bson:"points"`
	TasksSummingPoints []TaskHistory `json:"tasksSummingPoints" bson:"tasksSummingPoints"`
	CreatedAt          *time.Time    `json:"editedAt" bson:"editedAt"`
	BoardId            string        `json:"boardId" bson:"boardId"`
}

type PointsReportResponse struct {
	message string         `json:"message" bson:"message"`
	Data    []PointsReport `json:"data" bson:"data"`
}
