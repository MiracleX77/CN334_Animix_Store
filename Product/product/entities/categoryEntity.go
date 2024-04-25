package entities

import "gorm.io/gorm"

type (
	Category struct {
		gorm.Model
		Name   string `json:"name"`
		Status string `json:"status"`
	}
	InsertCategory struct {
		Name   string `json:"name"`
		Status string `json:"status"`
	}
	UpdateCategory struct {
		Name   string `json:"name"`
		Status string `json:"status"`
	}
)
