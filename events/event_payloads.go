package events

import (
	"encoding/binary"
	"fmt"
)

// EventPayloads
// Command Byte:  0x35
// Added In:      v0.1.0
// EventPayloads will be the first event in the byte stream. It enumerates all possible
// events and their associated payload length. An event being present in here does not
// mean that it WILL be encountered, merely that it might show up.
type EventPayloads struct {
	Payload  []byte
	Mappings map[byte]uint16
}

func (p EventPayloads) GetCommandByte() byte { return EventPayloadsByte }

func (p EventPayloads) GetName() string { return "Event Payloads" }

// GetPayloadLength will return the length of the payload associated with
// the command byte passed in.
func (p EventPayloads) GetPayloadLength(cmdByte byte) (uint16, error) {
	// log.Debugf("Getting payload length for byte %X", cmdByte)
	val, ok := p.Mappings[cmdByte]
	if !ok {
		return 0, fmt.Errorf("Length requested for byte that wasn't present in initial Event Payloads: %v", cmdByte)
	}
	return val, nil
}

// ParseEventPayloads will take in the EventPayloads event (first event of the file) and
// parse out the possible commands and their associated payload lengths.
func ParseEventPayloads(stream []byte, numCmds int) EventPayloads {
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
