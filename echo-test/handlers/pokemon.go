package handlers

import (
	"echo-test/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllKanto(c echo.Context) error {
	result, err := services.GetAllKanto()

	if err != nil {
		return c.String(http.StatusFailedDependency, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
