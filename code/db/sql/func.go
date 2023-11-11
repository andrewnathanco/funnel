package sql

import (
	"os"

	"github.com/ryanbradynd05/go-tmdb"
)

func createFileIfNotExists(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		// File does not exist, so create it
		_, err = os.Create(filename)
		return err
	}
	return nil
}

func getMovieByDecade(decade string) (*tmdb.MovieShort, error) {
	return nil, nil
}

func initDB() (*tmdb.MovieShort, error) {
	return nil, nil
}