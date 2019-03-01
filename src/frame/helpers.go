package frame

import (
	"math"
	"reflect"
	"strings"
)

// ------------------------------------------
// Operations -------------------------------
// ------------------------------------------
var Ops = map[string]func(float64, float64) Element{
	"+":  func(a, b float64) Element { return Element{Val: a + b} },
	"-":  func(a, b float64) Element { return Element{Val: a - b} },
	"*":  func(a, b float64) Element { return Element{Val: a * b} },
	"/":  func(a, b float64) Element { return Element{Val: a / b} },
	"%":  func(a, b float64) Element { return Element{Val: math.Mod(a, b) }},
	"<":  func(a, b float64) Element { return Element{Val: a < b} },
	"<=": func(a, b float64) Element { return Element{Val: a <= b} },
	">":  func(a, b float64) Element { return Element{Val: a > b} },
	">=": func(a, b float64) Element { return Element{Val: a >= b} },
}

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