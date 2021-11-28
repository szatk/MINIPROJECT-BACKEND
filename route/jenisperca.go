package route

import (
	"MINIPROJECT-BACKEND/controller"

	"github.com/labstack/echo/v4"
)

func NewJenisPerca(app *echo.Echo) {
	app.POST("jenisperca", controller.AddJenisPerca)
	app.GET("jenisperca", controller.GetAllJenisPerca)
	app.GET("jenisperca/:id", controller.GetJenisPercaById)
	app.PUT("jenisperca/:id", controller.UpdateJenisPerca)
	app.DELETE("jenisperca/:id", controller.DeleteJenisPerca)
}
