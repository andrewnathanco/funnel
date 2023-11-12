package config

import (
	"fmt"
	"log/slog"
	"os"
)

func ValidateENV() error {
	funnel_db := os.Getenv(EnvFunnelDB)
	slog.Info("USING: " + funnel_db)
	if funnel_db == "" {
		return fmt.Errorf("need environment variable %s", EnvFunnelDB)
	}

	session_key := os.Getenv(EnvSessionKey)
	slog.Info("SESSION KEY: " + session_key)
	if session_key == "" {
		return fmt.Errorf("need environment variable %s", EnvSessionKey)
	}

	tmdb_key := os.Getenv(EnvTMDBKey)
	if tmdb_key == "" {
		return fmt.Errorf("need environment variable %s", EnvTMDBKey)
	}

	return nil
}
