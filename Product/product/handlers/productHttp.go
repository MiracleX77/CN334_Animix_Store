package handlers

import (
	"strconv"

	productError "github.com/MiracleX77/CN334_Animix_Store/product/errors"
	"github.com/MiracleX77/CN334_Animix_Store/product/models"
	"github.com/MiracleX77/CN334_Animix_Store/product/usecases"

	"github.com/gofiber/fiber/v2/log"
	"github.com/labstack/echo/v4"
)

type productHttpHandler struct {
	productUsecase usecases.ProductUsecase
}

func NewProductHttpHandler(productUsecase usecases.ProductUsecase) ProductHandler {
	return &productHttpHandler{
		productUsecase: productUsecase,
	}
}

func (h *productHttpHandler) GetProductById(c echo.Context) error {
	//userId := c.Get("userId").(string)
	productId := c.Param("id")
	if err := h.productUsecase.CheckProductId(&productId); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	product, err := h.productUsecase.GetProductById(&productId)
	if err != nil {
		log.Errorf("Error getting dentist by id: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", product)
}

func (h *productHttpHandler) GetProductAll(c echo.Context) error {
	productes, err := h.productUsecase.GetProductAll()
	if err != nil {
		log.Errorf("Error getting all Product: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", productes)
}

func (h *productHttpHandler) GetProductAllByCategory(c echo.Context) error {
	categoryId := c.Param("id")
	productes, err := h.productUsecase.GetProductAllByKey("category_id", &categoryId)
	if err != nil {
		log.Errorf("Error getting all Product: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", productes)
}

func (h *productHttpHandler) GetProductAllByName(c echo.Context) error {
	productName := c.Param("name")
	productes, err := h.productUsecase.GetProductAllByKey("name", &productName)
	if err != nil {
		log.Errorf("Error getting all Product: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", productes)
}

func (h *productHttpHandler) InsertProduct(c echo.Context) error {
	reqBody := new(models.InsertProductModel)
	// if err := c.Bind(reqBody); err != nil {
	// 	log.Errorf("Error binding request body: %v", err)
	// 	return response(c, 400, "Bad request", nil)
	// }
	reqBody.AuthorId, _ = strconv.ParseUint(c.FormValue("author_id"), 10, 64)
	reqBody.CategoryId, _ = strconv.ParseUint(c.FormValue("category_id"), 10, 64)
	reqBody.PublisherId, _ = strconv.ParseUint(c.FormValue("publisher_id"), 10, 64)
	reqBody.Name = c.FormValue("name")
	description := c.FormValue("description")
	reqBody.Description = &description
	reqBody.Price, _ = strconv.ParseFloat(c.FormValue("price"), 64)
	reqBody.Stock, _ = strconv.Atoi(c.FormValue("stock"))
	//reqBody.Img = c.FormValue("img")

	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return validationErrorResponse(c, err)
	}

	file, err := c.FormFile("img")
	if err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, 400, "Bad request", nil)
	}

	src, err := file.Open()
	if err != nil {
		log.Errorf("Error opening the file: %v", err)
		return response(c, 500, "Internal Server Error", nil)
	}
	defer src.Close()

	if err := h.productUsecase.SendFileToApi(src, file.Filename); err != nil {
		log.Errorf("Error sending file to external API: %v", err)
		return response(c, 500, "Failed to upload image to external service", nil)
	}

	reqBody.Img = file.Filename
	if err := h.productUsecase.InsertProduct(reqBody); err != nil {
		log.Errorf("Error inserting product: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}

func (h *productHttpHandler) UpdateProduct(c echo.Context) error {
	productId := c.Param("id")
	reqBody := new(models.UpdateProductModel)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, 400, "Bad request", nil)
	}
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return validationErrorResponse(c, err)
	}

	if err := h.productUsecase.CheckProductId(&productId); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}

	if err := h.productUsecase.UpdateProduct(reqBody, &productId); err != nil {
		log.Errorf("Error updating product: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}

func (h *productHttpHandler) DeleteProduct(c echo.Context) error {
	productId := c.Param("id")
	if err := h.productUsecase.CheckProductId(&productId); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	if err := h.productUsecase.DeleteProduct(&productId); err != nil {
		log.Errorf("Error deleting product: %v", err)
		if _, ok := err.(*productError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}
