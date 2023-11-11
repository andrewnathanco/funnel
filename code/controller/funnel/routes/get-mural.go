package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetIndex(c echo.Context) error {
	// user_key := middleware.GetUserKeyFromContext(c)
	return c.Render(http.StatusOK, "index.html", nil)
}
