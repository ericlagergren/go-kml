package kml

import (
	"reflect"
)

var vX = reflect.Value{}

type valContainer struct {
	Vals []reflect.Value
}

// Finds the given TAG with the given ATTRibute. OFF is the offset of the
// element you want. -1 is all, 0 is the first, 1 is the second, etc.
// Will return an index error if OFF is out of the range.
func (k *KML) Find(tag, attr string, off int) []reflect.Value {

	vC := &valContainer{Vals: make([]reflect.Value, 0)}
	vC.traverse(k, tag, findTag)

	if off != -1 {
		return vC.Vals[off : off+1]
	}

	return vC.Vals
}

// Searches a struct a tag. Returns false if it's not found.
func findTag(v interface{}, tag string) (reflect.Value, bool) {
	val := reflect.ValueOf(v)

	f := val.FieldByName(tag)

	return f, f.IsValid()
}

// Traverse the KML structure, calling fn for each struct element.
func (vC *valContainer) traverse(v interface{}, tag string, fn func(interface{}, string) (reflect.Value, bool)) {

	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Struct {
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		} else if val.Kind() == reflect.Slice {
			for i := 0; i < val.Len(); i++ {
				vC.traverse(val.Slice(i, i+1), tag, fn)
			}
		} else {
			return
		}
	}

	var iface interface{}
	if val.CanInterface() {
		iface = val.Interface()
	}

	if rv, b := fn(iface, tag); b {
		vC.Vals = append(vC.Vals, rv)
	}

	n := val.NumField()
	for i := 0; i < n; i++ {

		var field interface{}

		tmp := val.Field(i)
		if !tmp.CanInterface() {
			continue
		}
		field = tmp.Interface()

		switch reflect.ValueOf(field).Kind() {
		case reflect.Ptr:
			vC.traverse(field, tag, fn)
		case reflect.Struct:
			if rv, b := fn(field, tag); b {
				vC.Vals = append(vC.Vals, rv)
			}
			vC.traverse(field, tag, fn)
		case reflect.Slice:
			sl := tmp.Slice(0, 1)
			if sl.Kind() != reflect.Struct {
				break
			}
			vC.traverse(field, tag, fn)
		}
	}
}
