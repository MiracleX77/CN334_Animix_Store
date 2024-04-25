package entities

import "gorm.io/gorm"

type (
	Publisher struct {
		gorm.Model
		Name   string `json:"name"`
		Status string `json:"status"`
	}
	InsertPublisher struct {
		Name   string `json:"name"`
		Status string `json:"status"`
	}
	UpdatePublisher struct {
		Name   string `json:"name"`
		Status string `json:"status"`
	}
)
