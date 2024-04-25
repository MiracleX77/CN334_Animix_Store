package handlers

import (
	productError "github.com/MiracleX77/CN334_Animix_Store/product/errors"
	"github.com/MiracleX77/CN334_Animix_Store/product/models"
	"github.com/MiracleX77/CN334_Animix_Store/product/usecases"

	"github.com/gofiber/fiber/v2/log"
	"github.com/labstack/echo/v4"
)

type publisherHttpHandler struct {
	publisherUsecase usecases.PublisherUsecase
}

func NewPublisherHttpHandler(publisherUsecase usecases.PublisherUsecase) PublisherHandler {
	return &publisherHttpHandler{
		publisherUsecase: publisherUsecase,
	}
}

func (h *publisherHttpHandler) GetPublisherById(c echo.Context) error {
	publisherId := c.Param("id")
	publisher, err := h.publisherUsecase.GetPublisherById(&publisherId)
	if err != nil {
		log.Errorf("Error getting dentist by id: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", publisher)
}

func (h *publisherHttpHandler) GetPublisherAll(c echo.Context) error {
	publisheres, err := h.publisherUsecase.GetPublisherAll()
	if err != nil {
		log.Errorf("Error getting all Publisher: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", publisheres)
}

func (h *publisherHttpHandler) InsertPublisher(c echo.Context) error {
	reqBody := new(models.InsertPublisherModel)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, 400, "Bad request", nil)
	}
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return validationErrorResponse(c, err)
	}
	if err := h.publisherUsecase.InsertPublisher(reqBody); err != nil {
		log.Errorf("Error inserting publisher: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}

func (h *publisherHttpHandler) UpdatePublisher(c echo.Context) error {
	publisherId := c.Param("id")
	reqBody := new(models.InsertPublisherModel)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, 400, "Bad request", nil)
	}
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return validationErrorResponse(c, err)
	}

	if err := h.publisherUsecase.UpdatePublisher(reqBody, &publisherId); err != nil {
		log.Errorf("Error updating publisher: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}

func (h *publisherHttpHandler) DeletePublisher(c echo.Context) error {
	publisherId := c.Param("id")
	if err := h.publisherUsecase.DeletePublisher(&publisherId); err != nil {
		log.Errorf("Error deleting publisher: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}
