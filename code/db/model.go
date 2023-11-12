package db

import "github.com/ryanbradynd05/go-tmdb"

const (
	SESSION_INIT  = "SESSION_INIT"
	SESSION_RATED = "SESSION_RATED"
)

type Session struct {
	UserKey       string `db:"user_key"`
	MovieKey      int    `db:"current_movie"`
	Movie         MovieShort
	SessionStatus string `db:"session_status"`
	Rating        int    `db:"rating"`
	AllRatings    int
}

type MovieShort struct {
	ID            int     `db:"id" json:"id"`
	Title         string  `db:"title" json:"title"`
	OriginalTitle string  `db:"original_title" json:"original_title"`
	ReleaseDate   string  `db:"release_date" json:"release_date"`
	Overview      string  `db:"overview" json:"overview"`
	VoteAverage   float32 `db:"vote_average" json:"vote_average"`
	VoteCount     uint32  `db:"vote_count" json:"vote_count"`
	Popularity    float32 `db:"popularity" json:"popularity"`
	Adult         bool    `db:"adult" json:"adult"`
	Video         bool    `db:"video" json:"video"`
	BackdropPath  string  `db:"backdrop_path" json:"backdrop_path"`
	PosterPath    string  `db:"poster_path" json:"poster_path"`
}

func ConvertTMDBToFunnel(tmdbShort tmdb.MovieShort) MovieShort {
	return MovieShort{
		ID:            tmdbShort.ID,
		Title:         tmdbShort.Title,
		OriginalTitle: tmdbShort.OriginalTitle,
		ReleaseDate:   tmdbShort.ReleaseDate,
		Overview:      tmdbShort.Overview,
		VoteAverage:   tmdbShort.VoteAverage,
		VoteCount:     tmdbShort.VoteCount,
		Popularity:    tmdbShort.Popularity,
		Adult:         tmdbShort.Adult,
		Video:         tmdbShort.Video,
		BackdropPath:  tmdbShort.BackdropPath,
		PosterPath:    tmdbShort.PosterPath,
	}
}

type FunnelMeta struct {
	ID              int `db:"id"`
	CurrentTMDBPage int `db:"current_tmdb_page"`
}

type Rating struct {
	MovieKey int `db:"movie_key"`
	Rating   int `db:"rating"`
}
