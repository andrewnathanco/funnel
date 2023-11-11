package funnel

import (
	"funnel/config"
	"os"
	"time"
)

func getVersion() string {
	return os.Getenv(config.EnvVersion)
}

func getMoviesLeft() int {
	return 1
}

func getDate() string {
	return time.Now().Format(time.RFC3339)
}

func getCurrentTheme() string {
	return "1980s"
}