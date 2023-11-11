package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthPing(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
