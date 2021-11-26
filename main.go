package main

import (
	"MINIPROJECT-BACKEND/route"

	"github.com/labstack/echo/v4"
)

func main() {
	// config.InitDB()
	// // config.InitLog()
	// config.InitMigration()

	app := echo.New()
	route.InitRoute(app)
	app.Start(":8080")
}
