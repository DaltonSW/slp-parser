package enums

// Character Select Screen IDs
const (
	CaptainFalconCSS    = 0x00
	DonkeyKongCSS       = 0x01
	FoxCSS              = 0x02
	GameAndWatchCSS     = 0x03
	KirbyCSS            = 0x04
	BowserCSS           = 0x05
	LinkCSS             = 0x06
	LuigiCSS            = 0x07
	MarioCSS            = 0x08
	MarthCSS            = 0x09
	MewtwoCSS           = 0x0A
	NessCSS             = 0x0B
	PeachCSS            = 0x0C
	PikachuCSS          = 0x0D
	IceClimbersCSS      = 0x0E
	JigglypuffCSS       = 0x0F
	SamusCSS            = 0x10
	YoshiCSS            = 0x11
	ZeldaCSS            = 0x12
	SheikCSS            = 0x13
	FalcoCSS            = 0x14
	YoungLinkCSS        = 0x15
	DrMarioCSS          = 0x16
	RoyCSS              = 0x17
	PichuCSS            = 0x18
	GanondorfCSS        = 0x19
	MasterHandCSS       = 0x1A
	WireframeMaleCSS    = 0x1B
	WireframeFemaleCSS  = 0x1C
	GigaBowserCSS       = 0x1D
	CrazyHandCSS        = 0x1E
	SandbagCSS          = 0x1F
	PopoCSS             = 0x20
	UserSelectOrNoneCSS = 0x21
)

// Character In-Game IDs
const (
	MarioInGame           = 0x00
	FoxInGame             = 0x01
	CaptainFalconInGame   = 0x02
	DonkeyKongInGame      = 0x03
	KirbyInGame           = 0x04
	BowserInGame          = 0x05
	LinkInGame            = 0x06
	SheikInGame           = 0x07
	NessInGame            = 0x08
	PeachInGame           = 0x09
	PopoInGame            = 0x0A
	NanaInGame            = 0x0B
	PikachuInGame         = 0x0C
	SamusInGame           = 0x0D
	YoshiInGame           = 0x0E
	JigglypuffInGame      = 0x0F
	MewtwoInGame          = 0x10
	LuigiInGame           = 0x11
	MarthInGame           = 0x12
	ZeldaInGame           = 0x13
	YoungLinkInGame       = 0x14
	DrMarioInGame         = 0x15
	FalcoInGame           = 0x16
	PichuInGame           = 0x17
	MrGameAndWatchInGame  = 0x18
	GanondorfInGame       = 0x19
	RoyInGame             = 0x1A
	MasterHandInGame      = 0x1B
	WireframeMaleInGame   = 0x1D
	WireframeFemaleInGame = 0x1E
	GigaBowserInGame      = 0x1F
	SandbagInGame         = 0x20
)

// Internal Stage IDs
const (
	PrincessPeachsCastleStage = 0x02
	RainbowCruiseStage        = 0x03
	KongoJungleStage          = 0x04
	JungleJapesStage          = 0x05
	GreatBayStage             = 0x06
	HyruleTempleStage         = 0x07
	BrinstarStage             = 0x08
	BrinstarDepthsStage       = 0x09
	YoshisStoryStage          = 0x0A
	YoshisIslandStage         = 0x0B
	FountainofDreamsStage     = 0x0C
	GreenGreensStage          = 0x0D
	CorneriaStage             = 0x0E
	VenomStage                = 0x0F
	PokemonStadiumStage       = 0x10
	PokeFloatsStage           = 0x11
	MuteCityStage             = 0x12
	BigBlueStage              = 0x13
	OnettStage                = 0x14
	FoursideStage             = 0x15
	IcicleMountainStage       = 0x16
	MushroomKingdomStage      = 0x18
	MushroomKingdomIIStage    = 0x19
	FlatZoneStage             = 0x1B
	DreamLandStage            = 0x1C
	YoshisIsland64Stage       = 0x1D
	KongoJungle64Stage        = 0x1E
)

// Item IDs
const (
	CapsuleItem            = 0x00
	BoxItem                = 0x01
	BarrelItem             = 0x02
	EggItem                = 0x03
	PartyBallItem          = 0x04
	BarrelCannonItem       = 0x05
	BobombItem             = 0x06
	MrSaturnDoseiItem      = 0x07
	HeartContainerItem     = 0x08
	MaximTomatoItem        = 0x09
	StarmanItem            = 0x0A
	HomeRunBatItem         = 0x0B
	BeamSwordItem          = 0x0C
	ParasolItem            = 0x0D
	GreenShellItem         = 0x0E
	RedShellItem           = 0x0F
	RayGunLGunItem         = 0x10
	FreezieFreezeItem      = 0x11
	FoodItem               = 0x12
	ProximityMineItem      = 0x13
	FlipperItem            = 0x14
	SuperScopeItem         = 0x15
	StarRodItem            = 0x16
	LipsStickItem          = 0x17
	FanHarisenItem         = 0x18
	FireFlowerItem         = 0x19
	SuperMushroomItem      = 0x1A
	MiniMushroomItem       = 0x1B
	WarpStarItem           = 0x1D
	ScrewAttackItem        = 0x1E
	BunnyHoodtem           = 0x1F
	MetalBoxItem           = 0x20
	CloakingDeviceItem     = 0x21
	PokeBallItem           = 0x22
	RayGunrecoileffectItem = 0x23
	StarRodStarItem        = 0x24
	LipsStickDustItem      = 0x25
)

// CPU AI Types
const (
	StayAI                   = 0x00
	EscapeAI                 = 0x02
	JumpAI                   = 0x03
	NormalAI                 = 0x04
	ManualAI                 = 0x05
	NanaAI                   = 0x06
	DefensiveAI              = 0x07
	StruggleAI               = 0x08
	FreakAI                  = 0x09
	CooperateAI              = 0x0A
	BombFestEventLinkAI      = 0x0B
	BombFestEventSamusAI     = 0x0C
	OnlyItemAI               = 0x0D
	HideAndSheikEventZeldaAI = 0x0E
	NoActAI                  = 0x0F
	AirAI                    = 0x10
	ItemAI                   = 0x11
	GuardEdgeAI              = 0x12
	CooperativeAI            = 0x14
	Coop2AI                  = 0x15
	Normal2AI                = 0x16
	MultiManMeleeAI          = 0x17
	EscapeAttackAI           = 0x18
	WalkAttackAI             = 0x19
	StayAttackAI             = 0x1A
)
