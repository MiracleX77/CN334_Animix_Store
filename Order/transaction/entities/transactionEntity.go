package entities

import (
	"gorm.io/gorm"

	orderEntity "github.com/MiracleX77/CN334_Animix_Store/order/entities"
)

type (
	Transaction struct {
		gorm.Model
		ProductId int `json:"product_id"`
		OrderId   int `json:"order_id"`
		Order     orderEntity.Order
		Status    string `json:"status"`
	}

	InsertTransaction struct {
		ProductId int    `json:"product_id"`
		OrderId   int    `json:"order_id"`
		Status    string `json:"status"`
	}

	UpdateTransaction struct {
		ProductId int    `json:"product_id"`
		OrderId   int    `json:"order_id"`
		Status    string `json:"status"`
	}
)
