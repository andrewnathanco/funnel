package template

import (
	"funnel/model"
	"html/template"
)

var (
	ErrorTemplates = []string{
		"view/error.html",
	}
)

func NewErrorTemplateController(
	func_map template.FuncMap,
) model.TemplateController {
	error_template_files := []string{}

	// add buttons
	error_template_files = append(error_template_files, ErrorTemplates...)

	error_template := template.Must(
		template.New("error_template").Funcs(func_map).ParseFiles(
			error_template_files...,
		))

	return model.TemplateController{
		Template: error_template,
		Name:     "error.html",
	}

}
