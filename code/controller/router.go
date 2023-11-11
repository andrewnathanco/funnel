package controller

import (
	"funnel/controller/funnel"
	"funnel/controller/health"
	"funnel/model"
)

type RouteController struct {
	Router     model.IRouter
	Controller model.IController
}

func GetRouteControllers() []RouteController {
	route_controllers := []RouteController{
		{
			Router:     funnel.NewFunnelRouter(),
			Controller: funnel.NewFunnelController(),
		},
		{
			Router:     health.NewHealthRouter(),
			Controller: health.NewHealthController(),
		},
	}

	return route_controllers
}
