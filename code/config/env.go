package config

import (
	"fmt"
	"log/slog"
	"os"
)

func ValidateENV() error {
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
