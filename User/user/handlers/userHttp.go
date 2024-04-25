package handlers

import (
	userError "github.com/MiracleX77/CN334_Animix_Store/user/errors"
	"github.com/MiracleX77/CN334_Animix_Store/user/models"
	"github.com/MiracleX77/CN334_Animix_Store/user/usecases"

	"github.com/gofiber/fiber/v2/log"
	"github.com/labstack/echo/v4"
)

type userHttpHandler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHttpHandler(userUsecase usecases.UserUsecase) UserHandler {
	return &userHttpHandler{
		userUsecase: userUsecase,
	}
}

func (h *userHttpHandler) GetUserById(c echo.Context) error {
	userId := c.Get("userId").(string)
	if err := h.userUsecase.CheckUserId(&userId); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*userError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	user, err := h.userUsecase.GetUserById(&userId)
	if err != nil {
		log.Errorf("Error getting dentist by id: %v", err)
		if _, ok := err.(*userError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", user)
}

func (h *userHttpHandler) GetUserAll(c echo.Context) error {
	users, err := h.userUsecase.GetUserAll()
	if err != nil {
		log.Errorf("Error getting all dentists: %v", err)
		if _, ok := err.(*userError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", users)
}

func (h *userHttpHandler) UpdateUser(c echo.Context) error {
	userId := c.Get("userId").(string)
	reqBody := new(models.UpdateModel)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, 400, "Bad request", nil)
	}
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return validationErrorResponse(c, err)
	}
	if err := h.userUsecase.CheckUserId(&userId); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*userError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}

	if err := h.userUsecase.UpdateUser(reqBody, &userId); err != nil {
		log.Errorf("Error updating user: %v", err)
		if _, ok := err.(*userError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}

func (h *userHttpHandler) DeleteUser(c echo.Context) error {
	userId := c.Get("userId").(string)
	if err := h.userUsecase.CheckUserId(&userId); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*userError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	if err := h.userUsecase.DeleteUser(&userId); err != nil {
		log.Errorf("Error deleting user: %v", err)
		if _, ok := err.(*userError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}
