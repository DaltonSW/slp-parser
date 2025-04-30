package game

import (
	"fmt"
	"strconv"
	"time"

	"go.dalton.dog/slp/file"
	"go.dalton.dog/slp/id"
)

func ParseGameMetadata(fileMetadata *file.Metadata) (*Metadata, error) {
	outMeta := &Metadata{}

	startAt, err := time.Parse(time.RFC3339Nano, fileMetadata.StartAt)
	if err != nil {
		return nil, err
	} else {
		outMeta.StartAt = startAt
	}

	outMeta.FinalFrame = int(fileMetadata.LastFrame)
	outMeta.ConsoleNick = fileMetadata.ConsoleNick
	outMeta.PlayedOn = ParsePlatform(fileMetadata.ConsoleNick)

	outMeta.Players = make([]PlayerMetadata, 4)

	for idx, data := range fileMetadata.Players {
		idxInt, _ := strconv.Atoi(idx)
		playerData := PlayerMetadata{
			Port: idxInt,
		}

		err := playerData.ParseData(data)
		if err != nil {
			return nil, err
		}

		outMeta.Players[playerData.Port] = playerData
	}

	return outMeta, nil
}

type Metadata struct {
	StartAt     time.Time
	FinalFrame  int
	PlayedOn    Platform
	ConsoleNick string

	Players []PlayerMetadata
}

func (m Metadata) String() string {
	out := "~~ Metadata ~~\n"
	out = out + fmt.Sprintf("Started At   : %v\n", m.StartAt)
	out = out + fmt.Sprintf("Last Frame   : %v\n", m.FinalFrame)
	out = out + fmt.Sprintf("Played On    : %v\n", m.PlayedOn)
	out = out + fmt.Sprintf("Console Nick : %v\n", m.ConsoleNick)

	out = out + "Ports:\n"
	out = out + fmt.Sprintf("%v\n", m.Players[0])
	out = out + fmt.Sprintf("%v\n", m.Players[1])
	out = out + fmt.Sprintf("%v\n", m.Players[2])
	out = out + fmt.Sprintf("%v\n", m.Players[3])

	return out
}

type PlayerMetadata struct {
	// 0-indexed port
	Port int

	// Info about the characters the player spent the game as.
	// This will typically contain only 1 entry, but swapping
	// between Zelda and Sheik can cause 2 entries.
	Characters []PlayerCharacter

	// Display name of the player
	NetplayName string

	// Slippi code
	SlippiCode string
}

func (pm PlayerMetadata) String() string {
	out := fmt.Sprintf("  Port %d: ", pm.Port+1)

	if pm.SlippiCode != "" {
		out += fmt.Sprintf("[%v] ", pm.SlippiCode)
	}

	if pm.NetplayName != "" {
		out += fmt.Sprintf("(%v) ", pm.NetplayName)
	}

	out += fmt.Sprintf("Played %v", pm.Characters)

	return out
}

func (pm *PlayerMetadata) ParseData(data file.PlayerMetadata) error {
	pm.NetplayName = data.Names["netplay"]
	pm.SlippiCode = data.Names["code"]
	pm.Characters = make([]PlayerCharacter, 0)

	for char, frames := range data.Characters {
		charInt, err := strconv.Atoi(char)
		if err != nil {
			return err
		}
		newChar := PlayerCharacter{
			CharacterID: id.CharacterInGameID(charInt),
			NumFrames:   int(frames),
		}
		pm.Characters = append(pm.Characters, newChar)
	}

	return nil
}

type PlayerCharacter struct {
	// In-Game ID of the character
	CharacterID id.CharacterInGameID

	// Number of frames the player spent as this character
	NumFrames int
}
