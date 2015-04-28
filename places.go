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
	XMLName         xml.Name         `xml:"Polygon"`
	OuterBoundaryIs *OuterBoundaryIs `xml:"outerBoundaryIs"`
	LinearRing      *LinearRing      `xml:"LinearRing"`
}

type OuterBoundaryIs struct {
	XMLName    xml.Name    `xml:"outerBoundaryIs"`
	LinearRing *LinearRing `xml:"LinearRing"`
}

type LinearRing struct {
	XMLName     xml.Name `xml:"LinearRing"`
	Coordinates string   `xml:"coordinates"`
}
