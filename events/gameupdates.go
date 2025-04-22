package events

// GameStart indicates that a game is starting, and contains the information necessary for initialization.
// Will ALWAYS occur once, and ALWAYS right after the Event Payloads event
type GameStart struct {
	ExtractCodeVersion Version
	GameInfo           GameInfoBlock // Docs say this will always be [312]uint8
	RandomSeed         uint32
	DashbackFix        uint32
	ShieldDropFix      uint32
	Nametag            [8]rune
	IsPAL              bool
	IsFrozenPS         bool

	MajorSceneNum uint8 // Should be 0x2 if game is VS Mode, 0x8 when online (has rollbacks)
	MinorSceneNum uint8 // "Mostly useless at the moment, should always be 0x2"

	DisplayName string
	ConnectCode string
	SlippiUID   string

	LanguageOption uint8 // 0 = Japanese, 1 = English

	MatchID    string // ID consisting of mode and time. Max 50 chars
	GameNumber uint32 // For the given MatchID, starts at 1
	Tiebreaker uint32 // For the given GameNumber, will be 0 if NOT a tiebreaker game
}

func ParseGameStart(payload []byte) GameStart {
	outEvent := GameStart{}

	return outEvent
}

// GameInfoBlock is 312-bytes worth of information that Melee reads from to initialize the game
type GameInfoBlock struct {
	GameInfoFlags GameInfoBitFlags

	// BitFieldOne   uint8 // Behavior of timer, quantity of character UI places, game mode
	// BitFieldTwo   uint8 // Friendly fire, BTT/Demo/Classic/Adventure/HRC/All-Star
	// BitFieldThree uint8 // Generally unknown. Bit 5 = "Single-Button Mode", Bit 8 = "Pause Mode"
	// BitFieldFour  uint8 // Behavior of UI elements during game pause

	BombRain uint8 // If 0, Bomb Rain is disabled. If anything else, bombs start to rain after 20s
	IsTeams  bool  // Whether or not the current game is a Teams match

	ItemSpawnBehavior    int8 // Controls the frequency at which items spawn
	SelfDestructScoreVal int8 // Stores the number (-2, -1, or 0) to subtract from score on an SD

	Stage     uint16 // Stage ID of the current stage
	GameTimer uint32 // Number of seconds the timer should start with

	// Bitfields that govern what items are enabled to spawn
	ItemSpawnBitFlags ItemSpawnBitFlags

	DamageRatio float32

	PlayerInfo [4]PlayerInfoBlock
}

type GameInfoBitFlags struct {
}

type ItemSpawnBitFlags struct {
	// Flags 1
	// 0x01 - Metal Box
	// 0x02 - Cloaking Device
	// 0x04 - Pokeball
	// 0x08 - Unknown
	// 0x10 - Unknown
	// 0x20 - Unknown
	// 0x40 - Unknown
	// 0x80 - Unknown

	// Flags 2
	// 0x01 - Fan
	// 0x02 - Fire Flower
	// 0x04 - Super Mushroom
	// 0x08 - Poison Mushroom
	// 0x10 - Hammer
	// 0x20 - Warp Star
	// 0x40 - Screw Attack
	// 0x80 - Bunny Hood

	// Flags 3
	// 0x01 - Ray Gun
	// 0x02 - Freezie
	// 0x04 - Food
	// 0x08 - Motion Sensor Bomb
	// 0x10 - Flipper
	// 0x20 - Super Scope
	// 0x40 - Star Rod
	// 0x80 - Lip's Stick

	// Flags 4
	// 0x01 - Heart Container
	// 0x02 - Maxim Tomato
	// 0x04 - Starman
	// 0x08 - Home Run Bat
	// 0x10 - Beam Sword
	// 0x20 - Parasol
	// 0x40 - Green Shell
	// 0x80 - Red Shell

	// Flags 5
	// 0x01 - Capsule
	// 0x02 - Box
	// 0x04 - Barrel
	// 0x08 - Egg
	// 0x10 - Party Ball
	// 0x20 - Barrel Cannon
	// 0x40 - Bob-omb
	// 0x80 - Mr. Saturn
}

// PlayerInfoBlock is a set of information present in the GameInfoBlock that's relevant to each player port
type PlayerInfoBlock struct {
	ExternalCharID  uint8
	PlayerType      uint8
	StockStartCount uint8
	CostumeIndex    uint8
	TeamShade       uint8
	Handicap        uint8
	TeamID          uint8
	PlayerBitField  uint8
	CPULevel        uint8
	DamageStart     uint16
	DamageSpawn     uint16
	OffenseRatio    float32
	DefenseRatio    float32
	ModelScale      float32
}

// Version represents a version of the Slippi extraction code.
// Consists of 4 parts, though the last part is unused
type Version struct {
	Major  uint8
	Minor  uint8
	Patch  uint8
	Unused uint8
}

// GameEnd indicates that the end of the game has occurred.
// Will ALWAYS occur once, and ALWAYS as the last event of the stream
type GameEnd struct {
	GameEndMethod    uint8    // 0 = Unresolved, 3 = Resolved, 1 = TIME!, 2 = GAME!, 7 = No Contest
	LRASIndex        int8     // Index of player who LRAS'd. -1 if N/A
	PlayerPlacements [4]uint8 // 0-indexed player positions. -1 if player not in game
}

func ParseGameEnd(payload []byte) GameEnd {
	outEvent := GameEnd{}

	return outEvent
}
