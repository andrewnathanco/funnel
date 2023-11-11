package funnel

import (
	"funnel/controller/funnel/routes"
	"funnel/model"

	"github.com/labstack/echo/v4"
)

type FunnelRouter struct {
}

func NewFunnelRouter() FunnelRouter {
	return FunnelRouter{}
}

func (mc FunnelController) GetRoutes() map[string]func(c echo.Context) error {
	router := map[string]func(c echo.Context) error{}
	router["index"] = routes.GetIndex
	return router
}

func (r FunnelRouter) ConfigureRouter(c model.IController, e *echo.Echo) {
	e.GET("/", c.GetRoutes()["index"])
}
