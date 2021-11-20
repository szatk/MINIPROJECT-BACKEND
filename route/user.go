package route

import (
	"MINIPROJECT-BACKEND/controller"
	"github.com/labstack/echo"
)

func NewUser(app *echo.Echo) {
	app.POST("/users", controller.CreateUserController)
}
