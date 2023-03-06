package routes

import (
	"github.com/labstack/echo/v4"

	"project-tenant/apis/controller"
)

func TenantRoute(e *echo.Echo) {
	// all routes related to user yaha add karna hai...
	e.POST("/create", controller.CreateTenant)
	e.GET("/get-tenant", controller.GetATenant)
	e.PUT("/update", controller.UpdateATenant)
	e.GET("/get-all-tenants", controller.GetAllTenants)
}
