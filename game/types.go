package game

import "strings"

type Seed uint32

type ActionState uint16

type Direction int

const (
	DIR_LEFT  = -1
	DIR_RIGHT = 1
	DIR_OTHER = 0
)

type IntVector2 struct {
	X int8
	Y int8
}

func NewIntVec(x, y int8) IntVector2 { return IntVector2{X: x, Y: y} }

type Vector2 struct {
	X float32
	Y float32
}

func NewVec(x, y float32) Vector2 { return Vector2{X: x, Y: y} }

type Platform string

func ParsePlatform(plat string) Platform {
	switch strings.ToLower(plat) {
	case "dolphin":
		return Dolphin
	case "nintendont":
		return Nintendont
	case "console":
		return Console
	case "network":
		return Network
	default:
		return None
	}
}

const (
	Dolphin    Platform = "Dolphin"
	Nintendont Platform = "Nintendont"
	Network    Platform = "Network"
	Console    Platform = "Console"
	None       Platform = "None"
)
