package entities

import (
	"gorm.io/gorm"
)

type (
	Delivery struct {
		gorm.Model
		AddressId      int      `json:"author_id"`
		Cost           *float64 `json:"cost"`
		Type           *string  `json:"type"`
		TrackingNumber *string  `json:"tracking_number"`
		Status         string   `json:"status"`
	}

	InsertDelivery struct {
		AddressId int    `json:"author_id" `
		Status    string `json:"status"`
	}

	UpdateDelivery struct {
		AddressId      int      `json:"author_id"`
		Cost           *float64 `json:"cost"`
		Type           *string  `json:"type"`
		TrackingNumber *string  `json:"tracking_number"`
		Status         string   `json:"status"`
	}
)
