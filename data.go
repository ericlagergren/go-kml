package kml

import "encoding/xml"

type ExtendedData struct {
	XMLName    xml.Name    `xml:"ExtendedData"`
	SchemaData *SchemaData `xml:"SchemaData"`
}

type SchemaData struct {
	XMLName    xml.Name     `xml:"SchemaData"`
	SchemaURL  string       `xml:"schemaUrl,attr"`
	SimpleData []SimpleData `xml:"SimpleData"`
}

type SimpleData struct {
	XMLName xml.Name    `xml:"SimpleData"`
	Name    string      `xml:"name,attr"`
	Value   interface{} `xml:",chardata"`
}
