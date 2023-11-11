package model

import "github.com/ryanbradynd05/go-tmdb"


type ListType string

const (
	GREEN_LIST = "GREEN_LIST"
	YELLOW_LIST = "YELLOW_LIST"
	RED_LIST = "RED_LIST"
)

type FunnelBoard struct {
	Movie []tmdb.MovieShort
}
