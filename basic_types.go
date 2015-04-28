package kml

import "strconv"

// Angles are strings because Go's `xml:",omitempty"` tags prevent the output
// of "empty" vales, (e.g. 0, nil, false, "", etc.), and some of the default
// values of some floats are 0.0. Leaving these as floats would cause some
// valid values to be suppressed.

type angle interface {
	Check() float64
}

type angle90 string // a value ≥−90 and ≤90
func (a angle90) Check() bool {
	float, err := strconv.ParseFloat(string(a), 64)
	if err != nil {
		return false
	}
	return float >= -90.00 && float <= 90.00
}

type anglePos90 string // a value ≥0 and ≤90
func (a anglePos90) Check() bool {
	float, err := strconv.ParseFloat(string(a), 64)
	if err != nil {
		return false
	}
	return float >= 0.0 && float <= 90.00
}

type angle180 string // a value ≥−180 and ≤180
func (a angle180) Check() bool {
	float, err := strconv.ParseFloat(string(a), 64)
	if err != nil {
		return false
	}
	return float >= -180.00 && float <= 180.00
}

type anglePos180 string // a value ≥0 and ≤180
func (a anglePos180) Check() bool {
	float, err := strconv.ParseFloat(string(a), 64)
	if err != nil {
		return false
	}
	return float >= 0.00 && float <= 180.00
}

type angle360 string // a value ≥−360 and ≤360
func (a angle360) Check() bool {
	float, err := strconv.ParseFloat(string(a), 64)
	if err != nil {
		return false
	}
	return float >= -360.00 && float <= 360.00
}

type anglePos360 string // a value ≥0 and ≤360
func (a anglePos360) Check() bool {
	float, err := strconv.ParseFloat(string(a), 64)
	if err != nil {
		return false
	}
	return float >= 0.00 && float <= 360.00
}

type (
	color     string        // hexBinary value: aabbggrr
	colorMode colorModeEnum // normal, random
)

type dateTime string // time.RFC3339Nano yyyy-mm-ddThh:mm:ss.ssszzzzzz

type gridOrigin gridOriginEnum

// http://www.web3d.org/specifications/kml2.2/documentation/kml22gx_vec2Type.html
// http://www.web3d.org/specifications/kml2.2/documentation/kml22gx_unitsEnumType.html#Link57
type vec2 struct {
	x, y           float64 // default 1.0
	xunits, yuints unitsEnum
}
