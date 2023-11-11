package model


type SessionStatus string

const (
	SESSION_INIT = "SESSION_INIT"
	SESSION_STARTED = "SESSION_STARTED"
	SESSION_OVER = "SESSION_OVER"
)

type Movie struct {
	Name string `json:",omitempty"`
	ID int `json:",omitempty"`
	Poster string `json:",omitempty"`
	Description string `json:",omitempty"`
	ReleaseDate string `json:",omitempty"`
}

type Tile struct {
	Penalty int 
	Selected bool
	Flipped bool 
	I int 
	J int 
}

type UserData struct {
	HardModeEnabled bool
}

type Mural struct {
	Game Game
	Session Session 
	UserStats UserStats 
	UserData UserData 
}

type Board struct {
	Size int `json:",omitempty"`
	Tiles [][]Tile `json:",omitempty"`
}

type Answer struct {
	Movie `json:",omitempty"`
	IsCorrect bool
	Selected bool
}

type SessionStats struct {
	Date string
	Score int
	Shareable string
}

type UserStats struct {
	UserKey string
	CurrentStreak int
	GamesPlayed int
	LongestStreak int
	BestScore int
	LastGame int
}

type Session struct {
	UserKey string 

	// current metadata
	CurrentScore int
	Board Board 

	// current state
	SelectedTile *Tile `json:",omitempty"`

	// answer data
	InputAnswer string `json:",omitempty"`

	SelectedAnswer *Answer `json:",omitempty"`
	SubmittedAnswer *Answer `json:",omitempty"`
	GameWon bool

	SessionStatus SessionStatus 
	SessionStats SessionStats
}

func NewSession(
	user_key string,
) Session {
	board := NewGameBoard(10)

	return Session{
		UserKey: user_key,
		Board: board,
		CurrentScore: 100,
		SessionStatus: SESSION_INIT,
	}
}


// this needs an even number to be populated
func populateTileZones(
	max int, 
	size int, 
	level int,
	all_tiles [][]Tile, 
) (int, int, [][]Tile) {
	// break case
	if size == 0 {
		return max, 0, all_tiles
	}

	// lets get all of the tiles
	i := 0
	for i < size {
		j := 0
		for j < size {
			// left
			penalty := (level + 1)  * 3
			all_tiles[i + level][level] = Tile{
				I: i + level,
				J: level,
				Penalty: penalty,
				Selected: false,
				Flipped: false,
			};

			// top
			all_tiles[level][j + level] = Tile{
				I: level,
				J: j + level,
				Penalty: penalty,
				Selected: false,
				Flipped: false,
			};

			// bottom
			all_tiles[max -  1 - level][j + level] = Tile{
				I: max -  1 - level,
				J: j + level,
				Penalty: penalty,
				Selected: false,
				Flipped: false,
			};

			// right
			all_tiles[i + level][max - 1 - level] = Tile{
				I: i + level,
				J: max -  1 - level,
				Penalty: penalty,
				Selected: false,
				Flipped: false,
			};

			j+= 1

		}
		i+= 1
	}

	return populateTileZones(max, size - 2, level + 1, all_tiles)
}

func NewTiles(size int) ([][]Tile) {
	// need to make tiles
	tiles := make([][]Tile, size)

	// need to make rows
	for i := range tiles {
		tiles[i] = make([]Tile, size)
	}

	_, _, tiles =  populateTileZones(size, size, 0, tiles)

	return tiles
}

func NewGameBoard(size int) (Board) {
	return Board{
		Size: size,
		Tiles: NewTiles(size),
	}
}


type Game struct {
	NumberOfSessions int
	GameKey int
	Date string
	CorrectAnswer Answer
	Answers []Answer
	IsCurrent bool
}