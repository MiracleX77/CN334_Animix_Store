package handlers

import (
	transactionError "github.com/MiracleX77/CN334_Animix_Store/transaction/errors"
	"github.com/MiracleX77/CN334_Animix_Store/transaction/usecases"

	"github.com/gofiber/fiber/v2/log"
	"github.com/labstack/echo/v4"
)

type transactionHttpHandler struct {
	transactionUsecase usecases.TransactionUsecase
}

func NewTransactionHttpHandler(transactionUsecase usecases.TransactionUsecase) TransactionHandler {
	return &transactionHttpHandler{
		transactionUsecase: transactionUsecase,
	}
}

func (h *transactionHttpHandler) GetTransactionById(c echo.Context) error {
	token := c.Get("token").(string)
	transactionId := c.Param("id")

	transaction, err := h.transactionUsecase.GetTransactionById(&transactionId, &token)
	if err != nil {
		log.Errorf("Error getting transaction by id: %v", err)
		if _, ok := err.(*transactionError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", transaction)
}

func (h *transactionHttpHandler) GetTransactionAllByOrderId(c echo.Context) error {
	token := c.Get("token").(string)
	orderId := c.Param("id")
	transactions, err := h.transactionUsecase.GetTransactionAllByOrderId(&orderId, &token)
	if err != nil {
		log.Errorf("Error getting all Transaction: %v", err)
		if _, ok := err.(*transactionError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", transactions)
}
