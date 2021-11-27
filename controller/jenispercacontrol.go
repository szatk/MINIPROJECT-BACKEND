package controller

import (
	"MINIPROJECT-BACKEND/config"
	"MINIPROJECT-BACKEND/model/respon"
	"MINIPROJECT-BACKEND/model/transaksi"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Add Jenis Perca
func AddJenisPerca(c echo.Context) error {
	var jenisPerca transaksi.JenisPerca
	c.Bind(&jenisPerca)

	result := config.DB.Create(&jenisPerca)
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
		Data:    &jenisPerca,
	})
}

// Get All Data Jenis Perca
func GetAllJenisPerca(c echo.Context) error {
	var jenisPerca = []transaksi.JenisPerca{}

	result := config.DB.Find(&jenisPerca)
	if result.Error != nil {
		return c.JSON(http.StatusNotAcceptable, respon.BaseRespon{
			Code:    http.StatusNotAcceptable,
			Message: "Cannot retrieve data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, respon.BaseRespon{
		Code:    http.StatusOK,
		Message: "Successful retrieve data",
		Data:    &jenisPerca,
	})
}

// Get Jenis Perca by ID
func GetJenisPercaById(c echo.Context) error {
	var jenisPerca transaksi.JenisPerca

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.First(&jenisPerca, id)

	if result.Error != nil {
		return c.JSON(http.StatusGone, respon.BaseRespon{
			Code:    http.StatusGone,
			Message: "Cannot retrieve data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, respon.BaseRespon{
		Code:    http.StatusOK,
		Message: "Successful retrieve data",
		Data:    &jenisPerca,
	})
}

// Update Jenis Perca
func UpdateJenisSampah(c echo.Context) error {
	var JenisPerca transaksi.JenisPerca
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.First(&JenisPerca, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotAcceptable, respon.BaseRespon{
			Code:    http.StatusNotAcceptable,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	c.Bind(&JenisPerca)
	result = config.DB.Save(&JenisPerca)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Cannot Update data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusAccepted, respon.BaseRespon{
		Code:    http.StatusAccepted,
		Message: "Successful update data",
		Data:    &JenisPerca,
	})
}

// Delete Jenis Perca
func DeleteJenisPerca(c echo.Context) error {
	var jenisPerca transaksi.JenisPerca

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.First(&jenisPerca, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotAcceptable, respon.BaseRespon{
			Code:    http.StatusNotAcceptable,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	result = config.DB.Where("id = ?", id).Unscoped().Delete(&jenisPerca)
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
		Data:    &jenisPerca,
	})
}
