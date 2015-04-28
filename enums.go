package kml

//go:generate stringer -type=altitudeModeEnum
type altitudeModeEnum int

const (
	ClampToGround altitudeModeEnum = iota
	RelativeToGround
	Absolute
)

//go:generate stringer -type=colorModeEnum
type colorModeEnum int

const (
	Normal colorModeEnum = iota
	Random
)

//go:generate stringer -type=displayModeEnum
type displayModeEnum int

const (
	Default displayModeEnum = iota
	Hide
)

//go:generate stringer -type=gridOriginEnum
type gridOriginEnum int

const (
	LowerLeft gridOriginEnum = iota
	UpperLeft
)

//go:generate stringer -type=listItemEnum
type listItemEnum int

const (
	Check listItemEnum = iota
	RadioFolder
	CheckOffOnly
	CheckHideChildren
)

//go:generate stringer -type=refreshModeEnum
type refreshModeEnum int

const (
	OnChange refreshModeEnum = iota
	OnInterval
	OnExpire
)

//go:generate stringer -type=unitsEnum
type unitsEnum int

const (
	Fraction unitsEnum = iota
	Pixels
	InsetPixels
)

//go:generate stringer -type=viewRefreshEnum
type viewRefreshEnum int

const (
	Never viewRefreshEnum = iota
	OnRequest
	OnStop
	OnRegion
)
