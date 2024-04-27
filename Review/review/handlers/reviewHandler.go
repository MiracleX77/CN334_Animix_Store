package handlers

import "github.com/labstack/echo/v4"

type ReviewHandler interface {
	GetReviewById(c echo.Context) error
	GetReviewByUserId(c echo.Context) error
	GetReviewByProductId(c echo.Context) error
	InsertReview(c echo.Context) error
	UpdateReview(c echo.Context) error
	DeleteReview(c echo.Context) error
}
