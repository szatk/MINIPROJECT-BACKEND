package route

import (
	"MINIPROJECT-BACKEND/factory"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	presenter := factory.Init()
	e.POST("/userPekerjaPerca", presenter.DriverPresenter.CreatePekerjaPerca)
}
