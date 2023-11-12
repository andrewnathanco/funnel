package routes

import (
	"funnel/db"
	"funnel/middleware"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SubmitRating(c echo.Context) error {
	user_key := middleware.GetUserKeyFromContext(c)
	session, err := db.FunnelDAL.GetSessionForUser(user_key)
	if err != nil {
		slog.Error(err.Error())
		return c.Render(http.StatusInternalServerError, "error.html", nil)
	}

	if err != nil {
		return c.Render(http.StatusBadRequest, "index.html", session)
	}

	movie, err := db.FunnelDAL.GetRandomMovie()
	if err != nil {
		slog.Error(err.Error())
		return c.Render(http.StatusInternalServerError, "error.html", nil)
	}

	rating := db.Rating{
		Rating:   session.Rating,
		MovieKey: session.MovieKey,
	}

	db.FunnelDAL.SaveRating(rating)

	// now reset the session
	session.MovieKey = movie.ID
	session.Movie = *movie
	session.Rating = 5
	session.SessionStatus = db.SESSION_INIT
	session.AllRatings = db.FunnelDAL.GetRatings()
	db.FunnelDAL.SaveSessionForUser(*session)
	return c.Render(http.StatusOK, "index.html", session)
}
