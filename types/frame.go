package types

import "go.dalton.dog/slp/events"

type Frame struct {
	Start events.FrameStart

	// Each of these contains 1 object per character. 4 ICs in doubles would be 8 characters
	PreFrame  [8]events.PreFrameUpdate
	PostFrame [8]events.PostFrameUpdate

	// ItemUpdates [15]events.ItemUpdate

	Bookend events.FrameBookend
}
