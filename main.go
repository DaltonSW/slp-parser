package main

import (
	"log"
	// "os"
	"path/filepath"

	"go.dalton.dog/slp/types"
)

func main() {
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

	file, err := types.LoadFile(path)
	if err != nil {
		log.Printf("Error processing file %v: %v", filename, err)
		return
	}

	// log.Printf("Success! Raw len: %v -- Metadata len: %v", len(file.Raw.Bytes), len(file.Metadata.Bytes))
	log.Printf("File %v\n%v", file.Filepath, file.Metadata)
}
