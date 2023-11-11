package funnel

import (
	"fmt"
	"funnel/config"
	"os"
	"time"

	"github.com/ryanbradynd05/go-tmdb"
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

func getThemes() []string {
	return []string{
		"2020s",
		"2010s",
		"2000s",
		"1990s",
		"1980s",
		"1970s",
		"1960s",
		"1950s",
		"1940s",
		"1930s",
		"1920s",
		"1910s",
		"1900s",
		"All",
	}
}

type option struct {
	Selected bool
}

func newOption(selected bool) option {
	return option{Selected: selected}
}

func getReleaseYear(movie tmdb.MovieShort) string {
	// we should be able to trust this, not just put an empty string
	layout := "2006-01-02"
	release_date, err := time.Parse(layout, movie.ReleaseDate)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", release_date.Year())
}