package models

import "time"

type ServiceActivity struct {
	Id uint `json:"id" bson:"id" gorm:"primary_key"`
	ServiceId string `json:"serviceId" bson:"serviceId"`
	ServiceName string `json:"serviceName" bson:"serviceName"`
	ActivityContent string `json:"activityContent" bson:"activityContent"`
	ActivityContentType string `json:"activityContentType" bson:"activityContentType"`
	Date time.Time `json:"date" bson:"date"`

}

type ServiceActivityResponse struct {
	message string `json:"message" bson:"message"`
	Data []ServiceActivity `json:"data" bson:"data"`

}

type CreateServiceActivity struct {
	ServiceId string `json:"serviceId" bson:"serviceId" binding:"required"`
	ServiceName string `json:"serviceName" bson:"serviceName" binding:"required"`
	ActivityContent string `json:"activityContent" bson:"activityContent" binding:"required"`
	ActivityContentType string `json:"activityContentType" bson:"activityContentType" binding:"required"`
	Date time.Time `json:"date" bson:"date" binding:"required"`
}