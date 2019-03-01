package frame

import (
	"reflect"
	"strings"
)

// ------------------------------------------
// Helper Functions -------------------------
// ------------------------------------------

func Eq(a, b []interface{}) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func is_numeric(t reflect.Type) bool {
	t_string := t.String()
	numeric_types := [...]string{"int", "float", "complex"}
	for _, s := range numeric_types {
		if strings.Contains(t_string, s) {
			return true
		}
	}
	return false
}
