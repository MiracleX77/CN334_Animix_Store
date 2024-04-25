package handlers

import (
	paymentError "github.com/MiracleX77/CN334_Animix_Store/payment/errors"
	"github.com/MiracleX77/CN334_Animix_Store/payment/usecases"

	"github.com/gofiber/fiber/v2/log"
	"github.com/labstack/echo/v4"
)

type paymentHttpHandler struct {
	paymentUsecase usecases.PaymentUsecase
}

func NewPaymentHttpHandler(paymentUsecase usecases.PaymentUsecase) PaymentHandler {
	return &paymentHttpHandler{
		paymentUsecase: paymentUsecase,
	}
}

func (h *paymentHttpHandler) GetPaymentById(c echo.Context) error {
	paymentId := c.Param("id")
	if err := h.paymentUsecase.CheckPaymentId(&paymentId); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*paymentError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	payment, err := h.paymentUsecase.GetPaymentById(&paymentId)
	if err != nil {
		log.Errorf("Error getting payment by id: %v", err)
		if _, ok := err.(*paymentError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", payment)
}
