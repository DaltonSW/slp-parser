package events

import (
	"encoding/binary"
	"fmt"
	"github.com/charmbracelet/log"
)

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

type EventPayloads struct {
	Payload  []byte
	Mappings map[byte]uint16
}

func (p EventPayloads) GetByte() byte {
	return EventPayloadsByte
}

func (p EventPayloads) String() string {
	return fmt.Sprint(p.Payload)
}

func (p EventPayloads) GetPayloadLength(cmdByte byte) (uint16, error) {
	log.Debugf("Getting payload length for byte %X", cmdByte)
	val, ok := p.Mappings[cmdByte]
	if !ok {
		return 0, fmt.Errorf("Length requested for byte that wasn't present in initial Event Payloads: %v", cmdByte)
	}
	return val, nil
}

func ParseEventPayloads(stream []byte, numCmds int) EventPayloads {
	log.Debugf("Trying to parse event payloads.\nNum Cmds: %v\nStream: %v\n", numCmds, stream)
	out := EventPayloads{Mappings: make(map[byte]uint16), Payload: stream}
	if len(stream) < 1 {
		return out
	}

	offset := 0

	for range numCmds {
		command := stream[offset]
		offset++
		length := binary.BigEndian.Uint16(stream[offset : offset+2])
		offset += 2

		out.Mappings[command] = length
	}

	return out
}

// ParseNextEvent will pass the given payload off to the appropriate event parser based on given commandByte
func ParseNextEvent(commandByte byte, payload []byte) Event {

	switch commandByte {
	// case EventPayloadsByte:
	// 	return ParseEventPayloads(payload)
	case GameStartByte:
		return ParseGameStart(payload)
	// case PreFrameUpdateByte:
	// 	return ParsePreFrameUpdate(payload)
	// case PostFrameUpdateByte:
	// 	return ParsePostFrameUpdate(payload)
	// case GameEndByte:
	// 	return ParseGameEnd(payload)
	// case FrameStartByte:
	// 	return ParseFrameStart(payload)
	// case ItemUpdateByte:
	// 	return ParseItemUpdate(payload)
	// case FrameBookendByte:
	// 	return ParseFrameBookend(payload)
	// case GeckoListByte:
	// 	return ParseGeckoList(payload)
	// case FountainPlatformsByte:
	// 	return ParseFountainPlatform(payload)
	// case WhispyBlowDirByte:
	// 	return ParseWhispyBlowDir(payload)
	// case StadiumTransformByte:
	// 	return ParseStadiumTransform(payload)
	// case MessageSplitterByte:
	// 	return ParseMessageSplitter(payload)
	default:
		return nil
	}
}
