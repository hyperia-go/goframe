package datatypes

import (
	"reflect"
	"strings"
)

func DType(data []interface{}) reflect.Type {
	return reflect.TypeOf(data[0])
}

func IsNumeric(t reflect.Type) bool {
	typeString := t.String()
	numericTypes := [...]string{"int", "float", "complex"}
	for _, s := range numericTypes {
		if strings.Contains(typeString, s) {
			return true
		}
	}
	return false
}