package controller

import (
	"MINIPROJECT-BACKEND/config"
	"MINIPROJECT-BACKEND/model/respon"
	"MINIPROJECT-BACKEND/model/transaksi"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Add Transaction
func AddTransaction(c echo.Context) error {
	var transaction transaksi.TransaksiReq
	c.Bind(&transaction)

	var totalQty int
	var totalPrice float32
	for _, item := range transaction.Detail {
		totalQty += item.Qty
		totalPrice += item.TotalPrice
	}

	transaction.TotalQty = totalQty
	transaction.TotalTransaksi = totalPrice

	res := config.DB.Create(&transaksi.Transaksi{
		HousePercaId:    transaction.HousePercaId,
		UserId:          transaction.UserId,
		PekerjaPercaId:  transaction.PekerjaPercaId,
		TotalQty:        transaction.TotalQty,
		TotalTransaksi:  transaction.TotalTransaksi,
		Status:          transaction.Status,
		DetailTransaksi: transaction.Detail,
	})

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Cannot save data to database",
			Data:    nil,
		})
	}

	var trans = transaksi.Transaksi{}

	config.DB.Preload("DetailTransaksi").Last(&trans)

	return c.JSON(http.StatusCreated, respon.BaseRespon{
		Code:    http.StatusCreated,
		Message: "Successful create data",
		Data:    &transaction,
	})
}

// Get All Data Transaction
func GetAllTransaction(c echo.Context) error {
	var trans = []transaksi.Transaksi{}

	result := config.DB.Preload("DetailTransaction").Find(&trans)
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
		Data:    &trans,
	})
}

// Get Jenis Sampah by ID
func GetTransactionById(c echo.Context) error {
	var trans transaksi.Transaksi

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.Preload("DetailTransaction").First(&trans, id)

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
		Data:    &trans,
	})
}

// Update Bank Sampah
func UpdateTansaction(c echo.Context) error {
	var transaction transaksi.Transaksi

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.Preload("DetailTransaction").First(&transaction, id)
	if result.Error != nil {
		return c.JSON(http.StatusGone, respon.BaseRespon{
			Code:    http.StatusGone,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	c.Bind(&transaction)
	result = config.DB.Save(&transaction)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Cannot Update data to database",
			Data:    nil,
		})
	}

	config.DB.Preload("DetailTransaction").First(&transaction, id)
	return c.JSON(http.StatusAccepted, respon.BaseRespon{
		Code:    http.StatusAccepted,
		Message: "Successful update data",
		Data:    &transaction,
	})
}
