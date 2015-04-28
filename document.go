package kml

import "encoding/xml"

const NameSpace = "http://www.opengis.net/kml/2.2"

// Ignoring the <?xml... ?> header, this is the "root" of the .kml file
type KML struct {
	object

	XMLName  xml.Name
	Document *Document
}

// Create a new file structure with a generic object with NameSpace and
// and empty document.
func NewFile() *KML {
	object := newObject()
	name := xml.Name{Local: "xml", Space: NameSpace}
	document := &Document{}
	return &KML{object, name, document}
}

// Beginning of the document; immediately follows KML.
type Document struct {
	object

	XMLName struct{} `xml:"Document"`

	Folder    *Folder    `xml:"Folder"`
	Placemark *Placemark `xml:"Placemark"`
}

// Set the Document tag on a KML structure.
func (k *KML) SetDocument(folder *Folder, placemark *Placemark) {
	k.Document.Folder = folder
	k.Document.Placemark = placemark
}

type Folder struct {
	object

	XMLName struct{} `xml:"Folder"`
	Name    string   `xml:"name"`
	Schema  *Schema  `xml:"Schema"`
}

// Set the Folder tag on a KML structure.
func (k *KML) SetFolder(name string, schema *Schema) {
	k.Document.Folder = &Folder{Name: name, Schema: schema}
}

type Schema struct {
	XMLName      struct{}      `xml:"Schema"`
	Name         string        `xml:"name,attr"`
	Id           string        `xml:"id,attr"`
	SimpleFields []SimpleField `xml:"SimpleField"`
}

// Set the Schema tag on a KML structure.
func (k *KML) SetSchema(name, id string, fields []SimpleField) {
	k.Document.Folder.Schema = &Schema{
		Name: name, Id: id, SimpleFields: fields,
	}
}

type SimpleField struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
}

// Append a SimpleField to the KML's schema.
func (k *KML) SetSimpleField(name, typ string) {
	k.Document.Folder.Schema.SimpleFields = append(
		k.Document.Folder.Schema.SimpleFields, SimpleField{name, typ},
	)
}
