package handlers

import "github.com/labstack/echo/v4"

type PaymentHandler interface {
	GetPaymentById(c echo.Context) error
}
