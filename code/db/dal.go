package db

import (
	"fmt"
	"funnel/model"

	"github.com/ryanbradynd05/go-tmdb"
)

var (
	MuralDAL IDAL

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
	InitDB() (error) 
	GetRandomMovieByDecade(decade string) (*tmdb.MovieShort, error) 
	GetNumberPerDecade(decade string) (int)
	InsertMoveIntoList(list string, movie_key string) (error)
	
	// session stuff
	GetBoardByKey(key string) (*model.FunnelBoard, error) 
	SetBoardByKey(key string, session model.FunnelBoard) (error) 
}
