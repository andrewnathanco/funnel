package movie

import (
	"fmt"
	"funnel/config"
	"os"

	"github.com/ryanbradynd05/go-tmdb"
)

type TMDBController struct {
	TMDBApi *tmdb.TMDb
}

func NewTMDBAPI() *tmdb.TMDb {
	config := tmdb.Config{
		APIKey:   os.Getenv(config.EnvTMDBKey),
		Proxies:  nil,
		UseProxy: false,
	}

	return tmdb.Init(config)
}

func NewTMDBController() *TMDBController {
	api := NewTMDBAPI()
	return &TMDBController{
		TMDBApi: api,
	}
}

func (mc TMDBController) GetAnswers(
	page_number int,
) ([]tmdb.MovieShort, error) {
	parameters := map[string]string{
		"page": fmt.Sprintf("%d", page_number),
	}

	movie_results, err := mc.TMDBApi.DiscoverMovie(parameters)
	if err != nil {
		return nil, err
	}

	return movie_results.Results, nil
}

func (mc TMDBController) GetTMDBMoviesByDecade(
	page_number int,
	decade string,
) ([]tmdb.MovieShort, error) {
	lower_bound, upper_boud := getDecadeBounds(decade)

	parameters := map[string]string{
		"page":                     fmt.Sprintf("%d", page_number),
		"primary_release_date.lte": upper_boud,
		"primary_release_date.gte": lower_bound,
	}

	movie_results, err := mc.TMDBApi.DiscoverMovie(parameters)
	if err != nil {
		return nil, err
	}

	return movie_results.Results, nil
}
