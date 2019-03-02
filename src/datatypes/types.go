package datatypes

import (
	"reflect"
	"strings"
)

func DType(data []interface{}) string {
	return reflect.TypeOf(data[0]).String()
}

func IsNumeric(typeString string) bool {
	numericTypes := [...]string{"int", "float", "complex"}
	for _, s := range numericTypes {
		if strings.Contains(typeString, s) {
			return true
		}
	}
	return false
}