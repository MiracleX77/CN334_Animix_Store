package models

import (
	"time"

	deliveryModel "github.com/MiracleX77/CN334_Animix_Store/delivery/models"
	paymentModel "github.com/MiracleX77/CN334_Animix_Store/payment/models"
)

type InsertOrderModel struct {
	UserId        uint64   `json:"user" validate:"required"`
	AddressId     uint64   `json:"address_id" validate:"required"`
	ListProductId []uint64 `json:"list_product_id" validate:"required"`
	Type          string   `json:"type" validate:"required"`
	Img           string   `json:"img" validate:"required"`
	Total         float64  `json:"total" validate:"required"`
}

type UpdateOrderModel struct {
	AddressId     uint64   `json:"address_id" validate:"required"`
	ListProductId []uint64 `json:"list_product_id" validate:"required"`
	Type          string   `json:"type" validate:"required"`
	Img           string   `json:"img" validate:"required"`
	Total         float64  `json:"total" validate:"required"`
}

type OrderModel struct {
	ID         uint64                      `json:"id"`
	UserId     uint64                      `json:"user_id"`
	Delivery   deliveryModel.DeliveryModel `json:"delivery"`
	Payment    paymentModel.PaymentModel   `json:"payment"`
	TotalPrice float64                     `json:"total_price"`
	Status     string                      `json:"status"`
	CreatedAt  time.Time                   `json:"created_at"`
	UpdatedAt  time.Time                   `json:"updated_at"`
}

type ListOrderModel struct {
	ID         uint64    `json:"id"`
	UserId     uint64    `json:"user_id"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
