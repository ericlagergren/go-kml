package kml

import "encoding/xml"

type Placemark struct {
	XMLName      xml.Name      `xml:"Placemark"`
	Style        *Style        `xml:"Style"`
	ExtendedData *ExtendedData `xml:"ExtendedData"`
	Polygon      *Polygon      `xml:"Polygon"`
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
