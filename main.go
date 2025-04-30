package slp

import (
	"path/filepath"

	"go.dalton.dog/bark"

	"go.dalton.dog/slp/file"
	"go.dalton.dog/slp/game"
)

// LoadFileFromPath will return a pointer to the file.File struct
// that was loaded and parsed from the slp file at the given filepath.
func LoadFileFromPath(filepath string) (*file.File, error) {
	return file.LoadFile(filepath)
}

// LoadGameFromPath will return a pointer to the game.Game struct
// that was loaded and parsed from the slp file at the given filepath.
func LoadGameFromPath(filepath string) (*game.Game, error) {
	file, err := file.LoadFile(filepath)
	if err != nil {
		return nil, err
	}

	return game.NewGameFromFile(file)
}

// LoadGameFromPath will return a pointer to the game.Game struct
// that was loaded and parsed from the given file.File struct.
func LoadGameFromFile(file *file.File) (*game.Game, error) {
	return game.NewGameFromFile(file)
}

func main() {
	bark.Init(bark.BarkOptions{})
	bark.SetDebugLevel(true)
	// samples, err := os.ReadDir("samples")
	// if err != nil {
	// 	log.Fatalf("Unable to read dir 'samples': %v", err)
	// }
	//
	// for _, entry := range samples {
	// 	loadEntry(entry.Name())
	// }
	loadEntry("Slippi_1.slp")
	loadEntry("Genesis_1.slp")
}

func loadEntry(filename string) {
	if filepath.Ext(filename) != ".slp" {
		return
	}

	path := filepath.Join("samples", filename)

	file, err := file.LoadFile(path)
	if err != nil {
		bark.Errorf("Error processing file (%v): %v", filename, err)
		return
	}

	// log.Printf("Success! Raw len: %v -- Metadata len: %v", len(file.Raw.Bytes), len(file.Metadata.Bytes))
	bark.Infof("File %v\n%v\n%v", file.Filepath, file.Metadata, file.Raw)
}
