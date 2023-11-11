package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"funnel/config"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

var (
	ErrCouldNotGetUserKey = fmt.Errorf("could not get user key")
	Store                 *sessions.CookieStore
)

func InitSession() {
	Store = sessions.NewCookieStore([]byte(os.Getenv(config.EnvSessionKey)))
}

func GetUserKeyFromContext(c echo.Context) string {
	user_session, err := Store.Get(c.Request(), "user-session")
	if err != nil {
		user_session = sessions.NewSession(Store, "user-session")
		user_session.Values["user-key"] = fmt.Sprintf("%v", uuid.New())
	}

	user_key, ok := user_session.Values["user-key"]
	if !ok {
		user_session = sessions.NewSession(Store, "user-session")
		user_session.Values["user-key"] = fmt.Sprintf("%v", uuid.New())
	}

	return user_key.(string)
}

// Process is the middleware function.
func GetUserKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user_session, err := Store.Get(c.Request(), "user-session")
		if err != nil {
			user_session = sessions.NewSession(Store, "user-session")
			user_session.Values["user-key"] = fmt.Sprintf("%v", uuid.New())
		}

		_, ok := user_session.Values["user-key"]
		if !ok {
			user_session.Values["user-key"] = fmt.Sprintf("%v", uuid.New())
		}

		Store.Save(c.Request(), c.Response().Writer, user_session)
		return next(c)
	}
}

func HashString(data string) string {
	// Calculate the SHA-256 hash
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
