package funnel

import (
	fun_tmpl "funnel/controller/funnel/template"
	"funnel/model"
	"html/template"
)

type FunnelController struct{}

func NewFunnelController() *FunnelController {
	return &FunnelController{}
}

func (mc FunnelController) GetTemplates() []model.TemplateController {
	func_map := template.FuncMap{
		"getVersion": getVersion,
		"moviesLeft": getMoviesLeft,
		"getDate": getDate,
		"getCurrentTheme": getCurrentTheme,
		"getThemes": getThemes,
		"newOption": newOption,
		"getReleaseYear": getReleaseYear,
	}

	templates := []model.TemplateController{
		fun_tmpl.NewIndexTemplateController(func_map),
		fun_tmpl.NewErrorTemplateController(func_map),
		fun_tmpl.NewBoardTemplateController(func_map),
	}

	return templates
}
