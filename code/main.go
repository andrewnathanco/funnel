package main

import (
	"errors"
	"fmt"
	"funnel/config"
	"funnel/controller"
	"funnel/db"
	"funnel/db/sql"
	mural_middleware "funnel/middleware"
	"funnel/worker"
	"html/template"
	"io"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	ErrCouldNotParseTempaltes = fmt.Errorf("could not parse templates")
)

type TemplateRenderer struct {
	templates map[string]*template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]

	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}

	return tmpl.ExecuteTemplate(w, name, data)
}

func Must(err error) {
	if err != nil {
		slog.Error(err.Error())
		panic(1)
	}
}

func main() {
	// load env
	Must(godotenv.Load())

	// validate env
	Must(config.ValidateENV())
	funnelDAL, err := sql.NewSQLiteDal(os.Getenv(config.EnvFunnelDB))
	Must(err)

	db.FunnelDAL = funnelDAL

	// start setting up
	e := echo.New()

	// define templates
	templates := map[string]*template.Template{}

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	mural_middleware.InitSession()
	e.Use(mural_middleware.GetUserKey)

	// setup schedular
	scheduler := worker.NewFunnelSchedular()

	// setup the project
	scheduler.InitProgram()
	Must(scheduler.RegisterWorkers())

	// start scheduler
	scheduler.StartScheduler()

	// Define your routes and handlers here
	// setup routes and controllers
	route_conrollers := controller.GetRouteControllers()

	for _, route_controller := range route_conrollers {
		// add templates
		for _, template := range route_controller.Controller.GetTemplates() {
			templates[template.Name] = template.Template
		}

		// add routes
		route_controller.Router.ConfigureRouter(route_controller.Controller, e)
	}

	e.Renderer = &TemplateRenderer{
		templates: templates,
	}

	// setup routes
	e.Static("/static", "./static")
	e.Logger.Fatal(e.Start(":2222"))
}
