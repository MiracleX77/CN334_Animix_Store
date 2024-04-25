package handlers

import "github.com/labstack/echo/v4"

type UserHandler interface {
	UpdateUser(c echo.Context) error
	GetUserById(c echo.Context) error
	GetUserAll(c echo.Context) error
	DeleteUser(c echo.Context) error
}
