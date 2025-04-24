package events

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
