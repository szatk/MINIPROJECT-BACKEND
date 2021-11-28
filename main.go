package main

import (
	"MINIPROJECT-BACKEND/config"
	"MINIPROJECT-BACKEND/route"

	"github.com/labstack/echo/v4"
)

// func main() {
// 	config.Connection()
// 	e := route.RouteVersion1()
// 	e.Start(":8080")
// }

func main() {
	config.InitDB()
	// config.InitLog()
	config.InitMigration()

	app := echo.New()

	route.NewHousePerca(app)
	route.NewJenisPerca(app)
	route.NewPekerjaPerca(app)
	route.NewTransaksiPerca(app)

	app.Start(":8080")
}
