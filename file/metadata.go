package file

import (
	"fmt"

	"github.com/jmank88/ubjson"
	"go.dalton.dog/bark"
)

func LoadMetadata(stream []byte) (*Metadata, error) {
	bark.Debug("Starting to load metadata bytes")
	meta := &Metadata{}

	if err := ubjson.Unmarshal(stream, meta); err != nil {
		return nil, err
	}

	return meta, nil
}

type Metadata struct {
	// StartAt: ISO 8601 formatted string representing the timestamp that the game started at
	StartAt string `ubjson:"startAt"`

	// LastFrame: Number of the last frame of the game. Can be used to determine game length without scrubbing
	LastFrame int32 `ubjson:"lastFrame"`

	// Players: Metadata for each player in the game
	Players map[string]PlayerMetadata `ubjson:"players"`

	// PlayedOn: Platform that the game was played on
	PlayedOn string `ubjson:"playedOn"`

	// ConsoleNick: Name of the console that the reply was created on
	ConsoleNick string `ubjson:"consoleNick"`
}

func (m Metadata) String() string {
	out := "~~ Metadata ~~\n"
	out = out + fmt.Sprintf("Started At   : %v\n", m.StartAt)
	out = out + fmt.Sprintf("Last Frame   : %v\n", m.LastFrame)
	out = out + fmt.Sprintf("Played On    : %v\n", m.PlayedOn)
	out = out + fmt.Sprintf("Console Nick : %v\n", m.ConsoleNick)
	out = out + fmt.Sprintf("Players      : %v\n", m.Players)
	// out = out + "Slots:\n"
	// out = out + fmt.Sprintf("  Port 1: %v\n", m.Players[0])
	// out = out + fmt.Sprintf("  Port 2: %v\n", m.Players[1])
	// out = out + fmt.Sprintf("  Port 3: %v\n", m.Players[2])
	// out = out + fmt.Sprintf("  Port 4: %v\n", m.Players[3])

	return out
}

type PlayerMetadata struct {
	// Characters: Number of frames spent as each character. Generally only useful for Zelda/Sheik, methinks
	Characters map[string]int32 `ubjson:"characters"`

	// Names: Contains the display name of the player, as well as the connect code if applicable
	Names map[string]string `ubjson:"names"`
}

func (pm PlayerMetadata) String() string {
	if len(pm.Characters) == 0 {
		return "NONE"
	}

	return fmt.Sprintf("%v %v", pm.Names, pm.Characters)
}
