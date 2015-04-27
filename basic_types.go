package kml

type angle90 float64    // a value ≥−90 and ≤90
type anglePos90 float64 // a value ≥0 and ≤90

type angle180 float64    // a value ≥−180 and ≤180
type anglePos180 float64 // a value ≥0 and ≤180

type angle360 float64    // a value ≥−360 and ≤360
type anglePos360 float64 // a value ≥0 and ≤360

type color string            // hexBinary value: aabbggrr
type colorMode colorModeEnum // normal, random

type dateTime string // time.RFC3339Nano yyyy-mm-ddThh:mm:ss.ssszzzzzz

type gridOrigin gridOriginEnum

// http://www.web3d.org/specifications/kml2.2/documentation/kml22gx_vec2Type.html
// http://www.web3d.org/specifications/kml2.2/documentation/kml22gx_unitsEnumType.html#Link57
type vec2 struct {
	x, y           float64 // default 1.0
	xunits, yuints unitsEnum
}
