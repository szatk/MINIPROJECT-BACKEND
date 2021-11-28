package route

import (
	"MINIPROJECT-BACKEND/controller"

	"github.com/labstack/echo/v4"
)

func NewTransaksiPerca(app *echo.Echo) {
	app.POST("transaksiperca", controller.AddTransaksi)
	app.GET("transaksiperca", controller.GetAllTransaksi)
	app.GET("transaksiperca/:id", controller.GetTransaksiById)
	app.PUT("transaksiperca/:id", controller.UpdateTransaksi)
	app.DELETE("transaksiperca/:id", controller.DeleteJenisPerca)
}
