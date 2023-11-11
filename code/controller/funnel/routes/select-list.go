package routes

import (
	"funnel/db"
	"funnel/middleware"
	"funnel/model"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func must(list string) string {
	if list != model.BLACK_LIST && list != model.GREEN_LIST && list != model.YELLOW_LIST {
		return model.GREEN_LIST
	}

	return list
}

func SelectList(c echo.Context) error {
	list := c.QueryParam("list")
	user_key := middleware.GetUserKeyFromContext(c)
	board, err := db.MuralDAL.GetBoardByKey(user_key)
	list = must(list)
	if err != nil {
		slog.Error(err.Error())
		return c.Render(http.StatusInternalServerError, "error.html", nil)
	}

	board.SelectedList = list
	db.MuralDAL.SetBoardByKey(user_key, *board)
	return c.Render(http.StatusOK, "board.html", board)
}
