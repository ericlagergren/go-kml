package kml

// This is an abstract base class and cannot be used directly in a KML file.
// It provides the id attribute, which allows unique identification of a KML
// element, and the targetId attribute.
// The id attribute must be assigned if the <Update> mechanism is to be used.
type object struct {
	id       string `xml:"id,attr"`
	targetId string `xml:"targetId,attr"`
}

func newObject() object {
	return object{}
}
