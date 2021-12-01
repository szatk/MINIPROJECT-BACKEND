package route

import (
	"MINIPROJECT-BACKEND/Middleware"
	HousePercaRoute "MINIPROJECT-BACKEND/route/HousePercaRoute"

	// JenisPercaRoute "MINIPROJECT-BACKEND/route/Jenisperca"
	// LoginRoute "MINIPROJECT-BACKEND/route/Jenisperca"
	// PekerjaPercaRoute "MINIPROJECT-BACKEND/route/PekerjaPerca"
	// TransaksiRoute "MINIPROJECT-BACKEND/route/Transaksi"
	// UsersRoute "golang-final-project/Routes/Users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteVersion1() *echo.Echo {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} => ${host} ${method} uri=${uri}  ||  ${status} => ${latency_human} ${remote_ip} \n",
	}))
	e.Use(middleware.BodyDump(Middleware.Log))

	r1 := e.Group("api/v1/")
	// UsersRoute.UsersRouter(r1)
	// LoginRoute.UserLoginRoute(r1)
	HousePercaRoute.HousePercaRouter(r1)
	// PekerjaPercaRoute.PekerjaPercaRouter(r1)
	// JenisPercaRoute.JenisPercaRouter(r1)
	// TransaksiRoute.TransaksiRouter(r1)

	return e
}
