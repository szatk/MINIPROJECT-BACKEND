package route

import (
	"MINIPROJECT-BACKEND/controller"
	// "MINIPROJECT-BACKEND/middleware"

	"github.com/labstack/echo/v4"
)

func NewUserPerca(app *echo.Echo) {
	app.POST("/users", controller.UserRegister)
	app.GET("/users", controller.GetAllUser)
	// app.GET("/users/:id", controller.GetUserByID, middleware.AuthJWTMiddleware)
	// app.DELETE("/users/:id", controller.DeleteUser, middleware.AuthJWTMiddleware)
	// app.PUT("/users/:id", controller.UpdateUser, middleware.AuthJWTMiddleware)
}
