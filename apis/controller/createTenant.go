package controller

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"

	"project-tenant/apis/service"
	"project-tenant/pkg/models"
)

func CreateTenant(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.Tenant
	defer cancel()
	//validate the request body
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, models.TenantResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.JSON(http.StatusBadRequest, models.TenantResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}
	result, err := service.CreateTenant(ctx, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.TenantResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	return c.JSON(http.StatusCreated, models.TenantResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}
