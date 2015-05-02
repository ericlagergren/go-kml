package kml

import "reflect"

type FieldVal struct {
	Vals []reflect.StructField
}

// Searches the *KML structure for the given tag. E.g., Document, Name, etc.
// Unexported tags cannot be found.
func (k *KML) Find(tag, attr string) *FieldVal {

	vals := &FieldVal{Vals: make([]reflect.StructField, 0)}
	vals.traverse(k, tag, attr, findTag)

	return vals
}

// Searches a struct a tag. Returns false if it's not found.
func findTag(v *interface{}, tag, attr string) (reflect.StructField, bool) {
	typ := reflect.TypeOf(*v)

	field, ok := typ.FieldByName(tag)
	if !ok {
		return field, false
	}

	return field, true
}

// Traverse the KML structure, calling fn for each struct element.
func (fv *FieldVal) traverse(v interface{}, tag, attr string, fn func(*interface{}, string, string) (reflect.StructField, bool)) {

	val := reflect.ValueOf(v)

	switch val.Kind() {
	case reflect.Struct:
	case reflect.Ptr:
		// Dereference pointer
		val = val.Elem()
	case reflect.Slice:
		// Check each item in the slice because it may be a struct!
		for i := 0; i < val.Len(); i++ {
			fv.traverse(val.Slice(i, i+1), tag, attr, fn)
		}
	default:
		// We only want the above 3 types.
		return
	}

	var iface interface{}
	if val.CanInterface() {
		iface = val.Interface()
	}

	if rv, b := fn(&iface, tag, attr); b {
		fv.Vals = append(fv.Vals, rv)
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
			fv.traverse(field, tag, attr, fn)
		case reflect.Struct:
			if rv, b := fn(&field, tag, attr); b {
				fv.Vals = append(fv.Vals, rv)
			}
			fv.traverse(field, tag, attr, fn)
		case reflect.Slice:
			sl := tmp.Slice(0, 1)
			if sl.Kind() != reflect.Struct {
				break
			}
			fv.traverse(field, tag, attr, fn)
		}
	}
}
