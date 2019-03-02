package element

import (
	"errors"
	"math"
	"reflect"
)

type Element struct {
	Val interface{}
	Type reflect.Type
}

var emptyElement = Element{Val: nil}

var Ops = map[string]func(Element, Element) (Element, error) {
	"+":  func(a, b Element) (Element, error) { return Element{Val: a.Val.(float64) + b.Val.(float64)}, nil },
	"-":  func(a, b Element) (Element, error) { return Element{Val: a.Val.(float64) - b.Val.(float64)}, nil },
	"*":  func(a, b Element) (Element, error) { return Element{Val: a.Val.(float64) * b.Val.(float64)}, nil },
	"/":  func(a, b Element) (Element, error) { return Element{Val: a.Val.(float64) / b.Val.(float64)}, nil },
	"%":  func(a, b Element) (Element, error) { return Element{Val: math.Mod(a.Val.(float64), b.Val.(float64))}, nil },
	"<":  func(a, b Element) (Element, error) { return Element{ Val: a.Val.(float64) < b.Val.(float64)}, nil },
	"<=": func(a, b Element) (Element, error) { return Element{ Val: a.Val.(float64) <= b.Val.(float64)}, nil },
	">":  func(a, b Element) (Element, error) { return Element{ Val: a.Val.(float64) > b.Val.(float64)}, nil },
	">=": func(a, b Element) (Element, error) { return Element{ Val: a.Val.(float64) >= b.Val.(float64)}, nil },
	"==": func(a, b Element) (Element, error) { return Element{ Val: a.Val == b.Val}, nil },
}

// ------------------------------------------
// Element Methods --------------------------
// ------------------------------------------
func New(value interface{}) Element {
	return Element{Val: value, Type: reflect.TypeOf(value)}
}
func (e *Element) AsFloat() error {
	switch e.Val.(type) {
	case uint8:
		e.Val = float64(e.Val.(uint8))
	case int8:
		e.Val = float64(e.Val.(uint8))
	case uint16:
		e.Val = float64(e.Val.(uint16))
	case int16:
		e.Val = float64(e.Val.(int16))
	case uint32:
		e.Val = float64(e.Val.(uint32))
	case int32:
		e.Val = float64(e.Val.(int32))
	case uint64:
		e.Val = float64(e.Val.(uint64))
	case int64:
		e.Val = float64(e.Val.(int64))
	case int:
		e.Val = float64(e.Val.(int))
	case float32:
		e.Val = float64(e.Val.(float32))
	case float64:
		e.Val = float64(e.Val.(float64))
	default:
		return errors.New("ArithmeticError: can only add numeric types")
	}
	return nil
}


// --------------------------------------------
// Element Functions --------------------------
// --------------------------------------------
func Op(e, x Element, op string) (Element, error) {

	// Ensure e has float value - strings cannot be cast to floats, so ignore these errors
	err_e := e.AsFloat()
	if err_e != nil && op != "==" {
		return emptyElement, err_e
	}

	// Ensure x has float value - strings cannot be cast to floats, so ignore these errors
	err_x := x.AsFloat()
	if err_x != nil && op != "==" {
		return emptyElement, err_x
	}

	// Return no error
	//fmt.Println(e)
	//fmt.Println(x)
	return Ops[op](e, x)
}

func Add(e, x Element) (Element, error) {
	return Op(e, x, "+")
}

func Diff(e, x Element) (Element, error) {
	return Op(e, x, "-")
}

func Prod(e, x Element) (Element, error) {
	return Op(e, x, "*")
}

func Quot(e, x Element) (Element, error) {
	return Op(e, x, "/")
}

func Mod(e, x Element) (Element, error) {
	return Op(e, x, "%")
}

func Eq(e, x Element) (Element, error) {
	return Op(e, x, "==")
}

func Le(e, x Element) (Element, error) {
	return Op(e, x, "<")
}

func Leq(e, x Element) (Element, error) {
	return Op(e, x, "<=")
}

func Ge(e, x Element) (Element, error) {
	return Op(e, x, ">")
}

func Geq(e, x Element) (Element, error) {
	return Op(e, x, ">=")
}
