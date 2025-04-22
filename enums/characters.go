package enums

// Character External IDs
const (
	CaptainFalconExt    = 0x00
	DonkeyKongExt       = 0x01
	FoxExt              = 0x02
	GameAndWatchExt     = 0x03
	KirbyExt            = 0x04
	BowserExt           = 0x05
	LinkExt             = 0x06
	LuigiExt            = 0x07
	MarioExt            = 0x08
	MarthExt            = 0x09
	MewtwoExt           = 0x0A
	NessExt             = 0x0B
	PeachExt            = 0x0C
	PikachuExt          = 0x0D
	IceClimbersExt      = 0x0E
	JigglypuffExt       = 0x0F
	SamusExt            = 0x10
	YoshiExt            = 0x11
	ZeldaExt            = 0x12
	SheikExt            = 0x13
	FalcoExt            = 0x14
	YoungLinkExt        = 0x15
	DrMarioExt          = 0x16
	RoyExt              = 0x17
	PichuExt            = 0x18
	GanondorfExt        = 0x19
	MasterHandExt       = 0x1A
	WireframeMaleExt    = 0x1B
	WireframeFemaleExt  = 0x1C
	GigaBowserExt       = 0x1D
	CrazyHandExt        = 0x1E
	SandbagExt          = 0x1F
	PopoExt             = 0x20
	UserSelectOrNoneExt = 0x21
)

// Character Internal IDs
const (
	MarioInt           = 0x00
	FoxInt             = 0x01
	CaptainFalconInt   = 0x02
	DonkeyKongInt      = 0x03
	KirbyInt           = 0x04
	BowserInt          = 0x05
	LinkInt            = 0x06
	SheikInt           = 0x07
	NessInt            = 0x08
	PeachInt           = 0x09
	PopoInt            = 0x0A
	NanaInt            = 0x0B
	PikachuInt         = 0x0C
	SamusInt           = 0x0D
	YoshiInt           = 0x0E
	JigglypuffInt      = 0x0F
	MewtwoInt          = 0x10
	LuigiInt           = 0x11
	MarthInt           = 0x12
	ZeldaInt           = 0x13
	YoungLinkInt       = 0x14
	DrMarioInt         = 0x15
	FalcoInt           = 0x16
	PichuInt           = 0x17
	MrGameAndWatchInt  = 0x18
	GanondorfInt       = 0x19
	RoyInt             = 0x1A
	MasterHandInt      = 0x1B
	WireframeMaleInt   = 0x1D
	WireframeFemaleInt = 0x1E
	GigaBowserInt      = 0x1F
	SandbagInt         = 0x20
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
