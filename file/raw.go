package file

import (
	"fmt"
	"strconv"

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
	bark.Debug(fmt.Sprintf("Starting to load raw bytes. Stream length: %v", len(stream)))
	raw := &Raw{Bytes: stream}

	offset := 0
	if stream[offset] != events.EventPayloadsByte {
		return nil, fmt.Errorf("Expected %v as first 'raw' byte but got %v", events.EventPayloadsByte, stream[offset])
	}
	offset++

	payloadSize := int(stream[offset])
	numCmds := (payloadSize - 1) / 3
	raw.EventPayloads = events.ParseEventPayloads(stream[offset+1:offset+payloadSize], numCmds)

	bark.Debug(fmt.Sprintf("Event Payloads Parsed: %v", raw.EventPayloads))

	stream = stream[offset+payloadSize:]

	for len(stream) > 0 {
		cmdByte := stream[0]
		payloadSize, err := raw.EventPayloads.GetPayloadLength(cmdByte)
		if err != nil {
			return nil, err
		}

		bark.Info(fmt.Sprintf("Loaded payload size of %v for cmdByte %X", payloadSize, cmdByte))

		payload := stream[:payloadSize+1]
		stream = stream[payloadSize+1:]

		event, err := events.ParseNextEventRaw(payload)
		if err != nil {
			return nil, err
		}
		if event != nil {
			bark.Info("Appending to raw.Events")
			raw.Events = append(raw.Events, event)
		}
	}

	// Loop until stream is empty
	//	Grab the command byte
	//	Grab the length byte and strip off that many bytes from the payload
	//	Pass those in to events.ParseNextEvent to get the appropriate event to add to the slice

	return raw, nil
}
