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
	name       string `xml:"name"`
	visibility bool   `xml:"visibility"`
	open       int    `xml:"open"` // bool, but either 1 or 0
	// author
	// name
	// link
	// href
	address string `xml:"address"`
	// AddressDetails
	phoneNumber  string `xml:"phoneNumber"` // RFC2806
	snippet      *Snippet
	description  string
	abstractView *abstractView `xml:"AbstractView"`
}

// A short description of the feature.
// Snippet has a maxLines attribute, an integer that specifies the maximum
// number of lines to display.
type Snippet struct {
	XMLName  xml.Name `xml:"Snippet"`
	maxLines int      `xml:"maxLines,attr"`
	snippet  string   `xml:",chardata"`
}
