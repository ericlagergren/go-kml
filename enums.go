package kml

//go:generate stringer -type=altitudeModeEnum
type altitudeModeEnum int

const (
	clampToGround altitudeModeEnum = iota
	relativeToGround
	absolute
)

//go:generate stringer -type=colorModeEnum
type colorModeEnum int

const (
	normal colorModeEnum = iota
	random
)

//go:generate stringer -type=displayModeEnum
type displayModeEnum int

const (
	Default displayModeEnum = iota // lowercase 'default' is reserved
	hide
)

//go:generate stringer -type=gridOriginEnum
type gridOriginEnum int

const (
	lowerLeft gridOriginEnum = iota
	upperLeft
)

//go:generate stringer -type=listItemEnum
type listItemEnum int

const (
	check listItemEnum = iota
	radioFolder
	checkOffOnly
	checkHideChildren
)

//go:generate stringer -type=refreshModeEnum
type refreshModeEnum int

const (
	onChange refreshModeEnum = iota
	onInterval
	onExpire
)

//go:generate stringer -type=unitsEnum
type unitsEnum int

const (
	fraction unitsEnum = iota
	pixels
	insetPixels
)

//go:generate stringer -type=viewRefreshEnum
type viewRefreshEnum int

const (
	never viewRefreshEnum = iota
	onRequest
	onStop
	onRegion
)
