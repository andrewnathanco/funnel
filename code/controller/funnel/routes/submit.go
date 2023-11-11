package routes

import (
	"funnel/db"
	"funnel/db/sql"
	"funnel/middleware"
	"funnel/model"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Submit(c echo.Context) error {
	user_key := middleware.GetUserKeyFromContext(c)
	board, err := db.MuralDAL.GetBoardByKey(user_key)
	if err != nil {
		slog.Error(err.Error())
		return c.Render(http.StatusInternalServerError, "error.html", nil)
	}

	movie_key := strconv.Itoa(board.Movie.ID)
	db.MuralDAL.InsertMoveIntoList(board.SelectedList, movie_key)
	movie, err := db.MuralDAL.GetRandomMovieByDecade(board.Theme)
	if err != nil {
		slog.Error(err.Error())
		return c.Render(http.StatusInternalServerError, "error.html", nil)
	}
	board.SelectedList = model.GREEN_LIST
	board.Movie = *movie
	board.MoviesLeft = db.MuralDAL.GetNumberPerDecade(sql.GetSQLDecade(board.Theme))
	db.MuralDAL.SetBoardByKey(user_key, *board)
	return c.Render(http.StatusOK, "index.html", board)
}
