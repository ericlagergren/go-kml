package kml

import "encoding/xml"

type Placemark struct {
	XMLName      xml.Name      `xml:"Placemark"`
	Style        *Style        `xml:"Style"`
	ExtendedData *ExtendedData `xml:"ExtendedData"`
	Polygon      *Polygon      `xml:"Polygon"`
}

func (k *KML) SetPlacemark(style *Style, data *ExtendedData, polygon *Polygon) {

	if k.Document.Folder.Placemark == nil {
		k.Document.Folder.Placemark = &Placemark{}
	}

	k.Document.Folder.Placemark.Style = style
	k.Document.Folder.Placemark.ExtendedData = data
	k.Document.Folder.Placemark.Polygon = polygon
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
