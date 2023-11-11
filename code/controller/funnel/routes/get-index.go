package routes

import (
	"funnel/db"
	"funnel/db/sql"
	"funnel/middleware"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetIndex(c echo.Context) error {
	user_key := middleware.GetUserKeyFromContext(c)
	board, err := db.MuralDAL.GetBoardByKey(user_key)
	if err != nil {
		slog.Error(err.Error())
		return c.Render(http.StatusInternalServerError, "error.html", nil)
	}

	board.MoviesLeft = db.MuralDAL.GetNumberPerDecade(sql.GetSQLDecade(board.Theme))
	return c.Render(http.StatusOK, "index.html", &board)
}
