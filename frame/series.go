package frame

import (
	"errors"
	"reflect"
	// "fmt"
)

var _ = reflect.TypeOf(0)

// ------------------------------------------
// Structs and interfaces -------------------
// ------------------------------------------
type Series struct {
	Name     string
	Elements Elements
	Type     reflect.Type
}

type SeriesInterface interface {

	// Get the values as an []interface{} object
	Get() []interface{}

	// Sum each element in the series
	Sum() float64
}

type Elements struct {
	Vals []Element
}

type Element struct {
	Val interface{}
}

var Ops = map[string]func(float64, float64) Element{
	"+":  func(a, b float64) Element { return Element{Val: a + b} },
	"-":  func(a, b float64) Element { return Element{Val: a - b} },
	"*":  func(a, b float64) Element { return Element{Val: a * b} },
	"/":  func(a, b float64) Element { return Element{Val: a / b} },
	"<":  func(a, b float64) Element { return Element{Val: a < b} },
	"<=": func(a, b float64) Element { return Element{Val: a <= b} },
	">":  func(a, b float64) Element { return Element{Val: a > b} },
	">=": func(a, b float64) Element { return Element{Val: a >= b} },
}

// ------------------------------------------
// Elements Methods -------------------------
// ------------------------------------------
func (elements Elements) Len() int {
	return len(elements.Vals)
}

func (elements Elements) Sum() float64 {
	sum := Element{Val: float64(0)}
	for _, x := range elements.Vals {
		sum, _ = sum.Add(x)
	}
	return sum.Val.(float64)
}

func (elements Elements) Prod() float64 {
	prod := Element{Val: float64(1)}
	for _, x := range elements.Vals {
		prod, _ = prod.Prod(x)
	}
	return prod.Val.(float64)
}

// ------------------------------------------
// Element Methods --------------------------
// ------------------------------------------

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

func (e Element) Op(x Element, op string) (Element, error) {

	// Ensure e has float value
	err_e := e.AsFloat()
	if err_e != nil {
		panic(err_e)
	}

	// Ensure x has float value
	err_x := x.AsFloat()
	if err_x != nil {
		panic(err_x)
	}

	// Return no error
	return Ops[op](e.Val.(float64), x.Val.(float64)), nil
}

func (e Element) Add(x Element) (Element, error) {
	return e.Op(x, "+")
}

func (e Element) Prod(x Element) (Element, error) {
	return e.Op(x, "*")
}

// ------------------------------------------
// Series Methods ---------------------------
// ------------------------------------------
func GoSeries(name string, data []interface{}, t reflect.Type) (Series, error) {
	if len(data) == 0 {
		return Series{Name: name}, errors.New("EmptyFrame")
	}
	s := Series{Name: name, Type: t}
	s_data := make([]Element, len(data))
	for i := range data {
		s_data[i] = Element{Val: data[i]}
	}
	elems := Elements{Vals: s_data}
	s.Elements = elems
	return s, nil
}

func (series Series) Get() []interface{} {
	ret := make([]interface{}, series.Elements.Len())
	for i, v := range series.Elements.Vals {
		ret[i] = v.Val
	}
	return ret
}

func (series Series) Sum() (float64, error) {
	if !is_numeric(series.Type) {
		return float64(0), errors.New("ArithmeticError: can only add numeric types")
	}
	return series.Elements.Sum(), nil
}

func (series Series) Product() (float64, error) {
	if !is_numeric(series.Type) {
		return float64(0), errors.New("ArithmeticError: can only add numeric types")
	}
	return series.Elements.Prod(), nil
}
