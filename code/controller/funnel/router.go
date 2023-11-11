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
	router["select-list"] = routes.SelectList
	router["change-theme"] = routes.ChangeTheme
	router["submit"] = routes.Submit
	return router
}

func (r FunnelRouter) ConfigureRouter(c model.IController, e *echo.Echo) {
	e.GET("/", c.GetRoutes()["index"])
	e.PUT("/funnel/change-theme", c.GetRoutes()["change-theme"])
	e.PUT("/funnel/select-list", c.GetRoutes()["select-list"])
	e.PUT("/funnel/submit", c.GetRoutes()["submit"])
}
