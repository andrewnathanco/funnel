package routes

import (
	"funnel/db"
	"funnel/middleware"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetIndex(c echo.Context) error {
	user_key := middleware.GetUserKeyFromContext(c)
	session, err := db.FunnelDAL.GetSessionForUser(user_key)
	if err != nil {
		slog.Error(err.Error())
		return c.Render(http.StatusInternalServerError, "error.html", nil)
	}

	movie, err := db.FunnelDAL.GetMovieFromKey(session.MovieKey)
	if err != nil {
		slog.Error(err.Error())
		return c.Render(http.StatusInternalServerError, "error.html", nil)
	}

	session.AllRatings = db.FunnelDAL.GetRatings()
	session.Movie = *movie
	return c.Render(http.StatusOK, "index.html", session)
}
