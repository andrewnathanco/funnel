package sql

import (
	"database/sql"
	"funnel/db"
	"funnel/model"
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
	db := &SQLiteDAL{ DB: database, }
	err = db.InitDB()

	// setup schema
	return db, err
}

func (dal *SQLiteDAL) PingDatabse() (error)  {
	return dal.DB.Ping()
}

func (dal *SQLiteDAL) InitDB() (error) {
	_, err := dal.DB.Exec(createBlackList)
	if err != nil {
		return err
	}

	_, err = dal.DB.Exec(createGreenList)
	if err != nil {
		return err
	}

	_, err = dal.DB.Exec(createYellowList)
	if err != nil {
		return err
	}

	_, err = dal.DB.Exec(createFunnelSessions)
	if err != nil {
		return err
	}

	return nil
}

func (dal *SQLiteDAL) GetRandomMovieByDecade(decade string) (*tmdb.MovieShort, error)  {
	return getMovieByDecade(decade, dal)
}

func (dal *SQLiteDAL) GetBoardByKey(key string) (*model.FunnelBoard, error) {
	var board *model.FunnelBoard
	var err error

	board, err = getSessionByKey(key, dal)
	if err != sql.ErrNoRows {
		if err != nil  {
			return nil, err
		}
	}

	if err == sql.ErrNoRows {
		new_board := model.NewFunnelBoard()
		movie, err := dal.GetRandomMovieByDecade("2020s")
		if err != nil {
			return nil, err
		}

		new_board.Movie = *movie
		board = &new_board
		err = insertBoard(key, board, dal)
		if err != nil {
			return nil, err
		}
	}

	return board, nil
}

func (dal *SQLiteDAL) SetBoardByKey(user_key string, board model.FunnelBoard) error {
	return insertBoard(user_key, &board, dal)
}

func (dal *SQLiteDAL) GetNumberPerDecade(decade string) (int) {
	return getNumberPerDecade(decade, dal)
}

func (dal *SQLiteDAL) InsertMoveIntoList(list string, movie_key string) (error) {
	return insertIntoList(list,movie_key, dal)
}