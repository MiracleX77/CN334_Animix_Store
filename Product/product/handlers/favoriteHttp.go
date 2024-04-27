package handlers

import (
	"strconv"

	productError "github.com/MiracleX77/CN334_Animix_Store/product/errors"
	"github.com/MiracleX77/CN334_Animix_Store/product/models"
	"github.com/MiracleX77/CN334_Animix_Store/product/usecases"

	"github.com/gofiber/fiber/v2/log"
	"github.com/labstack/echo/v4"
)

type favoriteHttpHandler struct {
	favoriteUsecase usecases.FavoriteUsecase
}

func NewFavoriteHttpHandler(favoriteUsecase usecases.FavoriteUsecase) FavoriteHandler {
	return &favoriteHttpHandler{
		favoriteUsecase: favoriteUsecase,
	}
}

func (h *favoriteHttpHandler) GetFavoriteAllByUserId(c echo.Context) error {
	userId := c.Get("userId").(string)
	favoritees, err := h.favoriteUsecase.GetFavoriteAllByUserId(&userId)
	if err != nil {
		log.Errorf("Error getting all Favorite: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", favoritees)
}

func (h *favoriteHttpHandler) InsertFavorite(c echo.Context) error {
	userId := c.Get("userId").(string)
	reqBody := new(models.InsertFavoriteModel)
	parsedUserId, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		log.Errorf("Error parsing userId: %v", err)
		return response(c, 400, "Bad request", nil)
	}
	reqBody.UserId = parsedUserId
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, 400, "Bad request", nil)
	}
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return validationErrorResponse(c, err)
	}
	if err := h.favoriteUsecase.InsertFavorite(reqBody); err != nil {
		log.Errorf("Error inserting favorite: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}

func (h *favoriteHttpHandler) DeleteFavorite(c echo.Context) error {
	favoriteId := c.Param("id")
	if err := h.favoriteUsecase.DeleteFavorite(&favoriteId); err != nil {
		log.Errorf("Error deleting favorite: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}
