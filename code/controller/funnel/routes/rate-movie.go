package routes

import (
	"funnel/db"
	"funnel/middleware"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RateMovie(c echo.Context) error {
	rating_str := c.FormValue("rate-slider")

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

	session.Movie = *movie
	session.SessionStatus = db.SESSION_RATED
	rating, err := strconv.ParseInt(rating_str, 10, 64)
	if err != nil {
		return c.Render(http.StatusBadRequest, "index.html", session)
	}

	session.Rating = int(rating)
	session.AllRatings = db.FunnelDAL.GetRatings()
	db.FunnelDAL.SaveSessionForUser(*session)
	return c.Render(http.StatusOK, "index.html", session)
}
