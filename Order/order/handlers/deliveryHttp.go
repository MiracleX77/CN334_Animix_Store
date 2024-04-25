package handlers

import (
	deliveryError "github.com/MiracleX77/CN334_Animix_Store/delivery/errors"
	"github.com/MiracleX77/CN334_Animix_Store/delivery/models"
	"github.com/MiracleX77/CN334_Animix_Store/delivery/usecases"

	"github.com/gofiber/fiber/v2/log"
	"github.com/labstack/echo/v4"
)

type deliveryHttpHandler struct {
	deliveryUsecase usecases.DeliveryUsecase
}

func NewDeliveryHttpHandler(deliveryUsecase usecases.DeliveryUsecase) DeliveryHandler {
	return &deliveryHttpHandler{
		deliveryUsecase: deliveryUsecase,
	}
}

func (h *deliveryHttpHandler) GetDeliveryById(c echo.Context) error {
	token := c.Get("token").(string)
	deliveryId := c.Param("id")
	if err := h.deliveryUsecase.CheckDeliveryId(&deliveryId); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*deliveryError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	delivery, err := h.deliveryUsecase.GetDeliveryById(&deliveryId, &token)
	if err != nil {
		log.Errorf("Error getting delivery by id: %v", err)
		if _, ok := err.(*deliveryError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", delivery)
}

func (h *deliveryHttpHandler) GetDeliveryAll(c echo.Context) error {
	token := c.Get("token").(string)
	deliveryes, err := h.deliveryUsecase.GetDeliveryAll(&token)
	if err != nil {
		log.Errorf("Error getting all Delivery: %v", err)
		if _, ok := err.(*deliveryError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", deliveryes)
}

func (h *deliveryHttpHandler) InsertDelivery(c echo.Context) error {
	reqBody := new(models.InsertDeliveryModel)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, 400, "Bad request", nil)
	}

	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return validationErrorResponse(c, err)
	}

	if err := h.deliveryUsecase.InsertDelivery(reqBody); err != nil {
		log.Errorf("Error inserting delivery: %v", err)
		if _, ok := err.(*deliveryError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}

func (h *deliveryHttpHandler) UpdateDelivery(c echo.Context) error {
	deliveryId := c.Param("id")
	reqBody := new(models.UpdateDeliveryModel)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, 400, "Bad request", nil)
	}
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return validationErrorResponse(c, err)
	}
	if err := h.deliveryUsecase.CheckDeliveryId(&deliveryId); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*deliveryError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}

	if err := h.deliveryUsecase.UpdateDelivery(reqBody, &deliveryId); err != nil {
		log.Errorf("Error updating delivery: %v", err)
		if _, ok := err.(*deliveryError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}

func (h *deliveryHttpHandler) DeleteDelivery(c echo.Context) error {
	deliveryId := c.Param("id")
	if err := h.deliveryUsecase.CheckDeliveryId(&deliveryId); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*deliveryError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	if err := h.deliveryUsecase.DeleteDelivery(&deliveryId); err != nil {
		log.Errorf("Error deleting delivery: %v", err)
		if _, ok := err.(*deliveryError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}
