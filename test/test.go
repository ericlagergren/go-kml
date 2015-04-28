package main

import (
	"encoding/xml"
	"fmt"
	"os"

	kml "../"
)

func main() {
	fields := []kml.SimpleField{
		{"STATEFP10", "string"},
		{"COUNTYFP10", "string"},
	}

	style := &kml.Style{
		nil, nil,
		&kml.LineStyle{"ff0000ff"},
		&kml.PolyStyle{"0"},
		nil, nil, nil,
	}

	v := kml.NewFile()
	v.SetFolder("foo", nil)
	v.SetSchema("county10", "county10", fields)
	v.SetSimpleField("GEOID10", "string")
	v.SetPlacemark(style, nil, nil)

	os.Stdout.WriteString(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "    ")

	err := enc.Encode(v)
	if err != nil {
		panic(err)
	}
	fmt.Println()
}
