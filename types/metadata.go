package types

type Metadata struct {
	Bytes []byte

	// StartAt: ISO 8601 formatted string representing the timestamp that the game started at
	StartAt string

	// LastFrame: Number of the last frame of the game. Can be used to determine game length without scrubbing
	LastFrame int

	// Players: Metadata for each player in the game
	Players []PlayerMetadata

	// PlayedOn: Platform that the game was played on
	PlayedOn string

	// ConsoleNick: Name of the console that the reply was created on
	ConsoleNick string
}

type PlayerMetadata struct {
	// Characters: Number of frames spent as each character. Generally only useful for Zelda/Sheik, methinks
	Characters map[uint8]uint8

	// Names: Contains the display name of the player, as well as the connect code if applicable
	Names map[string]string
}

func LoadMetadata(stream []byte) Metadata {
	meta := Metadata{Bytes: stream}

	return meta
}
