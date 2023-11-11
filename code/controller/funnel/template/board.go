package template

import (
	"funnel/model"
	"html/template"
)

var (
	BoardTemplates = []string{
		"view/board/board.tmpl",
		"view/board/board.html",
	}
)

func NewBoardTemplateController(
	func_map template.FuncMap,
) model.TemplateController {
	board_template_files := []string{}

	// add buttons
	board_template_files = append(board_template_files, BoardTemplates...)
	board_template_files = append(board_template_files, IndexTemplate...)

	board_template := template.Must(
		template.New("board_template").Funcs(func_map).ParseFiles(
			board_template_files...,
		))

	return model.TemplateController{
		Template: board_template,
		Name:     "board.html",
	}

}
