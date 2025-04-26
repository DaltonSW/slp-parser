package game

import "go.dalton.dog/slp/file"

type Game struct {
	Start    GameStart
	End      GameEnd
	Metadata Metadata
	Frames   []Frame
}

func NewGameFromFile(file file.File) *Game {
	out := &Game{}

	return out
}
