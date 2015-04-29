package kml

import (
	// "encoding/xml"
	"reflect"
)

// Finds the given TAG with the given ATTRibute. OFF is the offset of the
// element you want. 0 is the first, 1 is the second, etc.
func (k *KML) Find(tag, attr string, off int) reflect.Value {
	// Quick check in case it's a Document or something
	v := reflect.ValueOf(k).Elem()
	f := v.FieldByName(tag)

	vX := reflect.Value{}
	if f != vX {
		return f
	}

}

func (k *KML) traverse(v reflect.Value, tag, attr string, off int) reflect.Value {
	head := v

	n := v.NumField()
	for i := 0; i < n; i++ {
		k := v.Field(i)
		if k.Kind() == reflect.Ptr {
			k = k.Elem()
		}

		tmp := k.FieldByName(tag)

		if tmp != vX {
			return tmp
		}
	}

	return f
}
