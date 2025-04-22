package events

type FountainPlatform struct {
	FrameNumber int32
	Platform    uint8   // 0 = Right, 1 = Left
	Height      float32 // Platform's new height
}

func ParseFountainPlatform(payload []byte) FountainPlatform {
	outEvent := FountainPlatform{}

	return outEvent
}

type WhispyBlowDirection struct {
	FrameNumber int32
	Direction   uint8 // 0 = None, 1 = Left, 2 = Right
}

func ParseWhispyBlowDir(payload []byte) WhispyBlowDirection {
	outEvent := WhispyBlowDirection{}

	return outEvent
}

type StadiumTransformation struct {
	FrameNumber   int32
	SubEvent      uint16 // Sub-event for each transformation
	TransformType uint16 // 3 = Fire, 4 = Grass, 5 = Normal, 6 = Rock, 9 = Water
}

func ParseStadiumTransform(payload []byte) StadiumTransformation {
	outEvent := StadiumTransformation{}

	return outEvent
}

type GeckoList struct {
}

func ParseGeckoList(payload []byte) GeckoList {
	outEvent := GeckoList{}

	return outEvent
}

type MessageSplitter struct {
}

func ParseMessageSplitter(payload []byte) MessageSplitter {
	outEvent := MessageSplitter{}

	return outEvent
}
