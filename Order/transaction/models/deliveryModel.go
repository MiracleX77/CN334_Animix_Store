package models

import "time"

type InsertDeliveryModel struct {
	AddressId uint64 `json:"address_id" validate:"required"`
}

type UpdateDeliveryModel struct {
	AddressId      uint64   `json:"address_id" validate:"required"`
	Cost           *float64 `json:"cost"`
	Type           *string  `json:"type"`
	TrackingNumber *string  `json:"tracking_number"`
}

type DeliveryModel struct {
	ID             uint64       `json:"id"`
	Address        AddressModel `json:"address"`
	Cost           *float64     `json:"cost"`
	Type           *string      `json:"type"`
	TrackingNumber *string      `json:"tracking_number"`
	Status         string       `json:"status"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}

type AddressModel struct {
	ID          uint64      `json:"id"`
	UserId      uint64      `json:"user_id"`
	AddressLine *string     `json:"address_line"`
	Phone       string      `json:"phone"`
	Name        string      `json:"name"`
	SubDistrict SubDistrict `json:"sub_district"`
	District    District    `json:"district"`
	Province    Province    `json:"province"`
	Default     string      `json:"default"`
	Status      string      `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}
type SubDistrict struct {
	ID       uint64 `json:"id"`
	NameTh   string `json:"name_th"`
	NameEn   string `json:"name_en"`
	PostCode string `json:"post_code"`
}

type District struct {
	ID     uint64 `json:"id"`
	NameTh string `json:"name_th"`
	NameEn string `json:"name_en"`
}

type Province struct {
	ID     uint64 `json:"id"`
	NameTh string `json:"name_th"`
	NameEn string `json:"name_en"`
}
