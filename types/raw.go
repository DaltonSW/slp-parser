package types

import "go.dalton.dog/slp/events"

type Raw struct {
	Bytes  []byte
	Events []events.Event
}

func (r Raw) AddEvent(newEvent events.Event) {
	r.Events = append(r.Events, newEvent)
}

func LoadRaw(stream []byte) Raw {
	raw := Raw{Bytes: stream}

	// Loop until stream is empty
	//	Grab the command byte
	//	Grab the length byte and strip off that many bytes from the payload
	//	Pass those in to events.ParseNextEvent to get the appropriate event to add to the slice

	return raw
}
