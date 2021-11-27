package route

import (
	"MINIPROJECT-BACKEND/controller"

	"github.com/labstack/echo/v4"
)

func NewHousePerca(app *echo.Echo) {
	app.POST("houseperca", controller.HousePercaRegister)
	app.GET("houseperca", controller.GetAllHousePerca)
	app.GET("houseperca/:id", controller.GetHousePercaById)
	app.PUT("houseperca/:id", controller.UpdateHousePerca)
	app.DELETE("houseperca/:id", controller.DeleteHousePerca)
}
