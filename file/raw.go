package file

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/go-version"
	"go.dalton.dog/bark"
	"go.dalton.dog/slp/events"
)

type Raw struct {
	EventPayloads events.EventPayloads

	Bytes  []byte
	Events []events.EventRaw
}

func (r Raw) String() string {
	out := "~~ Raw ~~\n"
	// out = out + fmt.Sprintf("Payloads: %v", r.EventPayloads)
	out = out + "  Events: " + strconv.Itoa(len(r.Events))
	// out = out + "  Events:\n"
	// for idx, event := range r.Events {
	// 	out = out + fmt.Sprintf("    %v: %v %v\n", idx+1, event.GetCommandByte(), event)
	// }
	return out
}

func (r Raw) AddEvent(newEvent events.EventRaw) {
	r.Events = append(r.Events, newEvent)
}

func LoadRaw(stream []byte) (*Raw, error) {
	bark.Debugf("Starting to load raw bytes. Stream length: %v", len(stream))
	raw := &Raw{Bytes: stream}

	// First byte should *always* be the EventPayloadsByte
	if stream[0] != events.EventPayloadsByte {
		return nil, bark.NewErrorf("Expected %v as first 'raw' byte but got %v", events.EventPayloadsByte, stream[0])
	}

	// Get payload size and determine number of command:size pairs present
	payloadSize := int(stream[1])
	numCmds := (payloadSize - 1) / 3

	// Read out that many bytes
	raw.EventPayloads = events.ParseEventPayloads(stream[2:1+payloadSize], numCmds)

	bark.Debugf("Event Payloads Parsed: %v", raw.EventPayloads)

	// Chunk off already processed bytes to start iterating over events
	stream = stream[1+payloadSize:]

	verStr, err := getVersion(stream[:5])
	if err != nil {
		return nil, err
	}

	version, err := version.NewVersion(verStr)
	if err != nil {
		return nil, err
	}

	for len(stream) > 0 {
		cmdByte := stream[0]
		payloadSize, err := raw.EventPayloads.GetPayloadLength(cmdByte)
		if err != nil {
			return nil, err
		}

		bark.Infof("Loaded payload size of %v for cmdByte %X", payloadSize, cmdByte)

		payload := stream[:payloadSize+1]
		stream = stream[payloadSize+1:]

		event, err := events.ParseNextEventRaw(payload, version)
		if err != nil {
			return nil, err
		}
		if event != nil {
			raw.Events = append(raw.Events, event)
		}
	}

	return raw, nil
}

func getVersion(stream []byte) (string, error) {
	if len(stream) < 5 {
		return "", bark.NewErrorf("Byte slice passed for version extraction was too short!")
	}

	if stream[0] != events.GameStartByte {
		return "", bark.NewErrorf("Expected %v as first 'raw' byte but got %v", events.GameStartByte, stream[0])
	}

	return fmt.Sprintf("%d.%d.%d", stream[1], stream[2], stream[3]), nil

}
