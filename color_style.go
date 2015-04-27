package kml

// Styles found in section 12 of 07-147r2_OGC_KML_2.2.pdf
// http://read.pudn.com/downloads142/doc/617162/07-147r2_OGC_KML_2.2.pdf

import "encoding/xml"

type colorStyle struct {
	color color     `xml:"color"`
	mode  colorMode `xml:"colorMode"`
}

// 12.9
type Icon struct {
	href string `xml:"href"`
}

// 12.10
type LabelStyle struct {
	colorStyle
	scale float64 `xml:"scale"`
}

// High-level structure containing the styling for an
// AbstractStyleSelectorGroup.
// Section 12.2.1
type Style struct {
	XMLName    xml.Name      `xml:"Style"`
	Icon       *IconStyle    `xml:"IconStyle"`
	Label      *LabelStyle   `xml:"LabelStyle"`
	Line       *LineStyle    `xml:"LineStyle"`
	Poly       *PolyStyle    `xml:"PolyStyle"`
	Balloon    *BalloonStyle `xml:"BalloonStyle"`
	List       *ListStyle    `xml:"ListStyle"`
	SchemaData *SchemaData   `xml:"SchemaData"`
}

type IconStyle struct {
	scale   float64  `xml:"scale"`
	heading angle360 `xml:"heading"`
	icon    Icon     `xml:"Icon"`
	hotSpot vec2     `xml:"hotSpot"`
}

type LineStyle struct {
	XMLName xml.Name `xml:"LineStyle"`
	Color   string   `xml:"color"`
}

type PolyStyle struct {
	XMLName xml.Name `xml:"PolyStyle"`
	Fill    string   `xml:"fill"`
}

// 12.6
type BalloonStyle struct {
	XMLName     xml.Name        `xml:"BalloonStyle"`
	bgColor     color           `xml:"bgColor"`     // ffffffff
	textColor   color           `xml:"textColor"`   // ffffffff
	displayMode displayModeEnum `xml:"displayMode"` // default
}

// 12.13
type ListStyle struct {
	XMLName      xml.Name     `xml:"ListStyle"`
	listItemType listItemEnum `xml:"listItemType"` // check
	bgColor      color        `xml:"bgColor"`      // ffffffff
}
