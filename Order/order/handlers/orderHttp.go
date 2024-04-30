package handlers

import (
	"fmt"
	"strconv"
	"strings"

	orderError "github.com/MiracleX77/CN334_Animix_Store/order/errors"
	"github.com/MiracleX77/CN334_Animix_Store/order/models"
	"github.com/MiracleX77/CN334_Animix_Store/order/usecases"

	"github.com/gofiber/fiber/v2/log"
	"github.com/labstack/echo/v4"
)

type orderHttpHandler struct {
	orderUsecase usecases.OrderUsecase
}

func NewOrderHttpHandler(orderUsecase usecases.OrderUsecase) OrderHandler {
	return &orderHttpHandler{
		orderUsecase: orderUsecase,
	}
}

func (h *orderHttpHandler) GetOrderById(c echo.Context) error {
	token := c.Get("token").(string)
	orderId := c.Param("id")
	if err := h.orderUsecase.CheckOrderId(&orderId); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*orderError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	order, err := h.orderUsecase.GetOrderById(&orderId, &token)
	if err != nil {
		log.Errorf("Error getting order by id: %v", err)
		if _, ok := err.(*orderError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", order)
}

func (h *orderHttpHandler) GetOrderByUserId(c echo.Context) error {
	orders, err := h.orderUsecase.GetOrderByKey("user_id", c.Param("id"))
	if err != nil {
		log.Errorf("Error getting order by user id: %v", err)
		if _, ok := err.(*orderError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", orders)
}

func (h *orderHttpHandler) GetOrderByStatus(c echo.Context) error {
	orders, err := h.orderUsecase.GetOrderByKey("status", c.Param("status"))
	if err != nil {
		log.Errorf("Error getting order by status: %v", err)
		if _, ok := err.(*orderError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", orders)
}

func (h *orderHttpHandler) GetOrderAll(c echo.Context) error {
	orders, err := h.orderUsecase.GetOrderAll()
	if err != nil {
		log.Errorf("Error getting all Order: %v", err)
		if _, ok := err.(*orderError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", orders)
}

func (h *orderHttpHandler) InsertOrder(c echo.Context) error {

	userId := c.Get("userId").(string)
	if err := c.Request().ParseMultipartForm(1024); err != nil {
		return response(c, 400, "Failed to parse form", nil)
	}

	// สมมติว่าคุณต้องการรับ array ของ product ID จากฟิลด์ 'list_product_id[]'
	listProductIds := []uint64{}
	listProductStrs := c.Request().PostForm["list_product_id"] // ใช้ PostForm เพื่อรับ slice ของค่า
	fmt.Println(listProductStrs)
	for _, strID := range listProductStrs {
		ids := strings.Split(strID, ",")
		for _, strID := range ids {
			id, err := strconv.ParseUint(strID, 10, 64)
			if err != nil {
				return response(c, 400, "Invalid product ID", nil)
			}
			listProductIds = append(listProductIds, id)
		}
	}
	reqBody := new(models.InsertOrderModel)
	reqBody.UserId, _ = strconv.ParseUint(userId, 10, 64)
	reqBody.AddressId, _ = strconv.ParseUint(c.FormValue("address_id"), 10, 64)
	reqBody.Type = c.FormValue("type")
	reqBody.Total, _ = strconv.ParseFloat(c.FormValue("total"), 64)
	reqBody.ListProductId = listProductIds

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

	if err := h.orderUsecase.SendFileToApi(src, file.Filename); err != nil {
		log.Errorf("Error sending file to external API: %v", err)
		return response(c, 500, "Failed to upload image to external service", nil)
	}

	reqBody.Img = file.Filename

	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return validationErrorResponse(c, err)
	}

	if err := h.orderUsecase.InsertOrder(reqBody); err != nil {
		log.Errorf("Error inserting order: %v", err)
		if _, ok := err.(*orderError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}

func (h *orderHttpHandler) UpdateOrder(c echo.Context) error {
	orderId := c.Param("id")
	reqBody := new(models.UpdateOrderModel)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, 400, "Bad request", nil)
	}
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return validationErrorResponse(c, err)
	}
	if err := h.orderUsecase.CheckOrderId(&orderId); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*orderError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}

	if err := h.orderUsecase.UpdateStatusOrder(reqBody, &orderId); err != nil {
		log.Errorf("Error updating order: %v", err)
		if _, ok := err.(*orderError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}

func (h *orderHttpHandler) DeleteOrder(c echo.Context) error {
	orderId := c.Param("id")
	if err := h.orderUsecase.CheckOrderId(&orderId); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*orderError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	if err := h.orderUsecase.DeleteOrder(&orderId); err != nil {
		log.Errorf("Error deleting order: %v", err)
		if _, ok := err.(*orderError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}
