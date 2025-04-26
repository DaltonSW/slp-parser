package events

import "fmt"

// Command Bytes
const (
	EventPayloadsByte     = 0x35
	GameStartByte         = 0x36
	PreFrameUpdateByte    = 0x37
	PostFrameUpdateByte   = 0x38
	GameEndByte           = 0x39
	FrameStartByte        = 0x3A
	ItemUpdateByte        = 0x3B // Max of 15 of these per frame
	FrameBookendByte      = 0x3C
	GeckoListByte         = 0x3D
	FountainPlatformsByte = 0x3F
	WhispyBlowDirByte     = 0x40
	StadiumTransformByte  = 0x41
	MessageSplitterByte   = 0x10
)

type Event interface {
	GetByte() byte
	String() string
}

type EventRaw interface {
	GetCommandByte() byte
	GetEventName() string
}

// ParseNextEvent will pass the given payload off to the appropriate event parser based on given commandByte
// func ParseNextEventRaw(commandByte byte, payload []byte) (*EventRaw, error) {
func ParseNextEventRaw(payload []byte) (EventRaw, error) {
	var outEvent EventRaw

	commandByte := payload[0]

	switch commandByte {
	case EventPayloadsByte:
		return nil, nil
	case GameStartByte:
		outEvent = &GameStartRaw{}
	case PreFrameUpdateByte:
		outEvent = &PreFrameRaw{}
	case PostFrameUpdateByte:
		outEvent = &PostFrameRaw{}
	case GameEndByte:
		outEvent = &GameEndRaw{}
	case FrameStartByte:
		outEvent = &FrameStartRaw{}
	case ItemUpdateByte:
		outEvent = &ItemUpdateRaw{}
	case FrameBookendByte:
		outEvent = &FrameBookendRaw{}
	case GeckoListByte:
		outEvent = &GeckoListRaw{}
	case FountainPlatformsByte:
		outEvent = &FountainPlatformRaw{}
	case WhispyBlowDirByte:
		outEvent = &WhispyBlowDirectionRaw{}
	case StadiumTransformByte:
		outEvent = &PokemonTransformRaw{}
	case MessageSplitterByte:
		outEvent = &MessageSplitRaw{}
	default:
		return nil, fmt.Errorf("Tried to parse event with unsupported cmdByte: %b", commandByte)
	}

	err := UnpackRawEvent(outEvent, payload)

	if err != nil {
		return nil, err
	}

	return outEvent, nil
}
