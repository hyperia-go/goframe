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
	if !isNumeric(series.Type) {
		return float64(0), errors.New("ArithmeticError: can only add numeric types")
	}
	return series.Elements.Sum(), nil
}

func (series Series) Product() (float64, error) {
	if !isNumeric(series.Type) {
		return float64(0), errors.New("ArithmeticError: can only add numeric types")
	}
	return series.Elements.Prod(), nil
}

func (series Series) Max() (Element, error) {
	max, _ := series.Elements.Max()
	return max, nil
}

func (series Series) Argmax() (int, error) {
	_, index := series.Elements.Max()
	return index, nil
}
