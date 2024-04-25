package handlers

import (
	productError "github.com/MiracleX77/CN334_Animix_Store/product/errors"
	"github.com/MiracleX77/CN334_Animix_Store/product/models"
	"github.com/MiracleX77/CN334_Animix_Store/product/usecases"

	"github.com/gofiber/fiber/v2/log"
	"github.com/labstack/echo/v4"
)

type authorHttpHandler struct {
	authorUsecase usecases.AuthorUsecase
}

func NewAuthorHttpHandler(authorUsecase usecases.AuthorUsecase) AuthorHandler {
	return &authorHttpHandler{
		authorUsecase: authorUsecase,
	}
}

func (h *authorHttpHandler) GetAuthorById(c echo.Context) error {
	authorId := c.Param("id")
	author, err := h.authorUsecase.GetAuthorById(&authorId)
	if err != nil {
		log.Errorf("Error getting dentist by id: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", author)
}

func (h *authorHttpHandler) GetAuthorAll(c echo.Context) error {
	authores, err := h.authorUsecase.GetAuthorAll()
	if err != nil {
		log.Errorf("Error getting all Author: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", authores)
}

func (h *authorHttpHandler) InsertAuthor(c echo.Context) error {
	reqBody := new(models.InsertAuthorModel)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, 400, "Bad request", nil)
	}
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return validationErrorResponse(c, err)
	}
	if err := h.authorUsecase.InsertAuthor(reqBody); err != nil {
		log.Errorf("Error inserting author: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}

func (h *authorHttpHandler) UpdateAuthor(c echo.Context) error {
	authorId := c.Param("id")
	reqBody := new(models.InsertAuthorModel)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, 400, "Bad request", nil)
	}
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return validationErrorResponse(c, err)
	}

	if err := h.authorUsecase.UpdateAuthor(reqBody, &authorId); err != nil {
		log.Errorf("Error updating author: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}

func (h *authorHttpHandler) DeleteAuthor(c echo.Context) error {
	authorId := c.Param("id")
	if err := h.authorUsecase.DeleteAuthor(&authorId); err != nil {
		log.Errorf("Error deleting author: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}
