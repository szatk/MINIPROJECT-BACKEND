package route

import (
	"MINIPROJECT-BACKEND/controller"
	// "MINIPROJECT-BACKEND/middleware"

	"github.com/labstack/echo/v4"
)

func NewLoginPerca(app *echo.Echo) {
	app.POST("createlogin", controller.CreateUserLogin)
	app.POST("login", controller.Login)
	// app.GET("userlogin", controller.GetAllUserLogin, middleware.AuthJWTMiddleware)
	// app.GET("userlogin/:id", controller.GetUserLoginByID, middleware.AuthJWTMiddleware)
	// app.PUT("userlogin/:id", controller.UpdateUserLogin, middleware.AuthJWTMiddleware)
}
