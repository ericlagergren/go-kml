package kml

import "encoding/xml"

// This is an abstract element and cannot be used directly in a KML file.
// The following diagram shows how some of a Feature's elements appear
// in Google Earth.
//
// TODO atom:*
// TODO AddressDetails
type feature struct {
	object

	Name       string `xml:"name"`
	Visibility bool   `xml:"visibility"`
	Open       int    `xml:"open"` // bool, but either 1 or 0
	// author
	// name
	// link
	// href
	Address string `xml:"address"`
	// AddressDetails
	PhoneNumber  string `xml:"phoneNumber"` // RFC2806
	Snippet      *Snippet
	Description  string
	AbstractView *abstractView `xml:"AbstractView"`
}

// A short description of the feature.
// Snippet has a maxLines attribute, an integer that specifies the maximum
// number of lines to display.
type Snippet struct {
	XMLName  xml.Name `xml:"Snippet"`
	MaxLines int      `xml:"maxLines,attr"`
	snippet  string   `xml:",chardata"`
}
