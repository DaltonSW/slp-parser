package id

type MoveID int

const (
	None MoveID = iota
	NonStaling

	JabOne
	JabTwo
	JabThree
	RapidJabs

	DashAttack

	SideTilt
	UpTilt
	DownTilt

	SideSmash // Includes all angles as well as follow-up attacks like YLink's
	UpSmash
	DownSmash

	NeutralAir
	ForwardAir
	BackAir
	UpAir
	DownAir

	NeutralSpecial
	SideSpecial
	UpSpecial
	DownSpecial

	KirbyHatMario
	KirbyHatFox
	KirbyHatCFalcon
	KirbyHatDK
	KirbyHatBowser
	KirbyHatLink
	KirbyHatSheik
	KirbyHatNess
	KirbyHatPeach
	KirbyHatICs
	KirbyHatPikachu
	KirbyHatSamus
	KirbyHatYoshi
	KirbyHatJigglypuff
	KirbyHatMewtwo
	KirbyHatLuigi
	KirbyHatMarth
	KirbyHatZelda
	KirbyHatYLink
	KirbyHatDoc
	KirbyHatFalco
	KirbyHatPichu
	KirbyHatGNW
	KirbyHatGanon
	KirbyHatRoy

	UnknownOne
	UnknownTwo
	UnknownThree

	GetUpAttackBack
	GetUpAttackFront

	Pummel
	ForwardThrow
	BackThrow
	UpThrow
	DownThrow

	ForwardThrowCargo
	BackThrowCargo
	UpThrowCargo
	DownThrowCargo

	LedgeGetUpAttackSlow
	LedgeGetUpAttack

	BeamSwordJab
	BeamSwordTilt
	BeamSwordSmash
	BeamSwordDash

	HomeRunBatJab
	HomeRunBatTilt
	HomeRunBatSmash
	HomeRunBatDash

	ParasolJab
	ParasolTilt
	ParasolSmash
	ParasolDash

	FanJab
	FanTilt
	FanSmash
	FanDash

	StarRodJab
	StarRodTilt
	StarRodSmash
	StarRodDash

	LipsStickJab
	LipsStickTilt
	LipsStickSmash
	LipsStickDash

	OpenParasol

	RayGunShoot
	FireFlowerShoot

	ScrewAttack

	SuperScopeRapid
	SuperScopeCharged

	Hammer
)
