package events

// GameStartRaw (0x36 -- v0.1.0)
type GameStartRaw struct {
	CmdByte       uint8     `slp-offset:"0x00" slp-ver:"0.1.0"`
	SlippiVersion [4]uint8  `slp-offset:"0x01" slp-ver:"0.1.0"`
	GameInfoBlock [312]byte `slp-offset:"0x05" slp-ver:"0.1.0"`
	RandomSeed    uint32    `slp-offset:"0x13D" slp-ver:"0.1.0"`

	// Ports [4]PortMetadataRaw

	IsPAL          bool     `slp-offset:"0x1A1" slp-ver:"1.5.0"`
	FrozenPS       bool     `slp-offset:"0x1A2" slp-ver:"2.0.0"`
	MinorScene     uint8    `slp-offset:"0x1A3" slp-ver:"3.7.0"`
	MajorScene     uint8    `slp-offset:"0x1A4" slp-ver:"3.7.0"`
	LanguageOption uint8    `slp-offset:"0x2BD" slp-ver:"3.12.0"`
	MatchID        [51]byte `slp-offset:"0x2BE" slp-ver:"3.14.0"`
	GameNumber     uint32   `slp-offset:"0x2F1" slp-ver:"3.14.0"`
	TiebreakerNum  uint32   `slp-offset:"0x2F5" slp-ver:"3.14.0"`
}

// PreFrameRaw (0x37 -- v0.1.0)
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

// PostFrameRaw (0x38 -- v0.1.0)
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

// GameEndRaw (0x39 -- v0.1.0)
type GameEndRaw struct {
	CmdByte          uint8   `slp-offset:"0x00" slp-ver:"0.1.0"`
	EndMethod        uint8   `slp-offset:"0x01" slp-ver:"0.1.0"`
	LRASInitiator    int8    `slp-offset:"0x02" slp-ver:"2.0.0"`
	PlayerPlacements [4]int8 `slp-offset:"0x03" slp-ver:"3.13.0"`
}

// FrameStartRaw (0x3A -- v2.2.0)
type FrameStartRaw struct {
	CmdByte           uint8  `slp-offset:"0x00" slp-ver:"2.2.0"`
	FrameNumber       int32  `slp-offset:"0x01" slp-ver:"2.2.0"`
	RandomSeed        uint32 `slp-offset:"0x05" slp-ver:"2.2.0"`
	SceneFrameCounter uint32 `slp-offset:"0x09" slp-ver:"3.10.0"`
}

// FrameBookendRaw (0x3C -- v3.0.0)
type FrameBookendRaw struct {
	CmdByte     uint8  `slp-offset:"0x00" slp-ver:"0.1.0"`
	FrameNumber int32  `slp-offset:"0x01" slp-ver:"0.1.0"`
	Seed        uint32 `slp-offset:"0x05" slp-ver:"0.1.0"`
}

// ItemUpdateRaw (0x3B -- v3.0.0)
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
