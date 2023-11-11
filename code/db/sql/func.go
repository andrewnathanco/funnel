package sql

import (
	"encoding/json"
	"funnel/model"
	"os"
	"strconv"

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

func getMovieByDecade(decade string, dal *SQLiteDAL) (*tmdb.MovieShort, error) {
	decade = GetSQLDecade(decade)
	// now lets get back the res// now lets get back the rest
	rows := dal.DB.QueryRow(getMovieByDecadeQuery, decade)
	var movie_data string
	var movie tmdb.MovieShort
	var avg_score string
	err := rows.Scan(&movie_data, &avg_score)
	if err != nil  {
		return nil, err
	}

	err = json.Unmarshal([]byte(movie_data), &movie)
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func getSessionByKey(key string, dal *SQLiteDAL) (*model.FunnelBoard, error) {
	// now lets get back the res// now lets get back the rest
	rows := dal.DB.QueryRow(getSessionByKeyQuery, key)
	var session_data string
	var session model.FunnelBoard
	err := rows.Scan(&session_data)
	if err != nil  {
		return nil, err
	}

	err = json.Unmarshal([]byte(session_data), &session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func insertBoard(
	user_key string,
	board *model.FunnelBoard,
	dal *SQLiteDAL,
) (error) {
	board_marshalled, err := json.Marshal(board)
	if err != nil {
		return err
	}


	_, err = dal.DB.Exec(upsertBoardQuery, user_key, string(board_marshalled))
	if err != nil {
		return err
	}

	return nil
}

func GetSQLDecade(decade string) string {
	if decade == "All" {
		decade = "%"
	} else {
		if len(decade) > 1 {
			decade = decade[0:len(decade)-1]
			decade = replaceLastCharacter(decade, '%')
		}
	}
	
	return decade
}

func replaceLastCharacter(
	inputString string, 
	newChar rune,
) string {
	if len(inputString) == 0 {
		return inputString // Return the original string if it's empty
	}

	// Convert the string to a rune slice to work with individual characters
	strRunes := []rune(inputString)

	// Update the last character
	strRunes[len(strRunes)-1] = newChar

	// Convert the rune slice back to a string
	return string(strRunes)
}


func insertIntoList(list string, movie_key string, dal *SQLiteDAL) (error) {
	// now lets get back the res// now lets get back the rest
	if list == model.GREEN_LIST {
		_, err := dal.DB.Exec(greenListMovie, movie_key)
		if err != nil {
			return err
		}
	
		return nil
	}

	if list == model.YELLOW_LIST {
		_, err := dal.DB.Exec(yellowListMovie, movie_key)
		if err != nil {
			return err
		}
	
		return nil
	}

	if list == model.BLACK_LIST {
		_, err := dal.DB.Exec(blackListMovie, movie_key)
		if err != nil {
			return err
		}
	
		return nil
	}

	return nil
} 

func getNumberPerDecade(decade string, dal *SQLiteDAL) (int) {
	// now lets get back the res// now lets get back the rest
	rows := dal.DB.QueryRow(getNumberPerDecadeQuery, decade)
	var number_data string
	var answer_data string
	rows.Scan(&number_data, &answer_data)
	number, _ := strconv.ParseInt(number_data, 10, 64)
	return int(number)
} 