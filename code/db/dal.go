package db

import (
	"fmt"
)

var (
	FunnelDAL IDAL

	// errors
	ErrConnectToDatabase   = fmt.Errorf("could not connect to database")
	ErrCreateDatabaseFile  = fmt.Errorf("could not create database file")
	ErrPingDatabase        = fmt.Errorf("could not ping database")
	ErrSetupGameSchema     = fmt.Errorf("could not setup game schema")
	ErrGettingBoardFromDB  = fmt.Errorf("could not get board from db")
	ErrBoardNotFound       = fmt.Errorf("no board set")
	ErrCastingBoard        = fmt.Errorf("board not in correct format")
	ErrSettingCurrentGames = fmt.Errorf("could not set all boards")
)

type IDAL interface {
	PingDatabse() error

	// game stuff
	InitDB() error
	GetFunnelMeta() (*FunnelMeta, error)
	SaveMovies([]MovieShort) error
	SaveFunnelMeta(FunnelMeta) error
	GetRandomMovie() (*MovieShort, error)
	GetSessionForUser(user_key string) (*Session, error)
	GetMovieFromKey(movie_key int) (*MovieShort, error)
	SaveSessionForUser(session Session) error
	SaveRating(Rating) error
	GetRatings() int
}
