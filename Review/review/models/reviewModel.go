package models

import (
	"time"
)

type InsertReviewModel struct {
	UserId    uint64 `json:"user_id" validate:"required"`
	ProductId uint64 `json:"product_id" validate:"required"`
	Title     string `json:"title" validate:"required"`
	Content   string `json:"content" validate:"required"`
	Rating    int    `json:"rating" validate:"required"`
	Status    string `json:"status"`
}

type UpdateReviewModel struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	Rating  int    `json:"rating" validate:"required"`
	Status  string `json:"status"`
}

type ReviewModel struct {
	ID        uint64
	UserId    uint64
	ProductId uint64
	Title     string
	Content   string
	Rating    int
	Polarity  string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
