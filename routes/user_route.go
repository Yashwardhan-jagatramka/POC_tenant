package routes

import (
	"github.com/labstack/echo/v4"

	"project-tenant/controller"
)

func UserRoute(e *echo.Echo) {
	// all routes related to user yaha add karna hai...
	//e.Use(middleware.BasicAuth(func(key string, name string, context echo.Context) (bool, error) {
	//	if key == controller.KEY && name == "Yashwardhan" {
	//		return true, nil
	//	}
	//	return false, nil
	//}))
	e.POST("/create", controller.CreateTenant)
	e.GET("/get-tenant", controller.GetATenant)
	e.PUT("/update", controller.UpdateATenant)
	e.GET("/get-all-tenants", controller.GetAllTenants)
}
