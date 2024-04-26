package models

import "time"

type InsertTransactionModel struct {
	AddressId uint64 `json:"address_id" validate:"required"`
}

type UpdateTransactionModel struct {
	AddressId      uint64   `json:"address_id" validate:"required"`
	Cost           *float64 `json:"cost"`
	Type           *string  `json:"type"`
	TrackingNumber *string  `json:"tracking_number"`
}

type TransactionModel struct {
	ID        uint64       `json:"id"`
	Product   ProductModel `json:"product"`
	OrderId   uint64       `json:"order_id"`
	Price     float64      `json:"price"`
	Status    string       `json:"status"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type ProductModel struct {
	ID          uint64         `json:"id"`
	Author      AuthorModel    `json:"author"`
	Category    CategoryModel  `json:"category"`
	Publisher   PublisherModel `json:"publisher"`
	Name        string         `json:"name"`
	Description *string        `json:"description"`
	Price       float64        `json:"price"`
	Stock       int            `json:"stock"`
	ImgUrl      string         `json:"img_url"`
	Status      string         `json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type AuthorModel struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
type CategoryModel struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
type PublisherModel struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
