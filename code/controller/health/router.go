package health

import (
	"funnel/controller/health/routes"
	"funnel/model"

	"github.com/labstack/echo/v4"
)

type HealthRouter struct {
}

func NewHealthRouter() HealthRouter {
	return HealthRouter{}
}

func (hc HealthController) GetRoutes() map[string]func(c echo.Context) error {
	router := map[string]func(c echo.Context) error{}
	router["health"] = routes.Health
	router["healthping"] = routes.HealthPing
	return router
}

func (r HealthRouter) ConfigureRouter(c model.IController, e *echo.Echo) {
	e.GET("/health", c.GetRoutes()["health"])
	e.GET("/healthping", c.GetRoutes()["healthping"])
}
