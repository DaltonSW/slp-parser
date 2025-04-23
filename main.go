package main

import (
	"github.com/charmbracelet/log"
	// "os"
	"path/filepath"

	"go.dalton.dog/slp/file"
)

func main() {
	log.SetLevel(log.DebugLevel)
	// samples, err := os.ReadDir("samples")
	// if err != nil {
	// 	log.Fatalf("Unable to read dir 'samples': %v", err)
	// }
	//
	// for _, entry := range samples {
	// 	loadEntry(entry.Name())
	// }
	loadEntry("Slippi_1.slp")
	// loadEntry("Genesis_1.slp")
}

func loadEntry(filename string) {
	if filepath.Ext(filename) != ".slp" {
		return
	}

	path := filepath.Join("samples", filename)

	file, err := file.LoadFile(path)
	if err != nil {
		log.Error("Error processing file", "file", filename, "err", err)
		return
	}

	// log.Printf("Success! Raw len: %v -- Metadata len: %v", len(file.Raw.Bytes), len(file.Metadata.Bytes))
	log.Printf("File %v\n%v\n%v", file.Filepath, file.Metadata, file.Raw)
}
