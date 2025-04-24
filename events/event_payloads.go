package events

import (
	"encoding/binary"
	"fmt"
	"github.com/charmbracelet/log"
)

// Event:     Event Payloads
// Cmd Byte:  0x35
// Added:     v0.1.0
// EventPayloads will be the first event in the byte stream. It enumerates all possible
// events and their associated payload length. An event being present in here does not
// mean that it WILL be encountered, merely that it might show up.
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
