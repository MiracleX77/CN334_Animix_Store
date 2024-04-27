package handlers

import (
	"strconv"

	reviewError "github.com/MiracleX77/CN334_Animix_Store/review/errors"
	"github.com/MiracleX77/CN334_Animix_Store/review/models"
	"github.com/MiracleX77/CN334_Animix_Store/review/usecases"

	"github.com/gofiber/fiber/v2/log"
	"github.com/labstack/echo/v4"
)

type reviewHttpHandler struct {
	reviewUsecase usecases.ReviewUsecase
}

func NewReviewHttpHandler(reviewUsecase usecases.ReviewUsecase) ReviewHandler {
	return &reviewHttpHandler{
		reviewUsecase: reviewUsecase,
	}
}

func (h *reviewHttpHandler) GetReviewById(c echo.Context) error {
	token := c.Get("token").(string)
	reviewId := c.Param("id")
	if err := h.reviewUsecase.CheckReviewId(&reviewId); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*reviewError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	review, err := h.reviewUsecase.GetReviewById(&reviewId, &token)
	if err != nil {
		log.Errorf("Error getting review by id: %v", err)
		if _, ok := err.(*reviewError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", review)
}

func (h *reviewHttpHandler) GetReviewByUserId(c echo.Context) error {
	userId := c.Get("userId").(string)
	reviews, err := h.reviewUsecase.GetReviewByKey("user_id", userId)
	if err != nil {
		log.Errorf("Error getting review by user id: %v", err)
		if _, ok := err.(*reviewError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", reviews)
}

func (h *reviewHttpHandler) GetReviewByProductId(c echo.Context) error {
	reviews, err := h.reviewUsecase.GetReviewByKey("product_id", c.Param("id"))
	if err != nil {
		log.Errorf("Error getting review by status: %v", err)
		if _, ok := err.(*reviewError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", reviews)
}

func (h *reviewHttpHandler) InsertReview(c echo.Context) error {
	userId := c.Get("userId").(string)
	reqBody := new(models.InsertReviewModel)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, 400, "Bad request", nil)
	}
	userIdUint64, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		log.Errorf("Error converting userId to uint64: %v", err)
		return response(c, 400, "Bad request", nil)
	}
	reqBody.UserId = userIdUint64
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return validationErrorResponse(c, err)
	}

	if err := h.reviewUsecase.InsertReview(reqBody); err != nil {
		log.Errorf("Error inserting review: %v", err)
		if _, ok := err.(*reviewError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}

func (h *reviewHttpHandler) UpdateReview(c echo.Context) error {
	reviewId := c.Param("id")
	reqBody := new(models.UpdateReviewModel)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, 400, "Bad request", nil)
	}
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return validationErrorResponse(c, err)
	}
	if err := h.reviewUsecase.CheckReviewId(&reviewId); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*reviewError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}

	if err := h.reviewUsecase.UpdateReview(reqBody, &reviewId); err != nil {
		log.Errorf("Error updating review: %v", err)
		if _, ok := err.(*reviewError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}

func (h *reviewHttpHandler) DeleteReview(c echo.Context) error {
	reviewId := c.Param("id")
	if err := h.reviewUsecase.CheckReviewId(&reviewId); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*reviewError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	if err := h.reviewUsecase.DeleteReview(&reviewId); err != nil {
		log.Errorf("Error deleting review: %v", err)
		if _, ok := err.(*reviewError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}
