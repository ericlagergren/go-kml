package kml

type ExtendedData struct {
	SchemaData *SchemaData `xml:"SchemaData"`
}

type SchemaData struct {
	SchemaURL  string       `xml:"schemaUrl,attr"`
	SimpleData []SimpleData `xml:"SimpleData"`
}

type SimpleData struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}
