package route

import (
	"MINIPROJECT-BACKEND/controller"

	"github.com/labstack/echo/v4"
)

func NewPekerjaPerca(app *echo.Echo) {
	app.POST("pekerjaperca", controller.PekerjaPercaRegister)
	app.GET("pekerjaperca", controller.GetAllPekerjaPerca)
	app.GET("pekerjaperca/:id", controller.GetPekerjaPecraById)
	app.PUT("pekerjaperca/:id", controller.UpdatePekerjaPerca)
	app.DELETE("pekerjaperca/:id", controller.DeletePekerjaPerca)
}
