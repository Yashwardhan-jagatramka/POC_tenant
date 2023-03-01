package main

import (
	"github.com/labstack/echo/v4"

	"project-tenant/configs"
	"project-tenant/routes"
)

func main() {
	e := echo.New()

	configs.ConnectDB()

	routes.UserRoute(e)

	e.Logger.Fatal(e.Start(":6000"))
}
