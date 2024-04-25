package handlers

import "github.com/labstack/echo/v4"

type AddressHandler interface {
	UpdateAddress(c echo.Context) error
	GetAddressById(c echo.Context) error
	GetAddressAll(c echo.Context) error
	DeleteAddress(c echo.Context) error
	InsertAddress(c echo.Context) error
	GetProvince(c echo.Context) error
	GetDistrictByProvinceId(c echo.Context) error
	GetSubDistrictByDistrictId(c echo.Context) error
}
