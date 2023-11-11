package routes

import (
	"funnel/db"
	"funnel/db/sql"
	"funnel/middleware"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ChangeTheme(c echo.Context) error {
	theme := c.QueryParam("theme")
	user_key := middleware.GetUserKeyFromContext(c)
	board, err := db.MuralDAL.GetBoardByKey(user_key)
	if err != nil {
		slog.Error(err.Error())
		return c.Render(http.StatusInternalServerError, "error.html", nil)
	}

	movie, err := db.MuralDAL.GetRandomMovieByDecade(theme)
	if err != nil {
		slog.Error(err.Error())
		return c.Render(http.StatusInternalServerError, "error.html", nil)
	}

	board.Theme = theme
	board.Movie = *movie
	board.MoviesLeft = db.MuralDAL.GetNumberPerDecade(sql.GetSQLDecade(theme))
	db.MuralDAL.SetBoardByKey(user_key, *board)
	return c.Render(http.StatusOK, "index.html", board)
}
