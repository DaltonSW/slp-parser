package types

import (
	"log"
	"os"

	"go.dalton.dog/slp/utils"
)

type Game struct {
	Filepath string
	Raw      Raw
	Metadata Metadata
}

func LoadGame(filepath string) *Game {
	if utils.DoesFileNotExist(filepath) {
		return nil
	}

	game := &Game{
		Filepath: filepath,
	}

	// Load file data and validate loaded information
	fileData, err := os.ReadFile(filepath)
	if err != nil {
		log.Default().Fatal(err)
	}
	if len(fileData) < 1 {
		log.Default().Fatal("File Data loaded is empty")
	}

	// Open file
	// Read bytes in

	game.Raw = LoadRaw(fileData)
	game.Metadata = LoadMetadata(fileData)

	return game
}
