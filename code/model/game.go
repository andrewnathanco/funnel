package model

import (
	"github.com/ryanbradynd05/go-tmdb"
)



const (
	GREEN_LIST = "GREEN_LIST"
	YELLOW_LIST = "YELLOW_LIST"
	BLACK_LIST = "BLACK_LIST"
)

type FunnelBoard struct {
	Movie tmdb.MovieShort
	SelectedList string
	MoviesLeft int
	Theme string
}

func NewFunnelBoard() (FunnelBoard) {
	return FunnelBoard{
		SelectedList: GREEN_LIST,
		Theme: "2020s",
	}
}