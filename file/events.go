package file

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"

	"github.com/hashicorp/go-version"
)

// Command Bytes
// First byte of each event, dictates what kind of event it is.
const (
	EventPayloadsByte     = 0x35
	GameStartByte         = 0x36
	PreFrameUpdateByte    = 0x37
	PostFrameUpdateByte   = 0x38
	GameEndByte           = 0x39
	FrameStartByte        = 0x3A
	ItemUpdateByte        = 0x3B
	FrameBookendByte      = 0x3C
	GeckoListByte         = 0x3D
	FountainPlatformsByte = 0x3F
	WhispyBlowDirByte     = 0x40
	StadiumTransformByte  = 0x41
	MessageSplitterByte   = 0x10
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

// EventRaw represents a raw-parsed event from an slp file.
// Each struct is 1:1 mapped to an event on the Slippi spec.
// Spec here: https://github.com/project-slippi/slippi-wiki/blob/master/SPEC.md#game-start
type EventRaw interface {
	GetCommandByte() byte
	GetEventName() string
}

// ParseNextEvent will pass the given payload off to the appropriate event parser based on given commandByte
// func ParseNextEventRaw(commandByte byte, payload []byte) (*EventRaw, error) {
func ParseNextEventRaw(payload []byte, version *version.Version) (EventRaw, error) {
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
		outEvent = &StadiumTransformRaw{}
	case MessageSplitterByte:
		outEvent = &MessageSplitRaw{}
	default:
		return nil, fmt.Errorf("Tried to parse event with unsupported cmdByte: %b", commandByte)
	}

	err := UnpackRawEvent(outEvent, payload, version)

	if err != nil {
		return nil, err
	}

	return outEvent, nil
}

// UnpackRawEvent will take in a pointer to a RawEvent struct and a payload to unpack into it.
// The parsing is facilitated by reflection based on tags placed on the struct's properties.
func UnpackRawEvent(raw EventRaw, payload []byte, fileVersion *version.Version) error {

	// Verify that the payload's first byte matches the given event's Command Byte
	if payload[0] != raw.GetCommandByte() {
		return errors.New("Mismatched command byte")
	}

	// Verify that "raw" is a pointer to a struct
	rawVal := reflect.ValueOf(raw)
	if rawVal.Kind() != reflect.Pointer || rawVal.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("Expected pointer to EventRaw implementing struct, got %T instead", raw)
	}

	structVal := rawVal.Elem()
	structType := structVal.Type()

	for i := range structType.NumField() {
		field := structType.Field(i)

		fieldVerStr := field.Tag.Get("slp-ver")
		if fieldVerStr == "" {
			return fmt.Errorf("Found field without slp-ver tag: %v", field)
		}

		// Skip trying to load any properties added in a later version of the slp file generator
		fieldVer, err := version.NewVersion(fieldVerStr)
		if err != nil || fileVersion.LessThan(fieldVer) {
			continue
		}

		// Grab the byte offset from the property tag. Every EventRaw should exclusively have tag'd properties
		offsetHex := field.Tag.Get("slp-offset")
		if offsetHex == "" {
			return fmt.Errorf("Found field without slp-offset tag: %v", field)
		}

		// Parse the offset from hex to binary
		var offset int
		_, err = fmt.Sscanf(offsetHex, "0x%x", &offset)
		if err != nil {
			return fmt.Errorf("Invalid offset tag attempted to parse")
		}

		// Grab the size in bytes of the property to unpack into
		value := structVal.Field(i)
		size := binary.Size(value.Interface())
		if size < 0 {
			return fmt.Errorf("Unsupported size for field %s", field.Name)
		}

		// Verify that we won't exceed the payload bounds when trying to access it
		if offset+size > len(payload) {
			return fmt.Errorf("Payload too short for field %s with offset 0x%x", field.Name, offset)
		}

		// Slice off the bytes for this property and create a byte Reader for it
		fieldBytes := payload[offset : offset+size]
		reader := bytes.NewReader(fieldBytes)

		// Read the value at the offset into the property
		if err := binary.Read(reader, binary.BigEndian, value.Addr().Interface()); err != nil {
			return fmt.Errorf("Binary read failed for field %s: %v", field.Name, err)
		}

	}

	return nil
}

// GameStartRaw
// Command Byte: 0x36
// Added In: v0.1.0
type GameStartRaw struct {
	CmdByte       uint8     `slp-offset:"0x00" slp-ver:"0.1.0"`
	SlippiVersion [4]uint8  `slp-offset:"0x01" slp-ver:"0.1.0"`
	GameInfoBlock [312]byte `slp-offset:"0x05" slp-ver:"0.1.0"`
	RandomSeed    uint32    `slp-offset:"0x13D" slp-ver:"0.1.0"`

	// UCF Stuff
	PortOneDashback     uint32 `slp-offset:"0x141" slp-ver:"1.0.0"`
	PortOneShieldDrop   uint32 `slp-offset:"0x145" slp-ver:"1.0.0"`
	PortTwoDashback     uint32 `slp-offset:"0x149" slp-ver:"1.0.0"`
	PortTwoShieldDrop   uint32 `slp-offset:"0x14D" slp-ver:"1.0.0"`
	PortThreeDashback   uint32 `slp-offset:"0x151" slp-ver:"1.0.0"`
	PortThreeShieldDrop uint32 `slp-offset:"0x155" slp-ver:"1.0.0"`
	PortFourDashback    uint32 `slp-offset:"0x159" slp-ver:"1.0.0"`
	PortFourShieldDrop  uint32 `slp-offset:"0x15D" slp-ver:"1.0.0"`

	// In-Game Nametag
	PortOneNametag   [16]byte `slp-offset:"0x161" slp-ver:"1.3.0"`
	PortTwoNametag   [16]byte `slp-offset:"0x171" slp-ver:"1.3.0"`
	PortThreeNametag [16]byte `slp-offset:"0x181" slp-ver:"1.3.0"`
	PortFourNametag  [16]byte `slp-offset:"0x191" slp-ver:"1.3.0"`

	// Game Info
	IsPAL      bool  `slp-offset:"0x1A1" slp-ver:"1.5.0"`
	FrozenPS   bool  `slp-offset:"0x1A2" slp-ver:"2.0.0"`
	MinorScene uint8 `slp-offset:"0x1A3" slp-ver:"3.7.0"`
	MajorScene uint8 `slp-offset:"0x1A4" slp-ver:"3.7.0"`

	// Display Name
	//	(Shift JIS string @ 0x1A5+0x1F*i)
	//	"Max 15 chars + null terminator"
	PortOneDisplayName   [16]byte `slp-offset:"0x1A5" slp-ver:"3.9.0"`
	PortTwoDisplayName   [16]byte `slp-offset:"0x1C4" slp-ver:"3.9.0"`
	PortThreeDisplayName [16]byte `slp-offset:"0x1E3" slp-ver:"3.9.0"`
	PortFourDisplayName  [16]byte `slp-offset:"0x202" slp-ver:"3.9.0"`

	// Connect Code
	//	(Shift JIS string @ 0x221+0xA*i)
	//	"Max 7 1-byte chars + 2-byte '#' + null terminator"
	PortOneConnectCode   [10]byte `slp-offset:"0x221" slp-ver:"3.9.0"`
	PortTwoConnectCode   [10]byte `slp-offset:"0x22B" slp-ver:"3.9.0"`
	PortThreeConnectCode [10]byte `slp-offset:"0x235" slp-ver:"3.9.0"`
	PortFourConnectCode  [10]byte `slp-offset:"0x23F" slp-ver:"3.9.0"`

	// Slippi UID
	//	(string @ 0x249+0x1D*i)
	//	"Max 28 chars + null terminator"
	PortOneSlippiUID   [29]byte `slp-offset:"0x249" slp-ver:"3.11.0"`
	PortTwoSlippiUID   [29]byte `slp-offset:"0x266" slp-ver:"3.11.0"`
	PortThreeSlippiUID [29]byte `slp-offset:"0x283" slp-ver:"3.11.0"`
	PortFourSlippiUID  [29]byte `slp-offset:"0x2A0" slp-ver:"3.11.0"`

	LanguageOption uint8    `slp-offset:"0x2BD" slp-ver:"3.12.0"`
	MatchID        [51]byte `slp-offset:"0x2BE" slp-ver:"3.14.0"`
	GameNumber     uint32   `slp-offset:"0x2F1" slp-ver:"3.14.0"`
	TiebreakerNum  uint32   `slp-offset:"0x2F5" slp-ver:"3.14.0"`
}

func (GameStartRaw) GetCommandByte() byte { return GameStartByte }

func (GameStartRaw) GetEventName() string { return "Game Start" }

// PreFrameRaw contains information about the player's input information pre-processing
// Command Byte: 0x37
// Added In: v0.1.0
type PreFrameRaw struct {
	CmdByte          uint8   `slp-offset:"0x00" slp-ver:"0.1.0"`
	Frame            int32   `slp-offset:"0x01" slp-ver:"0.1.0"`
	PlayerIndex      uint8   `slp-offset:"0x05" slp-ver:"0.1.0"`
	IsFollower       bool    `slp-offset:"0x06" slp-ver:"0.1.0"`
	RandomSeed       uint32  `slp-offset:"0x07" slp-ver:"0.1.0"`
	ActionStateID    uint16  `slp-offset:"0x0B" slp-ver:"0.1.0"`
	XPosition        float32 `slp-offset:"0x0D" slp-ver:"0.1.0"`
	YPosition        float32 `slp-offset:"0x11" slp-ver:"0.1.0"`
	FacingDirection  float32 `slp-offset:"0x15" slp-ver:"0.1.0"`
	JoystickX        float32 `slp-offset:"0x19" slp-ver:"0.1.0"`
	JoystickY        float32 `slp-offset:"0x1D" slp-ver:"0.1.0"`
	CStickX          float32 `slp-offset:"0x21" slp-ver:"0.1.0"`
	CStickY          float32 `slp-offset:"0x25" slp-ver:"0.1.0"`
	Trigger          float32 `slp-offset:"0x29" slp-ver:"0.1.0"`
	ProcessedButtons uint32  `slp-offset:"0x2D" slp-ver:"0.1.0"`
	PhysicalButtons  uint16  `slp-offset:"0x31" slp-ver:"0.1.0"`
	PhysicalLTrigger float32 `slp-offset:"0x33" slp-ver:"0.1.0"`
	PhysicalRTrigger float32 `slp-offset:"0x37" slp-ver:"0.1.0"`
	RawAnalogX       int8    `slp-offset:"0x3B" slp-ver:"1.2.0"`
	Percent          float32 `slp-offset:"0x3C" slp-ver:"1.4.0"`
	RawAnalogY       int8    `slp-offset:"0x40" slp-ver:"3.15.0"`
	RawCStickX       int8    `slp-offset:"0x41" slp-ver:"3.17.0"`
	RawCStickY       int8    `slp-offset:"0x42" slp-ver:"3.17.0"`
}

func (PreFrameRaw) GetCommandByte() byte { return PreFrameUpdateByte }

func (PreFrameRaw) GetEventName() string { return "Pre-Frame" }

// PostFrameRaw
// Command Byte: 0x38
// Added In: v0.1.0
type PostFrameRaw struct {
	CmdByte                 uint8   `slp-offset:"0x00" slp-ver:"0.1.0"`
	Frame                   int32   `slp-offset:"0x01" slp-ver:"0.1.0"`
	PlayerIndex             uint8   `slp-offset:"0x05" slp-ver:"0.1.0"`
	IsFollower              bool    `slp-offset:"0x06" slp-ver:"0.1.0"`
	InternalCharacterID     uint8   `slp-offset:"0x07" slp-ver:"0.1.0"`
	ActionStateID           uint16  `slp-offset:"0x08" slp-ver:"0.1.0"`
	XPosition               float32 `slp-offset:"0x0A" slp-ver:"0.1.0"`
	YPosition               float32 `slp-offset:"0x0E" slp-ver:"0.1.0"`
	FacingDirection         float32 `slp-offset:"0x12" slp-ver:"0.1.0"`
	Percent                 float32 `slp-offset:"0x16" slp-ver:"0.1.0"`
	ShieldSize              float32 `slp-offset:"0x1A" slp-ver:"0.1.0"`
	LastHittingAttackID     uint8   `slp-offset:"0x1E" slp-ver:"0.1.0"`
	CurrentComboCount       uint8   `slp-offset:"0x1F" slp-ver:"0.1.0"`
	LastHitBy               uint8   `slp-offset:"0x20" slp-ver:"0.1.0"`
	StocksRemaining         uint8   `slp-offset:"0x21" slp-ver:"0.1.0"`
	ActionStateFrameCounter float32 `slp-offset:"0x22" slp-ver:"0.2.0"`
	StateBitFlags1          uint8   `slp-offset:"0x26" slp-ver:"2.0.0"`
	StateBitFlags2          uint8   `slp-offset:"0x27" slp-ver:"2.0.0"`
	StateBitFlags3          uint8   `slp-offset:"0x28" slp-ver:"2.0.0"`
	StateBitFlags4          uint8   `slp-offset:"0x29" slp-ver:"2.0.0"`
	StateBitFlags5          uint8   `slp-offset:"0x2A" slp-ver:"2.0.0"`
	MiscAS                  float32 `slp-offset:"0x2B" slp-ver:"2.0.0"`
	IsAirborne              bool    `slp-offset:"0x2F" slp-ver:"2.0.0"`
	LastGroundID            uint16  `slp-offset:"0x30" slp-ver:"2.0.0"`
	JumpsRemaining          uint8   `slp-offset:"0x32" slp-ver:"2.0.0"`
	LCancelStatus           uint8   `slp-offset:"0x33" slp-ver:"2.0.0"`
	HurtboxCollisionState   uint8   `slp-offset:"0x34" slp-ver:"2.1.0"`
	SelfInducedXSpeed       float32 `slp-offset:"0x35" slp-ver:"3.5.0"`
	SelfInducedYSpeed       float32 `slp-offset:"0x39" slp-ver:"3.5.0"`
	AttackBasedXSpeed       float32 `slp-offset:"0x3D" slp-ver:"3.5.0"`
	AttackBasedYSpeed       float32 `slp-offset:"0x41" slp-ver:"3.5.0"`
	SelfInducedGroundXSpeed float32 `slp-offset:"0x45" slp-ver:"3.5.0"`
	HitlagRemaining         float32 `slp-offset:"0x49" slp-ver:"3.8.0"`
	AnimationIndex          uint32  `slp-offset:"0x4D" slp-ver:"3.11.0"`
	InstanceHitBy           uint16  `slp-offset:"0x51" slp-ver:"3.16.0"`
	InstanceID              uint16  `slp-offset:"0x53" slp-ver:"3.16.0"`
}

func (PostFrameRaw) GetCommandByte() byte { return PostFrameUpdateByte }

func (PostFrameRaw) GetEventName() string { return "Post-Frame" }

// GameEndRaw
// Command Byte: 0x39
// Added In: v0.1.0
type GameEndRaw struct {
	CmdByte          uint8   `slp-offset:"0x00" slp-ver:"0.1.0"`
	EndMethod        uint8   `slp-offset:"0x01" slp-ver:"0.1.0"`
	LRASInitiator    int8    `slp-offset:"0x02" slp-ver:"2.0.0"`
	PlayerPlacements [4]int8 `slp-offset:"0x03" slp-ver:"3.13.0"`
}

func (GameEndRaw) GetCommandByte() byte { return GameEndByte }

func (GameEndRaw) GetEventName() string { return "Game End" }

// FrameStartRaw
// Command Byte: 0x3A
// Added In: v2.2.0
type FrameStartRaw struct {
	CmdByte           uint8  `slp-offset:"0x00" slp-ver:"2.2.0"`
	FrameNumber       int32  `slp-offset:"0x01" slp-ver:"2.2.0"`
	RandomSeed        uint32 `slp-offset:"0x05" slp-ver:"2.2.0"`
	SceneFrameCounter uint32 `slp-offset:"0x09" slp-ver:"3.10.0"`
}

func (FrameStartRaw) GetCommandByte() byte { return FrameStartByte }

func (FrameStartRaw) GetEventName() string { return "Frame Start" }

// ItemUpdateRaw
// Command Byte: 0x3B
// Added In: v3.0.0
type ItemUpdateRaw struct {
	CmdByte         uint8   `slp-offset:"0x00" slp-ver:"3.0.0"`
	Frame           int32   `slp-offset:"0x01" slp-ver:"3.0.0"`
	TypeID          uint16  `slp-offset:"0x05" slp-ver:"3.0.0"`
	State           uint8   `slp-offset:"0x07" slp-ver:"3.0.0"`
	FacingDirection float32 `slp-offset:"0x08" slp-ver:"3.0.0"`
	XVelocity       float32 `slp-offset:"0x0C" slp-ver:"3.0.0"`
	YVelocity       float32 `slp-offset:"0x10" slp-ver:"3.0.0"`
	XPosition       float32 `slp-offset:"0x14" slp-ver:"3.0.0"`
	YPosition       float32 `slp-offset:"0x18" slp-ver:"3.0.0"`
	DamageTaken     uint16  `slp-offset:"0x1C" slp-ver:"3.0.0"`
	ExpirationTimer float32 `slp-offset:"0x1E" slp-ver:"3.0.0"`
	SpawnID         uint32  `slp-offset:"0x22" slp-ver:"3.0.0"`
	Misc1           uint8   `slp-offset:"0x26" slp-ver:"3.2.0"`
	Misc2           uint8   `slp-offset:"0x27" slp-ver:"3.2.0"`
	Misc3           uint8   `slp-offset:"0x28" slp-ver:"3.2.0"`
	Misc4           uint8   `slp-offset:"0x29" slp-ver:"3.2.0"`
	Owner           int8    `slp-offset:"0x2A" slp-ver:"3.6.0"`
	InstanceID      uint16  `slp-offset:"0x2B" slp-ver:"3.16.0"`
}

func (ItemUpdateRaw) GetCommandByte() byte { return ItemUpdateByte }

func (ItemUpdateRaw) GetEventName() string { return "Item Update" }

// FrameBookendRaw
// Command Byte: 0x3C
// Added In: v3.0.0
type FrameBookendRaw struct {
	CmdByte     uint8  `slp-offset:"0x00" slp-ver:"0.1.0"`
	FrameNumber int32  `slp-offset:"0x01" slp-ver:"0.1.0"`
	Seed        uint32 `slp-offset:"0x05" slp-ver:"0.1.0"`
}

func (FrameBookendRaw) GetCommandByte() byte { return FrameBookendByte }

func (FrameBookendRaw) GetEventName() string { return "Frame Bookend" }

// GeckoListRaw (Currently Unimplemented)
// Command Byte: 0x3D
// Added In: v3.3.0
type GeckoListRaw struct{}

func (GeckoListRaw) GetCommandByte() byte { return GeckoListByte }

func (GeckoListRaw) GetEventName() string { return "Gecko List" }

// MessageSplitRaw
// Command Byte: 0x10
// Added In: v3.3.0
type MessageSplitRaw struct{}

func (MessageSplitRaw) GetCommandByte() byte { return MessageSplitterByte }

func (MessageSplitRaw) GetEventName() string { return "Message Split" }

// FountainPlatformRaw
// Command Byte: 0x3F
// Added In: v3.18.0
type FountainPlatformRaw struct {
	CmdByte     uint8   `slp-offset:"0x00" slp-ver:"3.18.0"`
	FrameNumber int32   `slp-offset:"0x01" slp-ver:"3.18.0"`
	Platform    uint8   `slp-offset:"0x05" slp-ver:"3.18.0"`
	Height      float32 `slp-offset:"0x06" slp-ver:"3.18.0"`
}

func (FountainPlatformRaw) GetCommandByte() byte { return FountainPlatformsByte }

func (FountainPlatformRaw) GetEventName() string { return "Fountain Platform" }

// WhispyBlowDirectionRaw
// Command Byte: 0x40
// Added In: v3.18.0
type WhispyBlowDirectionRaw struct {
	CmdByte     uint8 `slp-offset:"0x00" slp-ver:"3.18.0"`
	FrameNumber int32 `slp-offset:"0x01" slp-ver:"3.18.0"`
	Direction   uint8 `slp-offset:"0x05" slp-ver:"3.18.0"`
}

func (WhispyBlowDirectionRaw) GetCommandByte() byte { return WhispyBlowDirByte }

func (WhispyBlowDirectionRaw) GetEventName() string { return "Whispy Blow Direction" }

// StadiumTransformRaw
// Command Byte: 0x41
// Added In: v3.18.0
type StadiumTransformRaw struct {
	CmdByte        uint8  `slp-offset:"0x00" slp-ver:"3.18.0"`
	FrameNumber    int32  `slp-offset:"0x01" slp-ver:"3.18.0"`
	TransformEvent uint16 `slp-offset:"0x05" slp-ver:"3.18.0"`
	TransformType  uint16 `slp-offset:"0x07" slp-ver:"3.18.0"`
}

func (StadiumTransformRaw) GetCommandByte() byte { return StadiumTransformByte }

func (StadiumTransformRaw) GetEventName() string { return "Stadium Transform" }
