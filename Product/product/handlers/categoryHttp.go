package handlers

import (
	productError "github.com/MiracleX77/CN334_Animix_Store/product/errors"
	"github.com/MiracleX77/CN334_Animix_Store/product/models"
	"github.com/MiracleX77/CN334_Animix_Store/product/usecases"

	"github.com/gofiber/fiber/v2/log"
	"github.com/labstack/echo/v4"
)

type categoryHttpHandler struct {
	categoryUsecase usecases.CategoryUsecase
}

func NewCategoryHttpHandler(categoryUsecase usecases.CategoryUsecase) CategoryHandler {
	return &categoryHttpHandler{
		categoryUsecase: categoryUsecase,
	}
}

func (h *categoryHttpHandler) GetCategoryById(c echo.Context) error {
	categoryId := c.Param("id")
	category, err := h.categoryUsecase.GetCategoryById(&categoryId)
	if err != nil {
		log.Errorf("Error getting dentist by id: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", category)
}

func (h *categoryHttpHandler) GetCategoryAll(c echo.Context) error {
	categoryes, err := h.categoryUsecase.GetCategoryAll()
	if err != nil {
		log.Errorf("Error getting all Category: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", categoryes)
}

func (h *categoryHttpHandler) InsertCategory(c echo.Context) error {
	reqBody := new(models.InsertCategoryModel)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, 400, "Bad request", nil)
	}
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return validationErrorResponse(c, err)
	}
	if err := h.categoryUsecase.InsertCategory(reqBody); err != nil {
		log.Errorf("Error inserting category: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}

func (h *categoryHttpHandler) UpdateCategory(c echo.Context) error {
	categoryId := c.Param("id")
	reqBody := new(models.InsertCategoryModel)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, 400, "Bad request", nil)
	}
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return validationErrorResponse(c, err)
	}

	if err := h.categoryUsecase.UpdateCategory(reqBody, &categoryId); err != nil {
		log.Errorf("Error updating category: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}

func (h *categoryHttpHandler) DeleteCategory(c echo.Context) error {
	categoryId := c.Param("id")
	if err := h.categoryUsecase.DeleteCategory(&categoryId); err != nil {
		log.Errorf("Error deleting category: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}
