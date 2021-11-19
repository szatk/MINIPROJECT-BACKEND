package ProdusenSepatuHandler

import (
	"fmt"
	"golang-final-project/Driver/SepatuDriver"
	"golang-final-project/Models/ProdusenSepatu"
	"golang-final-project/Models/Response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type APIEnv struct {
	DB *gorm.DB
}

// Register Produsen Sepatu
func (a *APIEnv) ProdusenSepatuRegister(c echo.Context) error {
	var produsenSepatu ProdusenSepatu.ProdusenSepatu
	c.Bind(&produsenSepatu)

	err := SepatuDriver.CreateProdusenSepatu(a.DB, &produsenSepatu)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "error save data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, Response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "successful create data",
		Data:    &produsenSepatu,
	})
}

// Get All Produsen Sepatu Data
func (a *APIEnv) GetAllBankSampah(c echo.Context) error {
	produsenSepatu, err := SepatuDriver.GetAllProdusenSepatu(a.DB)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "cannot retrieve data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, Response.BaseResponse{
		Code:    http.StatusOK,
		Message: "successful retrieve data",
		Data:    &produsenSepatu,
	})
}

// Get Produsen Sepatu by ID
func (a *APIEnv) GetProdusenSepatuById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "path parameter invalid",
			Data:    nil,
		})
	}

	produsenSepatu, exists, err := SepatuDriver.GetProdusenSepatuByID(fmt.Sprint(id), a.DB)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "cannot retrieve data from database",
			Data:    nil,
		})
	}

	if !exists {
		return c.JSON(http.StatusNotFound, Response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "data not found",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, Response.BaseResponse{
		Code:    http.StatusOK,
		Message: "successful retrieve data",
		Data:    &produsenSepatu,
	})
}

// Update Produsen Sepatu
func (a *APIEnv) UpdateProdusenSepatu(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "path parameter invalid",
			Data:    nil,
		})
	}

	produsenSepatu, exists, err := SepatuDriver.GetProdusenSepatuByID(fmt.Sprint(id), a.DB)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotAcceptable, Response.BaseResponse{
			Code:    http.StatusNotAcceptable,
			Message: "data not found",
			Data:    nil,
		})
	}

	if !exists {
		return c.JSON(http.StatusNotFound, Response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "record not exist",
			Data:    nil,
		})
	}

	c.Bind(&produsenSepatu)
	if err := SepatuDriver.UpdateProdusenSepatu(a.DB, &produsenSepatu, fmt.Sprint(id)); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "cannot update data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusAccepted, Response.BaseResponse{
		Code:    http.StatusAccepted,
		Message: "successful update data",
		Data:    &produsenSepatu,
	})
}

func (a *APIEnv) DeleteProdusenSepatu(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "path parameter invalid",
			Data:    nil,
		})
	}

	_, exists, err := SepatuDriver.GetProdusenSepatuByID(fmt.Sprint(id), a.DB)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "error process to database",
			Data:    nil,
		})
	}

	if !exists {
		return c.JSON(http.StatusNotFound, Response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "record not exist",
			Data:    nil,
		})
	}

	err = SepatuDriver.DeleteProdusenSepatu(fmt.Sprint(id), a.DB)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "cannot delete data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusAccepted, Response.BaseResponse{
		Code:    http.StatusAccepted,
		Message: "successful delete data",
		Data:    nil,
	})
}
