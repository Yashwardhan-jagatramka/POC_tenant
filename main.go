package main

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"project-tenant/configs"
	"project-tenant/controller"
	"project-tenant/routes"
)

func main() {
	e := echo.New()

	configs.ConnectDB()

	routes.UserRoute(e)

	e.Use(middleware.Recover())
	e.Pre(middleware.BasicAuth(func(name string, key string, context echo.Context) (bool, error) {
		if name == "Yash" && key == controller.KEY {
			return true, nil
		}
		return false, errors.New("Authentication Failed")
	}))

	e.Logger.Fatal(e.Start(":6000"))
}
