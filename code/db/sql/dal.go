package sql

import (
	"database/sql"
	"fmt"
	"funnel/db"
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
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

type SQLiteDAL struct {
	DB *sqlx.DB
}

func NewSQLiteDal(file string) (*SQLiteDAL, error) {
	err := createFileIfNotExists(file)
	if err != nil {
		slog.Error(err.Error())
		return nil, db.ErrCreateDatabaseFile
	}

	database, err := sqlx.Open("sqlite3", file)
	if err != nil {
		slog.Error(err.Error())
		return nil, db.ErrConnectToDatabase
	}

	err = database.Ping()
	if err != nil {
		return nil, db.ErrPingDatabase
	}
	db := &SQLiteDAL{DB: database}
	err = db.InitDB()

	// setup schema
	return db, err
}

func (dal *SQLiteDAL) PingDatabse() error {
	return dal.DB.Ping()
}

func (dal *SQLiteDAL) InitDB() error {
	_, err := dal.DB.Exec(createMoviesTable)
	if err != nil {
		return err
	}

	_, err = dal.DB.Exec(createMetaTable)
	if err != nil {
		return err
	}

	_, err = dal.DB.Exec(createSessionTable)
	if err != nil {
		return err
	}

	return nil
}

func (dal *SQLiteDAL) GetFunnelMeta() (*db.FunnelMeta, error) {
	rows, err := dal.DB.Queryx("select * from funnel_meta")
	if err != nil {
		return nil, err
	}

	var funnel_meta db.FunnelMeta
	for rows.Next() {
		err := rows.StructScan(&funnel_meta)
		if err != nil {
			return nil, err
		}
	}

	return &funnel_meta, nil
}

func (dal *SQLiteDAL) SaveFunnelMeta(meta db.FunnelMeta) error {
	_, err := dal.DB.NamedExec(updateMeta, meta)
	return err
}

func (dal *SQLiteDAL) SaveMovies(movies []db.MovieShort) error {
	_, err := dal.DB.NamedExec(insertMovies, movies)
	return err
}

func (dal *SQLiteDAL) GetRandomMovie() (*db.MovieShort, error) {
	movie := db.MovieShort{}
	err := dal.DB.Get(&movie, getRandomMovie)
	return &movie, err
}

func (dal *SQLiteDAL) GetSessionForUser(user_key string) (*db.Session, error) {
	var session db.Session
	err := dal.DB.Get(&session, getSessionForUser, user_key)
	if err == sql.ErrNoRows {
		movie, rand_mov_err := dal.GetRandomMovie()
		if rand_mov_err != nil {
			slog.Error(fmt.Sprintf("random movie: %s", rand_mov_err.Error()))
			return nil, rand_mov_err
		}

		session = db.Session{
			UserKey:       user_key,
			MovieKey:      movie.ID,
			SessionStatus: db.SESSION_INIT,
			Rating:        5,
		}
	} else if err != nil {
		slog.Error(fmt.Sprintf("session: %s", err.Error()))
		return nil, err
	}

	return &session, nil
}

func (dal *SQLiteDAL) SaveSessionForUser(session db.Session) error {
	_, err := dal.DB.NamedExec(saveSessionForUser, session)
	return err
}

func (dal *SQLiteDAL) GetMovieFromKey(movie_key int) (*db.MovieShort, error) {
	movie := db.MovieShort{}
	err := dal.DB.Get(&movie, "select * from movies where id = ?", movie_key)
	return &movie, err
}
