package models

import "time"

type InsertAddressModel struct {
	AddressLine   string `json:"address_line" `
	Phone         string `json:"phone" validate:"required"`
	Name          string `json:"name" validate:"required"`
	SubDistrictId uint64 `json:"sub_district_id" validate:"required"`
	DistrictId    uint64 `json:"district_id" validate:"required"`
	ProvinceId    uint64 `json:"province_id" validate:"required"`
	Default       string `json:"default" validate:"required"`
}

type UpdateAddressModel struct {
	ID            uint64 `json:"id" validate:"required" `
	AddressLine   string `json:"address_line" `
	Phone         string `json:"phone" validate:"required"`
	Name          string `json:"name" validate:"required"`
	SubDistrictId uint64 `json:"sub_district_id" validate:"required"`
	DistrictId    uint64 `json:"district_id" validate:"required"`
	ProvinceId    uint64 `json:"province_id" validate:"required"`
	Default       string `json:"default" validate:"required"`
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
