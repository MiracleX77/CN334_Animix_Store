package handlers

import "github.com/labstack/echo/v4"

type TransactionHandler interface {
	GetTransactionById(c echo.Context) error
	GetTransactionAllByOrderId(c echo.Context) error
}
