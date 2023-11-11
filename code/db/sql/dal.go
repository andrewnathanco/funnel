package sql

import (
	"database/sql"
	"funnel/db"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ryanbradynd05/go-tmdb"
)

type SQLiteDAL struct {
	DB *sql.DB
}

func NewSQLiteDal(file string) (*SQLiteDAL, error) {
	err := createFileIfNotExists(file)
	if err != nil {
		slog.Error(err.Error())
		return nil, db.ErrCreateDatabaseFile
	}

	database, err := sql.Open("sqlite3", file)
	if err != nil {
		slog.Error(err.Error())
		return nil, db.ErrConnectToDatabase
	}

	err = database.Ping()
	if err != nil {
		return nil, db.ErrPingDatabase
	}

	// setup schema
	return &SQLiteDAL{
		DB: database,
	}, nil
}

func (dal *SQLiteDAL) PingDatabse() (error)  {
	return dal.DB.Ping()
}

func (dal *SQLiteDAL) InitDB() (error) {
	return nil
}

func (dal *SQLiteDAL) GetRadomMovieByDecade(decade string) (*tmdb.MovieShort, error)  {
	return getMovieByDecade(decade)
}