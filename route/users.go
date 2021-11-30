package route

import (
	"MINIPROJECT-BACKEND/config"
	"MINIPROJECT-BACKEND/controller"
	"MINIPROJECT-BACKEND/middleware"

	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)



func NewUserPerca (app *echo.Echo) {
	
	app.POST("/users", controller.UserRegister)
	app.GET("/users", controller.GetAllUser)
	app.GET("/users/:id", controller.GetUserById, middleware.AuthJWTMiddleware)
	app.DELETE("/users/:id", controller.DeleteUser middleware.AuthJWTMiddleware)
	app.PUT("/users/:id", controller.UpdateUser, middleware.AuthJWTMiddleware)
}