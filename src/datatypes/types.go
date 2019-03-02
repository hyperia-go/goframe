package datatypes

import (
	"reflect"
	"strings"
)

func DType(data []interface{}) reflect.Kind {
	return reflect.TypeOf(data[0]).Kind()
}

func IsNumeric(t reflect.Kind) bool {
	typeString := t.String()
	numericTypes := [...]string{"int", "float", "complex"}
	for _, s := range numericTypes {
		if strings.Contains(typeString, s) {
			return true
		}
	}
	return false
}