package db

import (
	"fmt"

	"github.com/ryanbradynd05/go-tmdb"
)

var (
	MuralDAL IDAL
	BlacklistDAL IDAL

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

	// session stuf
	InitDB() (error) 
	GetRadomMovieByDecade(decade string) (*tmdb.MovieShort, error) 
}
