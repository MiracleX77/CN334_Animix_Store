package entities

import (
	delivery "github.com/MiracleX77/CN334_Animix_Store/delivery/entities"
	payment "github.com/MiracleX77/CN334_Animix_Store/payment/entities"
	"gorm.io/gorm"
)

type (
	Order struct {
		gorm.Model
		UserId     int `json:"user_id"`
		DeliveryId int `json:"delivery_id"`
		Delivery   delivery.Delivery
		PaymentId  int `json:"payment_id"`
		Payment    payment.Payment
		TotalPrice float64 `json:"total_price"`
		Status     string  `json:"status"`
	}

	InsertOrder struct {
		UserId     int     `json:"user_id"`
		DeliveryId int     `json:"delivery_id"`
		PaymentId  int     `json:"payment_id"`
		TotalPrice float64 `json:"total_price"`
		Status     string  `json:"status"`
	}

	UpdateOrder struct {
		Status string `json:"status"`
	}
)
