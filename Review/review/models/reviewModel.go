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
	ID        uint64    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	ProductId uint64    `json:"product_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Rating    int       `json:"rating"`
	Polarity  string    `json:"polarity"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserModel struct {
	ID        uint64    `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_on"`
	Status    string    `json:"status"`
}
