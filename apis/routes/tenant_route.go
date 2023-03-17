package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"project-tenant/apis/controller"
	"project-tenant/apis/logger"
)

func TenantRoute(e *echo.Echo) {
	// all routes related to tanent yaha add karna hai...

	Logfile := logger.CreateLogFile()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: Logfile,
	}))
	e.POST("/create", controller.CreateTenant)
	e.GET("/get-tenant", controller.GetATenant)
	e.PUT("/update", controller.UpdateATenant)
	e.GET("/get-all-tenants", controller.GetAllTenants)
}
