package route

import (
	"MINIPROJECT-BACKEND/config"
	"MINIPROJECT-BACKEND/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JenisPercaRouter(route *echo.Group) {
	jwtSecretKey := os.Getenv("SECRET_JWT")
	jwt := middleware.JWT([]byte(jwtSecretKey))

	api := controller.APIEnv{DB: config.DB}

	route.POST("jenisperca", api.AddJenisPerca, jwt)
	route.GET("jenisperca", api.GetAllJenisPerca, jwt)
	route.GET("jenisperca/:id", api.GetJenisPercaById, jwt)
	route.PUT("jenisperca/:id", api.UpdateJenisPerca, jwt)
	route.DELETE("jenisperca/:id", api.DeleteJenisPerca, jwt)
}

// import (
// 	"MINIPROJECT-BACKEND/controller"

// 	"github.com/labstack/echo/v4"
// )

// func NewJenisPerca(app *echo.Echo) {
// 	app.POST("jenisperca", controller.AddJenisPerca)
// 	app.GET("jenisperca", controller.GetAllJenisPerca)
// 	app.GET("jenisperca/:id", controller.GetJenisPercaById)
// 	app.PUT("jenisperca/:id", controller.UpdateJenisPerca)
// 	app.DELETE("jenisperca/:id", controller.DeleteJenisPerca)
// }
