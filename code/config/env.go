package config

import (
	"fmt"
	"log/slog"
	"os"
)

func ValidateENV() error {
	blacklist_db := os.Getenv(EnvFunnelDB)
	slog.Info("USING: " + blacklist_db)
	if blacklist_db == "" {
		return fmt.Errorf("need environment variable %s", EnvFunnelDB)
	}

	mural_db := os.Getenv(EnvMuralDB)
	slog.Info("USING: " + mural_db)
	if mural_db == "" {
		return fmt.Errorf("need environment variable %s", EnvMuralDB)
	}

	session_key := os.Getenv(EnvSessionKey)
	slog.Info("SESSION KEY: " + session_key)
	if session_key == "" {
		return fmt.Errorf("need environment variable %s", EnvSessionKey)
	}

	return nil
}
