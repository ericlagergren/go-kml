package kml

import "encoding/xml"

type Placemark struct {
	XMLName      xml.Name      `xml:"Placemark"`
	Style        *Style        `xml:"Style"`
	ExtendedData *ExtendedData `xml:"ExtendedData"`
	Polygon      *Polygon      `xml:"Polygon"`
}

func (k *KML) SetPlacemark(style *Style, data *ExtendedData, polygon *Polygon) {

	if k.Document.Placemark == nil {
		k.Document.Placemark = &Placemark{}
	}

	k.Document.Placemark.Style = style
	k.Document.Placemark.ExtendedData = data
	k.Document.Placemark.Polygon = polygon
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
