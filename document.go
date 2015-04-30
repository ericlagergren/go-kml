package kml

import "encoding/xml"

const NameSpace = "http://www.opengis.net/kml/2.2"

// Ignoring the <?xml... ?> header, this is the "root" of the .kml file
type KML struct {
	object

	XMLName  xml.Name
	Document *Document
}

// Create a new file structure with a generic object (with NameSpace) and
// and an empty document.
func NewFile() *KML {
	object := newObject()
	name := xml.Name{Local: "xml", Space: NameSpace}
	document := new(Document)
	return &KML{object, name, document}
}

// Beginning of the document; immediately follows KML.
type Document struct {
	object

	Folder *Folder `xml:"Folder"`
}

// Set the Document tag content on a KML structure.
func (k *KML) SetDocument(folder *Folder) {
	k.Document = &Document{Folder: folder}
}

type Folder struct {
	object

	XMLName   struct{}    `xml:"Folder"`
	Name      string      `xml:"name"`
	Schema    *Schema     `xml:"Schema"`
	Placemark []Placemark `xml:"Placemark"`
}

// Set the Folder tag content on a KML structure.
func (k *KML) SetFolder(name string, schema *Schema, placemark *Placemark) {
	k.Document.Folder = &Folder{
		Name:      name,
		Schema:    schema,
		Placemark: make([]Placemark, 1),
	}
}

type Schema struct {
	XMLName     struct{}      `xml:"Schema"`
	Name        string        `xml:"name,attr"`
	Id          string        `xml:"id,attr"`
	SimpleField []SimpleField `xml:"SimpleField"`
}

// Set the Schema tag on a KML structure.
func (k *KML) SetSchema(name, id string, fields []SimpleField) {
	k.Document.Folder.Schema = &Schema{
		Name: name, Id: id, SimpleField: fields,
	}
}

type SimpleField struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
}

// Append a SimpleField to the KML's schema.
func (k *KML) SetSimpleField(name, typ string) {
	k.Document.Folder.Schema.SimpleField = append(
		k.Document.Folder.Schema.SimpleField, SimpleField{name, typ},
	)
}
