package route

import (
	housepercaroute "MINIPROJECT-BACKEND/route/houseperca"

	"github.com/labstack/echo/v4"
)

func RouteVersion1() *echo.Echo {
	e := echo.New()
	r1 := e.Group("api/v1/")
	housepercaroute.HousePercaRouter(r1)

	return e
}
