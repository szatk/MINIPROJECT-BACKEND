package route

import (
	"MINIPROJECT-BACKEND/controller"

	"github.com/labstack/echo"
)

func NewUser(app *echo.Echo) {
	app.POST("/userlogin", controller.CreateUserController)
	// app.POST("/user", controller.Update)
}
