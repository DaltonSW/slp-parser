package events

// FrameStart transfers the RNG seed at the start of a frame's processing.
// This prevents desyncs "such as the knockback spiral animation desync".
type FrameStart struct {
	FrameNumber     int32  // Number of the frame. Starts at -123. Frame 0 is when timer countdown stops
	RandomSeed      uint32 // Random seed at the start of the frame
	SceneFrameCount uint32 // Scene frame counter. Starts at 0, and counts up even if the game is paused
}

// PreFrameUpdate will occur exactly once per frame per character (ICs are 2 characters).
// Contains the information required to reconstruct a replay. Information is collected
// right BEFORE controller input processing takes place.
type PreFrameUpdate struct {
	FrameNumber   int32  // Number of the frame. Starts at -123. Frame 0 is when timer countdown stops
	PlayerIndex   uint8  // Between 0 and 3. Port is index + 1
	IsFollower    bool   // true for Nana, false otherwise
	RandomSeed    uint32 // Random seed at this point
	ActionStateID uint16 // Indicates the action state the character is in

	// Player X and Y position
	PosX float32
	PosY float32

	FacingDir float32 // -1 for left, 1 for right, 0 "in some rare cases with items"

	// Joystick and C-Stick coordinates
	JoyX    float32
	JoyY    float32
	CStickX float32
	CStickY float32

	Trigger float32

	// The lower uint16 of ProcessedButtons is the same structure as PhysicalButtons
	ProcessedButtons uint32 // Joysticks, C-Sticks, and trigger analog presses
	PhysicalButtons  uint16 // D-Pads, ABXY, Z, Start, trigger digital presses

	PhysicalL float32
	PhysicalR float32

	// Joystick and C-Stick coordinates (After UCF is applied)
	JoyUCFX    int8
	JoyUCFY    int8
	CStickUCFX int8
	CStickUCFY int8

	Percent float32 // Current damage percent
}

func ParsePreFrameUpdate(payload []byte) PreFrameUpdate {
	outEvent := PreFrameUpdate{}

	return outEvent
}

// PostFrameUpdate will occur exactly once per frame per character (ICs are 2 characters).
// Contains the information for making decisions about game states. Information is collected
// at the end of game's collision detection.
type PostFrameUpdate struct {
	FrameNumber    int32 // Number of the frame. Starts at -123. Frame 0 is when timer countdown stops
	PlayerIndex    uint8
	IsFollower     bool
	InternalCharID uint8
	ActionStateID  uint16

	PosX float32
	PosY float32

	FacingDir  float32
	Percent    float32
	ShieldSize float32

	LastHittingAttackID uint8
	CurrentComboCount   uint8 // Combos as defined by the game

	LastHitBy       uint8 // Index of the player that last hit this character
	StocksRemaining uint8

	ActionStateFrameCounter float32 // Number of frames action state has been active

	StateBitFlags PostFrameStateBitFlags

	HitstunRemaining float32 // Can be used for different stuff, but this is remaining hitstun when in hitstun

	GroundAirState bool   // 0 = grounded, 1 = aerial
	LastGroundID   uint16 // ID of the last piece of ground the character stood on

	JumpsRemaining uint8
	LCancelStatus  uint8 // 0 = None, 1 = Successful, 2 = Unsuccessful

	HurtboxCollisionState uint8 // 0 = Vulnerable, 1 = Invulnerable, 2 = Intangible

	SelfInducedAirSpeedX float32 // Negative is left, positive is right
	SelfInducedSpeedY    float32

	AttackBasedSpeedX float32
	AttackBasedSpeedY float32

	SelfInducedGroundSpeedX float32

	HitlagFramesRemaining float32
	AnimationIndex        uint32
	InstanceHitBy         uint16
	InstanceID            uint16 // Serially generated, unique ID for each new action state across all fighters
}

func ParsePostFrameUpdate(payload []byte) PostFrameUpdate {
	outEvent := PostFrameUpdate{}

	return outEvent
}

// The boolean representations of the various state flags used in PostFrame Update
type PostFrameStateBitFlags struct {
	// Flags 1
	// 0x01 - Unknown
	// 0x02 - Is Absorber Active (G&W Bucket)
	// 0x04 - Unknown
	// 0x08 - Active when reflector doesn't change projectile ownership (Mewtwo Side-B)
	// 0x10 - Is Reflector Active
	// 0x20 - Unknown
	// 0x40 - Unknown
	// 0x80 - Unknown

	// Flags 2
	// 0x01 - Unknown
	// 0x02 - Unknown
	// 0x04 - Has temp intangibility or invincibility from subaction
	// 0x08 - Is Fast-Falling
	// 0x10 - Is defender in hitlag (doesn't count shield hitlag)
	// 0x20 - Is in hitlag
	// 0x40 - Unknown
	// 0x80 - Unknown

	// Flags 3
	// 0x01 - Unknown
	// 0x02 - Unknown
	// 0x04 - Is holding other character due to a grab
	// 0x08 - Unknown
	// 0x10 - Unknown
	// 0x20 - Unknown
	// 0x40 - Unknown
	// 0x80 - Is shield active

	// Flags 4
	// 0x01 - Unknown
	// 0x02 - Is in hitstun
	// 0x04 - Owner's detection hitbox touching the shield bubble
	// 0x08 - Unknown
	// 0x10 - Unknown
	// 0x20 - Powershield active
	// 0x40 - Unknown
	// 0x80 - Unknown

	// Flags 5
	// 0x01 - Unknown
	// 0x02 - Is cloaking device
	// 0x04 - Unknown
	// 0x08 - Is follower (Nana)
	// 0x10 - Is inactive (Zelda/Sheik when not in use. Should always be 0 during replays)
	// 0x20 - Unknown
	// 0x40 - Is dead
	// 0x80 - Is offscreen
}

// FrameBookend is send to determine that the entire frame's worth of data has been transferred.
// ALWAYS sent at the very end of the frame's transfer.
type FrameBookend struct {
	FrameNumber int32 // Number of the frame. Starts at -123. Frame 0 is when timer countdown stops
	LatestFrame int32 // Non-rollback should always have this equal FrameNumber
}

func ParseFrameBookend(payload []byte) FrameBookend {
	outEvent := FrameBookend{}

	return outEvent
}

// ItemUpdate is sent pertaining to an item. Up to 15 items can have data extracted per frame.
type ItemUpdate struct {
	FrameNumber int32   // Number of the frame. Starts at -123. Frame 0 is when timer countdown stops
	TypeID      uint16  // Type of item
	StateID     uint8   // State the item is in. "Mostly undocumented"
	FacingDir   float32 // -1 = left, 1 = right, "0 in some cases"

	// Velocity and Position trackers
	VelX float32
	VelY float32
	PosX float32
	PosY float32

	DamageTaken     uint16
	ExpirationTimer float32 // Number of frames remaining before item expires

	SpawnID uint32 // Incremented as new items spawn (max of 15)

	SamusMissleType uint8 // 0 = Homing, 1 = Super

	// 0 = Smile,        1 = T-Eyes, 2 = Line Eyes, 3 = Circle Eyes,
	// 4 = Upward Curve, 5 = Wink,   6 = Dot Eyes,  7 = Stitch Face
	PeachTurnipFace uint8

	SamusMewTwoIsLaunched    uint8 // 0 = false, 1 = true
	SamusMewTwoCurrentCharge uint8 // Current charge power for their charge shots

	PlayerOwner int8   // 0-3 for the player that owns the item. -1 if no owner
	InstanceID  uint16 // "Inherited instance ID of the owner. 0 when not owned"
}

func ParseItemUpdate(payload []byte) ItemUpdate {
	outEvent := ItemUpdate{}

	return outEvent
}
