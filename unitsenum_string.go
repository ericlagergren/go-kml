// generated by stringer -type=unitsEnum; DO NOT EDIT

package kml

import "fmt"

const _unitsEnum_name = "FractionPixelsInsetPixels"

var _unitsEnum_index = [...]uint8{0, 8, 14, 25}

func (i unitsEnum) String() string {
	if i < 0 || i+1 >= unitsEnum(len(_unitsEnum_index)) {
		return fmt.Sprintf("unitsEnum(%d)", i)
	}
	return _unitsEnum_name[_unitsEnum_index[i]:_unitsEnum_index[i+1]]
}
