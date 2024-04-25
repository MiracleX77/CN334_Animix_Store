package entities

import "gorm.io/gorm"

type (
	Author struct {
		gorm.Model
		Name        string  `json:"name"`
		Description *string `json:"description"`
		Status      string  `json:"status"`
	}
	InsertAuthor struct {
		Name        string  `json:"name"`
		Description *string `json:"description"`
		Status      string  `json:"status"`
	}
	UpdateAuthor struct {
		Name        string  `json:"name"`
		Description *string `json:"description"`
		Status      string  `json:"status"`
	}
)
