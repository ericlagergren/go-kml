package kml

// import "encoding/xml"

// This is an abstract element and cannot be used directly in a KML file.
// This element is extended by the <Camera> and <LookAt> elements.
type abstractView struct {
	object
}

type Camera struct {
	abstractView
	longitude angle180
	latitude  angle90
	altitude  float64
	heading   angle360
	tilt      anglePos180
	roll      angle180
}

type LookAt struct {
	abstractView
	longitude    angle180
	latitude     angle90
	altitude     float64
	heading      angle360
	tilt         anglePos180
	_range       float64 // 'range' is a reserved word
	altitudeMode altitudeModeEnum
}
