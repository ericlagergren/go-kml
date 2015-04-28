package kml

// import "encoding/xml"

func newAbstractView() abstractView {
	return abstractView{object: newObject()}
}

// This is an abstract element and cannot be used directly in a KML file.
// This element is extended by the <Camera> and <LookAt> elements.
type abstractView struct {
	object
}

func NewCamera() *Camera {
	return &Camera{abstractView: newAbstractView()}
}

type Camera struct {
	abstractView

	Longitude angle180    `xml:"longitude"`
	Latitude  angle90     `xml:"latitude"`
	Altitude  float64     `xml:"altitude"`
	Heading   angle360    `xml:"heading"`
	Tilt      anglePos180 `xml:"tilt"`
	Roll      angle180    `xml:"roll"`
}

func NewLookAt() *LookAt {
	return &LookAt{abstractView: newAbstractView()}
}

type LookAt struct {
	abstractView

	Longitude    angle180         `xml:"longitude"`
	Latitude     angle90          `xml:"latitude"`
	Altitude     float64          `xml:"altitude"`
	Heading      angle360         `xml:"heading"`
	Tilt         anglePos180      `xml:"tilt"`
	Range        float64          `xml:"range"`
	AltitudeMode altitudeModeEnum `xml:"altitudeMode"`
}
