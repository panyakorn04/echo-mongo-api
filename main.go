package main

import (
	"github.com/labstack/echo/v4"
	"github.com/panyakorn4/echo-mongo-api/config"
	"github.com/panyakorn4/echo-mongo-api/routes"
)

func main() {
	e := echo.New()

	// run database connection
	config.ConnectDB()

	// routers
	routes.UserRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}
