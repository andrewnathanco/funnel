package model

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

var (
	ErrCouldNotParseTempaltes = fmt.Errorf("could not parse templates")
)

type TemplateRenderer struct {
	Templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)

}
type TemplateController struct {
	Template *template.Template
	Name string
}


