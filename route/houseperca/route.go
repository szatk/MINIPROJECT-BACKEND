package housepercaroute

import (
	"MINIPROJECT-BACKEND/controller"

	"github.com/labstack/echo/v4"
)

func HousePercaRouter(route *echo.Group) {

	route.POST("houseperca", controller.HousePercaRegister)
	route.GET("houseperca", controller.GetAllHousePerca)
	route.GET("houseperca/:id", controller.GetHousePercaById)
	route.PUT("houseperca/:id", controller.UpdateHousePerca)
	route.DELETE("houseperca/:id", controller.DeleteHousePerca)
}
