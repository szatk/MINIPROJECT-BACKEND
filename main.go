<<<<<<< HEAD
=======
package main

import (
	"MINIPROJECT-BACKEND/config"
	"MINIPROJECT-BACKEND/route"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	var db config.DBConfig

	errorEnv := godotenv.Load(".env")

	if errorEnv != nil {
		log.Fatalf("Error loading .env file")
	}

	db.Host = os.Getenv("DB_HOST")
	db.Port = os.Getenv("DB_PORT")
	db.User = os.Getenv("DB_USER")
	db.Password = os.Getenv("DB_PASSWORD")
	db.DBName = os.Getenv("DB_NAME")

	config.Connection(db)
	e := route.RouteVersion1()

	port := os.Getenv("PORT")
	e.Start(":" + port)
}

// import (
// 	"MINIPROJECT-BACKEND/config"
// 	"MINIPROJECT-BACKEND/route"

// 	"github.com/labstack/echo/v4"
// )

// func main() {
// 	config.InitDB()
// 	config.InitLog()
// 	config.InitMigration()

// 	app := echo.New()

// 	route.NewHousePerca(app)
// 	route.NewJenisPerca(app)
// 	route.NewPekerjaPerca(app)
// 	route.NewTransaksiPerca(app)
// 	route.NewUserPerca(app)
// 	route.NewLoginPerca(app)

// 	app.Start(":8080")
// }
>>>>>>> OjekPerca
