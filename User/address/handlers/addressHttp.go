package handlers

import (
	addressError "github.com/MiracleX77/CN334_Animix_Store/address/errors"
	"github.com/MiracleX77/CN334_Animix_Store/address/models"
	"github.com/MiracleX77/CN334_Animix_Store/address/usecases"

	"github.com/gofiber/fiber/v2/log"
	"github.com/labstack/echo/v4"
)

type addressHttpHandler struct {
	addressUsecase usecases.AddressUsecase
}

func NewAddressHttpHandler(addressUsecase usecases.AddressUsecase) AddressHandler {
	return &addressHttpHandler{
		addressUsecase: addressUsecase,
	}
}

func (h *addressHttpHandler) GetAddressById(c echo.Context) error {
	//userId := c.Get("userId").(string)
	addressId := c.Param("id")
	if err := h.addressUsecase.CheckAddressId(&addressId); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*addressError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	address, err := h.addressUsecase.GetAddressById(&addressId)
	if err != nil {
		log.Errorf("Error getting dentist by id: %v", err)
		if _, ok := err.(*addressError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", address)
}

func (h *addressHttpHandler) GetAddressAll(c echo.Context) error {
	userId := c.Get("userId").(string)
	addresses, err := h.addressUsecase.GetAddressAll(&userId)
	if err != nil {
		log.Errorf("Error getting all dentists: %v", err)
		if _, ok := err.(*addressError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", addresses)
}

func (h *addressHttpHandler) InsertAddress(c echo.Context) error {
	reqBody := new(models.InsertAddressModel)
	userId := c.Get("userId").(string)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, 400, "Bad request", nil)
	}
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return validationErrorResponse(c, err)
	}
	if err := h.addressUsecase.InsertAddress(reqBody, &userId); err != nil {
		log.Errorf("Error inserting address: %v", err)
		if _, ok := err.(*addressError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}

func (h *addressHttpHandler) UpdateAddress(c echo.Context) error {
	addressId := c.Param("id")
	reqBody := new(models.UpdateAddressModel)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, 400, "Bad request", nil)
	}
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return validationErrorResponse(c, err)
	}
	if err := h.addressUsecase.CheckAddressId(&addressId); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*addressError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}

	if err := h.addressUsecase.UpdateAddress(reqBody, &addressId); err != nil {
		log.Errorf("Error updating address: %v", err)
		if _, ok := err.(*addressError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}

func (h *addressHttpHandler) DeleteAddress(c echo.Context) error {
	addressId := c.Param("id")
	if err := h.addressUsecase.CheckAddressId(&addressId); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*addressError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	if err := h.addressUsecase.DeleteAddress(&addressId); err != nil {
		log.Errorf("Error deleting address: %v", err)
		if _, ok := err.(*addressError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}

func (h *addressHttpHandler) GetProvince(c echo.Context) error {
	provinces, err := h.addressUsecase.GetProvince()
	if err != nil {
		log.Errorf("Error getting all provinces: %v", err)
		if _, ok := err.(*addressError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", provinces)
}

func (h *addressHttpHandler) GetDistrictByProvinceId(c echo.Context) error {
	provinceId := c.Param("id")

	districts, err := h.addressUsecase.GetDistrictByProvinceId(&provinceId)
	if err != nil {
		log.Errorf("Error getting all districts by province id: %v", err)
		if _, ok := err.(*addressError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", districts)
}

func (h *addressHttpHandler) GetSubDistrictByDistrictId(c echo.Context) error {
	districtId := c.Param("id")

	subDistricts, err := h.addressUsecase.GetSubDistrictByDistrictId(&districtId)
	if err != nil {
		log.Errorf("Error getting all sub districts by district id: %v", err)
		if _, ok := err.(*addressError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", subDistricts)
}
