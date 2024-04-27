package entities

import "gorm.io/gorm"

type (
	Favorite struct {
		gorm.Model
		ProductId int `json:"product_id"`
		Product   Product
		UserId    int    `json:"user_id"`
		Status    string `json:"status"`
	}
	InsertFavorite struct {
		ProductId int    `json:"product_id"`
		UserId    int    `json:"user_id"`
		Status    string `json:"status"`
	}
	UpdateFavorite struct {
		Status string `json:"status"`
	}
)
