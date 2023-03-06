package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"

	"project-tenant/apis/service"
	"project-tenant/pkg/models"
)

func GetATenant(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	userId, err1 := strconv.Atoi(c.QueryParam("tenantId"))
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, models.TenantResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err1.Error()}})
	}
	user2, err3 := service.GetATenant(c, ctx, userId)
	if err3 != nil {
		return c.JSON(http.StatusInternalServerError, models.TenantResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err3.Error()}})
	}
	return c.JSON(http.StatusOK, models.TenantResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": user2}})
}
