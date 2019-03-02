package frame

import (
	"reflect"
	"strings"
)

// ------------------------------------------
// Operations -------------------------------
// ------------------------------------------
func isNumeric(t reflect.Type) bool {
	typeString := t.String()
	numericTypes := [...]string{"int", "float", "complex"}
	for _, s := range numericTypes {
		if strings.Contains(typeString, s) {
			return true
		}
	}
	return false
}