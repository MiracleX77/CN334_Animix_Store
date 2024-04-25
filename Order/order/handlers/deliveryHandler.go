package handlers

import "github.com/labstack/echo/v4"

type DeliveryHandler interface {
	UpdateDelivery(c echo.Context) error
	GetDeliveryById(c echo.Context) error
	GetDeliveryAll(c echo.Context) error
	DeleteDelivery(c echo.Context) error
	InsertDelivery(c echo.Context) error
}
