package types

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"

	"go.dalton.dog/slp/utils"
)

type File struct {
	Filepath string
	Raw      *Raw
	Metadata *Metadata
}

func LoadFile(filepath string) (*File, error) {
	noFile, err := utils.DoesFileNotExist(filepath)
	if noFile {
		return nil, err
	}

	file := &File{
		Filepath: filepath,
	}

	fileData, err := os.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	if len(fileData) < 1 {
		return nil, errors.New("Loaded file was empty")
	}

	// rawBytes, metadataBytes, err := splitFileBytes(fileData)
	_, metadataBytes, err := splitFileBytes(fileData)

	if err != nil {
		return nil, err
	}

	// raw, err := LoadRaw(rawBytes)
	// file.Raw = LoadRaw(rawBytes)
	metadata, err := LoadMetadata(metadataBytes)
	if err != nil {
		return nil, err
	}
	file.Metadata = metadata

	return file, nil
}

func splitFileBytes(stream []byte) ([]byte, []byte, error) {
	// For file structures: An ASCII char in brackets will be the ASCII value, an integer will be that value in hex ([3] == 0x03 ; [a] == 0x61)

	trimmedStream, rawLen, err := stripRawHeader(stream)

	if err != nil {
		return nil, nil, err
	}

	raw := trimmedStream[:rawLen]
	remainder := trimmedStream[rawLen:]

	metadata, err := stripMetadataHeader(remainder)
	if err != nil {
		return nil, nil, err
	}

	return raw, metadata, nil
}

func stripRawHeader(stream []byte) ([]byte, int, error) {
	// Index in the byte stream
	offset := 0

	// Validating [{][U][3][r][a][w]

	if stream[offset] != '{' {
		return nil, 0, errors.New("Expected UBJSON object start")
	}
	offset++

	if stream[offset] != 'U' {
		return nil, 0, errors.New("Expected 'U' type marker for raw key length")
	}
	offset++

	// keyLen should always be 3 for 'raw'
	keyLen := int(stream[offset])
	if keyLen != 3 {
		return nil, 0, fmt.Errorf("Expected first key length of 3, got %v", keyLen)
	}
	offset++

	if offset+keyLen > len(stream) {
		return nil, 0, errors.New("Unexpected EOF while trying to read 'raw' key")
	}
	key := string(stream[offset : offset+keyLen])

	if key != "raw" {
		return nil, 0, fmt.Errorf("Expected 'raw' key, got %v", key)
	}
	offset += keyLen

	// Starting key looks good!

	// Validating [[][$][U][#][l][x][x][x][x]

	if stream[offset] != '[' {
		return nil, 0, errors.New("Expected '[' after 'raw' key to indicate array start ")
	}
	offset++

	if stream[offset] != '$' || stream[offset+1] != 'U' {
		return nil, 0, errors.New("Expected optimized array of type $U to start 'raw' key")
	}
	offset += 2

	if stream[offset] != '#' || stream[offset+1] != 'l' {
		return nil, 0, errors.New("Expected array length marker #l for 'raw' key")
	}
	offset += 2

	if offset+4 > len(stream) {
		return nil, 0, errors.New("Unexpected EOF while trying to read 'raw' array length")
	}
	rawLen := int(binary.BigEndian.Uint32(stream[offset : offset+4]))
	offset += 4

	if offset+rawLen > len(stream) {
		return nil, 0, errors.New("Unexpected EOF while trying to read 'raw' data")
	}

	return stream[offset:], rawLen, nil
}

func stripMetadataHeader(stream []byte) ([]byte, error) {
	// Index in the byte stream
	offset := 0

	// Validating [U][8][m][e][t][a][d][a][t][a]

	if stream[offset] != 'U' {
		return nil, errors.New("Expected 'U' type marker for 'metadata' key length")
	}
	offset++

	// keyLen should always be 8 for 'metadata'
	keyLen := int(stream[offset])
	if keyLen != 8 {
		return nil, fmt.Errorf("Expected 'metadata' key length of 8, got %v", keyLen)
	}
	offset++

	if offset+keyLen > len(stream) {
		return nil, errors.New("Unexpected EOF while trying to read 'metadata' key")
	}
	key := string(stream[offset : offset+keyLen])

	if key != "metadata" {
		return nil, fmt.Errorf("Expected 'metadata' key, got %v", key)
	}
	offset += keyLen

	// Metadata key looks good!
	// Remainder should be "{ <metadata stuff> }"

	return stream[offset:], nil
}
