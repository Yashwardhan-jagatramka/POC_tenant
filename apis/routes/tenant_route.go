package routes

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"project-tenant/apis/controller"
)

var Logfile, _ = os.Create("logfile.txt")

func TenantRoute(e *echo.Echo) {
	// all routes related to tanent yaha add karna hai...

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: Logfile,
	}))
	e.POST("/create", controller.CreateTenant)
	e.GET("/get-tenant", controller.GetATenant)
	e.PUT("/update", controller.UpdateATenant)
	e.GET("/get-all-tenants", controller.GetAllTenants)
}
