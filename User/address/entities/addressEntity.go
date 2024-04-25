package entities

import (
	user "github.com/MiracleX77/CN334_Animix_Store/user/entities"
	"gorm.io/gorm"
)

type (
	Address struct {
		gorm.Model
		UserId        int `json:"user_id"`
		User          user.User
		AddressLine   *string `json:"address_line"`
		Phone         string  `json:"phone"`
		Name          string  `json:"name"`
		SubDistrictId int     `json:"sub_district_id"`
		SubDistrict   SubDistrict
		DistrictId    int `json:"district_id"`
		District      District
		ProvinceId    int `json:"province_id"`
		Province      Province
		Default       string `json:"default"`
		Status        string `json:"status"`
	}

	UpdateAddress struct {
		gorm.Model
		UserId        int     `json:"user_id"`
		AddressLine   *string `json:"address_line"`
		Phone         string  `json:"phone"`
		Name          string  `json:"name"`
		SubDistrictId int     `json:"sub_district_id"`
		DistrictId    int     `json:"district_id"`
		ProvinceId    int     `json:"province_id"`
		Default       string  `json:"default"`
		Status        string  `json:"status"`
	}

	InsertAddress struct {
		UserId        int     `json:"user_id"`
		AddressLine   *string `json:"address_line"`
		Phone         string  `json:"phone"`
		Name          string  `json:"name"`
		SubDistrictId int     `json:"sub_district_id"`
		DistrictId    int     `json:"district_id"`
		ProvinceId    int     `json:"province_id"`
		Default       string  `json:"default"`
		Status        string  `json:"status"`
	}
)
