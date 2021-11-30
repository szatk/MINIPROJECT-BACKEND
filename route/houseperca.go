package route

import (
	"MINIPROJECT-BACKEND/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func HousePercaRouter(route *echo.Group) {
	jwtSecretKey := os.Getenv("SECRET_JWT")
	jwt := middleware.JWT([]byte(jwtSecretKey))

	api := controller.APIEnv{DB: Database.DB}

	route.POST("houseperca", api.HousePercaRegister, jwt)
	route.GET("houseperca", api.GetAllHousePerca, jwt)
	route.GET("houseperca/:id", api.GetHousePercaById, jwt)
	route.PUT("houseperca/:id", api.UpdateHousePerca, jwt)
	route.DELETE("houseperca/:id", api.DeleteHousePerca, jwt)
}

// import (
// 	"MINIPROJECT-BACKEND/controller"

// 	"github.com/labstack/echo/v4"
// )

// func NewHousePerca(app *echo.Echo) {
// 	app.POST("houseperca", controller.HousePercaRegister)
// 	app.GET("houseperca", controller.GetAllHousePerca)
// 	app.GET("houseperca/:id", controller.GetHousePercaById)
// 	app.PUT("houseperca/:id", controller.UpdateHousePerca)
// 	app.DELETE("houseperca/:id", controller.DeleteHousePerca)
// }
