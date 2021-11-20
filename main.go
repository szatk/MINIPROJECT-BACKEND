package main

import (
	"MINIPROJECT-BACKEND/config"
	"MINIPROJECT-BACKEND/route"
	"github.com/labstack/echo"
)

func main() {
	config.InitDB()
	// config.InitLog()
	config.InitMigration()

	app := echo.New()
	
	route.NewUser(app)
	
	app.Start(":8080")
}
