package template

import (
	"funnel/model"
	"html/template"
)

var (
	IndexTemplate = []string{
		"view/index.html",
		"view/index.tmpl",
		"view/decade-selector.tmpl",
		"view/option/red-option.tmpl",
		"view/option/yellow-option.tmpl",
		"view/option/green-option.tmpl",
	}
)

func NewIndexTemplateController(
	func_map template.FuncMap,
) model.TemplateController {
	index_template_files := []string{}

	// add buttons
	index_template_files = append(index_template_files, IndexTemplate...)
	index_template_files = append(index_template_files, BoardTemplates...)

	index_template := template.Must(
		template.New("index").Funcs(func_map).ParseFiles(
			index_template_files...,
		))

	return model.TemplateController{
		Template: index_template,
		Name:     "index.html",
	}

}
