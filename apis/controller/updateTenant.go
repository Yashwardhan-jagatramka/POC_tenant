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

func UpdateATenant(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId, _ := strconv.Atoi(c.QueryParam("tenantId"))
	defer cancel()
	//validate the request body
	var reqTenant models.Tenant
	if err := c.Bind(&reqTenant); err != nil {
		return c.JSON(http.StatusBadRequest, models.TenantResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	//use the validator library to validate required fields
	if validationErr := validate.Struct(reqTenant); validationErr != nil {
		return c.JSON(http.StatusBadRequest, models.TenantResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}
	result, err := service.UpdateTenant(ctx, userId, reqTenant)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.TenantResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	return c.JSON(http.StatusOK, models.TenantResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": result}})
}
