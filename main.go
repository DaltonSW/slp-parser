package main

import (
	"path/filepath"

	"go.dalton.dog/bark"
	"go.dalton.dog/slp/file"
)

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
