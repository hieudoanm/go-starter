package health_controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
}
