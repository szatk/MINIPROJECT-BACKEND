package main

import (
	"MINIPROJECT-BACKEND/config"
	"MINIPROJECT-BACKEND/route"
)

func main() {
	config.Connection()
	e := route.RouteVersion1()
	e.Start(":8080")
}
