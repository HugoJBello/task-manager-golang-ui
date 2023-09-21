package models

import "time"

type PointsReport struct {
	Week      int        `json:"week" bson:"week"`
	Points    int        `json:"points" bson:"points"`
	CreatedAt *time.Time `json:"editedAt" bson:"editedAt"`
	BoardId   string     `json:"boardId" bson:"boardId"`
}

type PointsReportResponse struct {
	message string         `json:"message" bson:"message"`
	Data    []PointsReport `json:"data" bson:"data"`
}
