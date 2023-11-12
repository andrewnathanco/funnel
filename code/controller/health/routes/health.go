package routes

import (
	"funnel/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthStatus string

const (
	DATABSE_FAILED = "FAILED"
	SUCCESS        = "SUCCESS"
)

type HealthObject struct {
	Status HealthStatus
}

func Health(c echo.Context) error {
	var health HealthObject
	err := db.FunnelDAL.PingDatabse()
	if err != nil {
		health.Status = DATABSE_FAILED
	}

	health.Status = SUCCESS
	return c.JSON(http.StatusOK, health)
}
