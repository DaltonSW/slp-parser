package types

import (
	// "bytes"
	"fmt"
	// "log"
	// "strconv"

	"github.com/jmank88/ubjson"
)

// HACK: Try and figure out how to make this decode / unmarshal in a proper way before moving on to `raw`

type Metadata struct {
	// Bytes []byte

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
	Names PlayerNames `ubjson:"names"`
}

func (pm PlayerMetadata) String() string {
	if len(pm.Characters) == 0 {
		return "NONE"
	}

	return fmt.Sprintf("%v %v", pm.Names, pm.Characters)
}

type PlayerCharacter struct {
	// CharacterID: ID of the character
	CharacterID string

	// NumFrames: Number of frames the player spent as this character
	NumFrames int32
}

func (pc PlayerCharacter) String() string {
	return fmt.Sprintf("Played %v for %v frames", pc.CharacterID, pc.NumFrames)
}

type PlayerNames struct {
	// Netplay: Display name of the player
	Netplay string `ubjson:"netplay"`

	// Code: Slippi code
	Code string `ubjson:"code"`
}

func (pn PlayerNames) String() string {
	return fmt.Sprintf("%v [%v]", pn.Netplay, pn.Code)
}

func LoadMetadata(stream []byte) (*Metadata, error) {
	meta := &Metadata{}

	if err := ubjson.Unmarshal(stream, meta); err != nil {
		return nil, err
	}

	return meta, nil

	// meta.StartAt = rawMap["startAt"].(string)
	// meta.LastFrame = rawMap["lastFrame"].(int32)
	// meta.PlayedOn = rawMap["playedOn"].(string)
	//
	// if rawMap["consoleNick"] != nil {
	// 	meta.ConsoleNick = rawMap["consoleNick"].(string)
	// }
	//
	// rawPlayers := rawMap["players"].(map[string]any)
	//
	// for _, playerData := range rawPlayers {
	// 	// slotInt, err := strconv.Atoi(slot)
	// 	// if err != nil {
	// 	// 	log.Default().Printf("Invalid slot index")
	// 	// 	continue
	// 	// }
	//
	// 	player := PlayerMetadata{
	// 		// Characters: make(map[string]PlayerCharacter),
	// 	}
	// 	playerMap := playerData.(map[string]any)
	//
	// 	characterMap := playerMap["characters"].(map[string]any)
	//
	// 	for id, frames := range characterMap {
	// 		pc := PlayerCharacter{}
	//
	// 		pc.CharacterID = id
	// 		pc.NumFrames = frames.(int32)
	//
	// 		// player.Characters = append(player.Characters, pc)
	// 	}
	//
	// 	nameMap := playerMap["names"].(map[string]any)
	// 	player.Names.Code = nameMap["code"].(string)
	// 	player.Names.Netplay = nameMap["netplay"].(string)
	//
	// 	// if nameMap["code"] != nil {
	// 	// 	player.Names.Code = nameMap["code"]
	// 	// }
	//
	// 	// meta.Players[slotInt] = player
	// }
	//
	// return meta, nil
}
