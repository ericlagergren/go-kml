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

	simpleData := []kml.SimpleData{
		{"STATEFP10", "53"},
		{"COUNTYFP10", "001"},
	}

	data := &kml.ExtendedData{&kml.SchemaData{"#county10", simpleData}}

	coords := "-118.98254805808614,46.911347333288994 -118.98251205658734,46.914382331244724 -118.98240605208865,46.923506325099368 -118.98237105059073,46.926547323051238 -118.98236905054021,46.926646322984489 -118.98236205023512,46.927266322566929"

	polygon := new(kml.Polygon)
	polygon.OuterBoundary = &kml.OuterBoundaryIs{
		LinearRing: &kml.LinearRing{coords},
	}

	v := kml.NewFile()
	v.SetFolder("foo", nil)
	v.SetSchema("county10", "county10", fields)
	v.SetSimpleField("GEOID10", "string")
	v.SetPlacemark(nil, data, polygon)

	os.Stdout.WriteString(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "    ")

	err := enc.Encode(v)
	if err != nil {
		panic(err)
	}
	fmt.Println()
}
