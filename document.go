package kml

import "encoding/xml"

// Ignoring the <?xml... ?> header, this is the "root" of the .kml file
type KML struct {
	XMLName  xml.Name `xml:"kml"`
	XMLNs    string   `xml:"xmlns,attr"`
	Document *Document
}

// Beginning of the document; immediately follows KML.
type Document struct {
	XMLName   xml.Name   `xml:"Document"`
	Folder    *Folder    `xml:"Folder"`
	Schema    *Schema    `xml:"Schema"`
	Placemark *Placemark `xml:"Placemark"`
}
type Folder struct {
	XMLName xml.Name `xml:"Folder"`
	Name    string   `xml:">name"`
}
type Schema struct {
	XMLName      xml.Name      `xml:"Schema"`
	Name         string        `xml:"name,attr"`
	Id           string        `xml:"id,attr"`
	SimpleFields []SimpleField `xml:"SimpleField"`
}
type SimpleField struct {
	XMLName xml.Name `xml:"SimpleField"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
}
