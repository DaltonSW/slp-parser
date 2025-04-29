package game

// Frame represents the information contained within a single frame of gameplay
type Frame struct {
	// Number of the frame. Starts at -123. Frame 0 is when timer countdown stops
	FrameNumber int

	// Scene frame counter. Starts at 0, and counts up even if the game is paused
	SceneFrameCount int

	// Random seed at the start of the frame
	RandomSeed Seed

	PortInfo []PortInfo
	Items    []ItemInfo
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
