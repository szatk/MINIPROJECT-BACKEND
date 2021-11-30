package route

import (
	"MINIPROJECT-BACKEND/controller"
	"MINIPROJECT-BACKEND/middleware"

	"github.com/labstack/echo/v4"
)

func NewLoginPerca (app *echo.Echo) {
app.POST("/users/login", controller.LoginUserController)
}
