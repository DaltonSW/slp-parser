package game

import "go.dalton.dog/slp/id"

type PortInfo struct {
	LeaderData   CharacterFrameData
	FollowerData CharacterFrameData
}

type CharacterFrameData struct {
	PreFrameInfo  CharPreFrameInfo
	PostFrameInfo CharPostFrameInfo
}

type CharPreFrameInfo struct {
	RandomSeed    Seed
	ActionStateID id.ActionState

	JoystickPos Vector2
	CStickPos   Vector2
	FacingDir   Direction

	RawJoystickPos IntVector2
	RawCStickPos   IntVector2

	CurrentDamage float32
}

type CharPostFrameInfo struct {
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

	FacingDir float32 // -1 for left, 1 for right, 0 "in some rare cases with items"

	Trigger float32

	// The lower uint16 of ProcessedButtons is the same structure as PhysicalButtons
	ProcessedButtons uint32 // Joysticks, C-Sticks, and trigger analog presses
	PhysicalButtons  uint16 // D-Pads, ABXY, Z, Start, trigger digital presses

	PhysicalL float32
	PhysicalR float32

	Percent float32 // Current damage percent
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
