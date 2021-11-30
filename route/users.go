package route

import (
	"MINIPROJECT-BACKEND/config"
	"MINIPROJECT-BACKEND/controller"
	"MINIPROJECT-BACKEND/Middleware"

	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)



func NewUserPerca (app *echo.Echo) {
	
	app.POST("/users", controller.UserRegister)
	app.GET("/users", controller.GetAllUser)
	app.GET("/users/:id", controller.GetUserById, Middleware.AuthJWTMiddleware)
	app.DELETE("/users/:id", controller.DeleteUser Middleware.AuthJWTMiddleware)
	app.PUT("/users/:id", controller.UpdateUser, Middleware.AuthJWTMiddleware)
}