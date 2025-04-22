package main

import (
	"log"
	"os"
	"path/filepath"

	"go.dalton.dog/slp/types"
)

func main() {
	samples, err := os.ReadDir("samples")
	if err != nil {
		log.Fatalf("Unable to read dir 'samples': %v", err)
	}

	for _, entry := range samples {
		if filepath.Ext(entry.Name()) != ".slp" {
			continue
		}

		path := filepath.Join("samples", entry.Name())

		file, err := types.LoadFile(path)
		if err != nil {
			log.Printf("Error processing file %v: %v", file.Filepath, err)
			continue
		}

		// log.Printf("Success! Raw len: %v -- Metadata len: %v", len(file.Raw.Bytes), len(file.Metadata.Bytes))
		log.Printf("File %v -- Metadata: %v", file.Filepath, file.Metadata.Bytes)
	}
}
