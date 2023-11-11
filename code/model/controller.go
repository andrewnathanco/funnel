package model

import "github.com/labstack/echo/v4"

type IController interface {
	GetTemplates() []TemplateController
	GetRoutes() map[string]func(c echo.Context) error
}

type IRouter interface {
	ConfigureRouter(IController, *echo.Echo)
}
