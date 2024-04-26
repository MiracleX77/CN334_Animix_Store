package handlers

import "github.com/labstack/echo/v4"

type OrderHandler interface {
	UpdateOrder(c echo.Context) error
	GetOrderById(c echo.Context) error
	GetOrderByUserId(c echo.Context) error
	GetOrderByStatus(c echo.Context) error
	GetOrderAll(c echo.Context) error
	DeleteOrder(c echo.Context) error
	InsertOrder(c echo.Context) error
}
