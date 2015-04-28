package kml

type Placemark struct {
	Style        *Style        `xml:"Style"`
	ExtendedData *ExtendedData `xml:"ExtendedData"`
	Polygon      *Polygon      `xml:"Polygon"`
}

func (k *KML) SetPlacemark(style *Style, data *ExtendedData, polygon *Polygon) {

	if k.Document.Folder.Placemark == nil {
		k.Document.Folder.Placemark = make([]Placemark, 1)
	}

	placemark := Placemark{style, data, polygon}

	k.Document.Folder.Placemark = append(
		k.Document.Folder.Placemark, placemark)
}

type Polygon struct {
	Extrude       string           `xml:"extrude,omitempty"`
	AltitudeMode  altitudeModeEnum `xml:"altitudeMode,omitempty"`
	OuterBoundary *OuterBoundaryIs `xml:"outerBoundaryIs"`
	InnerBoundary *InnerBoundaryIs `xml:"innerBoundaryIs"`
}

type OuterBoundaryIs struct {
	LinearRing *LinearRing `xml:"LinearRing"`
}

type InnerBoundaryIs struct {
	LinearRing *LinearRing `xml:"LinearRing"`
}

type LinearRing struct {
	Coordinates string `xml:"coordinates"`
}
