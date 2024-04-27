package entities

import (
	"gorm.io/gorm"
)

type (
	Review struct {
		gorm.Model
		ProductId int    `json:"product_id"`
		UserId    int    `json:"user_id"`
		Title     string `json:"title"`
		Content   string `json:"content"`
		Rating    int    `json:"rating"`
		Polarity  string `json:"polarity"`
		Status    string `json:"status"`
	}

	InsertReview struct {
		ProductId int    `json:"product_id"`
		UserId    int    `json:"user_id"`
		Title     string `json:"title"`
		Content   string `json:"content"`
		Rating    int    `json:"rating"`
		Polarity  string `json:"polarity"`
		Status    string `json:"status"`
	}

	UpdateReview struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		Rating   int    `json:"rating"`
		Polarity string `json:"polarity"`
		Status   string `json:"status"`
	}
)
