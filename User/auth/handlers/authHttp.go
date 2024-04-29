package handlers

import (
	"net/http"

	authError "github.com/MiracleX77/CN334_Animix_Store/auth/errors"
	"github.com/MiracleX77/CN334_Animix_Store/auth/models"
	"github.com/MiracleX77/CN334_Animix_Store/auth/usecases"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type authHttpHandler struct {
	authUsecase usecases.AuthUsecase
}

func NewAuthHttpHandler(authUsecase usecases.AuthUsecase) AuthHandler {
	return &authHttpHandler{
		authUsecase: authUsecase,
	}
}

func (h *authHttpHandler) Register(c echo.Context) error {
	reqBody := new(models.RegisterData)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, http.StatusBadRequest, "Bad request", nil)
	}
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return response(c, http.StatusBadRequest, "Bad request", nil)
	}
	if err := h.authUsecase.CheckData(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*authError.ServerInternalError); ok {
			return response(c, http.StatusInternalServerError, "Server Internal Error", nil)
		} else {
			return response(c, http.StatusBadRequest, err.Error(), nil)
		}
	}
	if err := h.authUsecase.RegisterDataProcessing(reqBody); err != nil {
		return response(c, http.StatusInternalServerError, "Processing Data failed", nil)
	}

	return response(c, http.StatusOK, "Success", nil)
}

func (h *authHttpHandler) Login(c echo.Context) error {
	reqBody := new(models.LoginData)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, http.StatusBadRequest, "Bad request", nil)
	}
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return response(c, http.StatusBadRequest, "Bad request", nil)
	}

	token, err := h.authUsecase.LoginDataProcession(reqBody)
	if err != nil {
		log.Errorf("Error login: %v", err)
		if _, ok := err.(*authError.ServerInternalError); ok {
			return response(c, http.StatusInternalServerError, "Server Internal Error", nil)
		} else {
			return response(c, http.StatusBadRequest, err.Error(), nil)
		}
	}
	return response(c, http.StatusOK, "Success", token)

}
