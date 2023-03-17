package main

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"project-tenant/apis/controller"
	"project-tenant/apis/routes"
	"project-tenant/pkg/configs"
)

func main() {
	e := echo.New()

	configs.ConnectDB()

	routes.TenantRoute(e)

	e.Use(middleware.Recover())
	e.Pre(middleware.BasicAuth(func(name string, key string, context echo.Context) (bool, error) {
		if name == "Yash" && key == controller.KEY {
			return true, nil
		}
		return false, errors.New("Authentication Failed")
	}))

	e.Logger.Fatal(e.Start(":6000"))
	defer routes.Logfile.Close()
}
