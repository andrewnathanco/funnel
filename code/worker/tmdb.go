package worker

import (
	"fmt"
	"funnel/controller/movie"
	"funnel/db"
	"log/slog"
)

type TMDBWorker struct{}

func NewTMDBWorker() TMDBWorker {
	return TMDBWorker{}
}

func (tw TMDBWorker) CacheAnswers() {
	slog.Info("Caching Answers From TMDB")
	movie_controller := movie.NewTMDBController()

	// get current page
	funnel_meta, err := db.FunnelDAL.GetFunnelMeta()
	if err != nil {
		slog.Error(fmt.Errorf("could not get current movie page: %w", err).Error())
		return
	}

	decades := []string{
		"2020",
		"2010",
		"2000",
		"1990",
		"1980",
		"1970",
		"1960",
		"1950",
		"1940",
	}

	for _, decade := range decades {
		tmdb_movies, err := movie_controller.GetTMDBMoviesByDecade(funnel_meta.CurrentTMDBPage+1, decade)
		if err != nil {
			slog.Error(fmt.Errorf("could not get movies: %w", err).Error())
			return
		}

		// cache answers
		var movies []db.MovieShort
		for _, tmdb_movie := range tmdb_movies {
			movie := db.ConvertTMDBToFunnel(tmdb_movie)
			movies = append(movies, movie)
		}

		err = db.FunnelDAL.SaveMovies(movies)
		if err != nil {
			slog.Error(fmt.Errorf("could not cache answers: %w", err).Error())
			return
		}
	}

	funnel_meta.CurrentTMDBPage += 1
	err = db.FunnelDAL.SaveFunnelMeta(*funnel_meta)
	if err != nil {
		slog.Error(fmt.Errorf("could not update meta: %w", err).Error())
		return
	}
}
