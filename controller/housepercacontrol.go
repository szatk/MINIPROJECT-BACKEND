package controller

import (
	"MINIPROJECT-BACKEND/config"
	"MINIPROJECT-BACKEND/model/houseperca"
	"MINIPROJECT-BACKEND/model/respon"
	"MINIPROJECT-BACKEND/ojekperca/housepercaojek"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type APIEnv struct {
	DB *gorm.DB
}

// Create House Perca
func (a *APIEnv) HousePercaRegister(c echo.Context) error {
	var housePerca houseperca.HousePerca
	c.Bind(&housePerca)

	result := config.DB.Create(&housePerca)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Error save data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, respon.BaseRespon{
		Code:    http.StatusCreated,
		Message: "Successful create data",
		Data:    &housePerca,
	})
}

// Get All House Perca Data
func (a *APIEnv) GetAllHousePerca(c echo.Context) error {
	var housePerca = []houseperca.HousePerca{}

	result := config.DB.Find(&housePerca)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Cannot retrieve data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, respon.BaseRespon{
		Code:    http.StatusOK,
		Message: "Successful retrieve data",
		Data:    &housePerca,
	})
}

// Get House Perca by ID
func (a *APIEnv) GetHousePercaById(c echo.Context) error {
	var housePerca houseperca.HousePerca

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.First(&housePerca, id)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Cannot retrieve data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, respon.BaseRespon{
		Code:    http.StatusOK,
		Message: "Successful retrieve data",
		Data:    &housePerca,
	})
}

// Update House Perca
func (a *APIEnv) UpdateHousePerca(c echo.Context) error {
	var housePerca houseperca.HousePerca

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	housePerca, exists, err := housepercaojek.GetHousePercaByID(fmt.Sprint(id), a.DB)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "cannot retrieve data from database",
			Data:    nil,
		})
	}

	if !exists {
		return c.JSON(http.StatusNotFound, respon.BaseRespon{
			Code:    http.StatusNotFound,
			Message: "data not found",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusAccepted, respon.BaseRespon{
		Code:    http.StatusAccepted,
		Message: "Successful update data",
		Data:    &housePerca,
	})
}

// Delete House Perca
func (a *APIEnv) DeleteHousePerca(c echo.Context) error {
	var housePerca houseperca.HousePerca

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.First(&housePerca, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotAcceptable, respon.BaseRespon{
			Code:    http.StatusNotAcceptable,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	result = config.DB.Where("id = ?", id).Unscoped().Delete(&housePerca)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Cannot Delete data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusAccepted, respon.BaseRespon{
		Code:    http.StatusAccepted,
		Message: "Successful Delete data",
		Data:    &housePerca,
	})
}
