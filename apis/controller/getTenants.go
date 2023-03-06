package controller

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"

	"project-tenant/apis/service"
	"project-tenant/pkg/models"
)

func GetAllTenants(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []models.Tenant
	defer cancel()
	ans, err := service.GetAllTenants(c, users, ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.TenantResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	return c.JSON(http.StatusOK, models.TenantResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": ans}})
}
