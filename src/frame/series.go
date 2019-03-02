package frame

import (
	"datatypes"
	"errors"
	"fmt"
	"frame/element"
	"reflect"
)

var _ = reflect.TypeOf(0)

// ------------------------------------------
// Structs and interfaces -------------------
// ------------------------------------------
type Series struct {
	Name     string
	Elements element.ElementArray
	Index 	 []int
	Type     reflect.Kind
}

// ------------------------------------------
// Series Methods ---------------------------
// ------------------------------------------
func GoSeries(name string, data []interface{}) (Series, error) {
	if len(data) == 0 {
		return Series{Name: name}, errors.New("EmptyFrame")
	}

	// TODO - Implement this properly.
	t := datatypes.DType(data)

	// Create the series
	s := Series{Name: name, Type: t}
	s_data := make([]element.Element, len(data))
	index := make([]int, len(data))
	for i := range data {
		s_data[i] = element.New(data[i])
		index[i] = i
	}
	elems := element.ElementArray{Vals: s_data}
	s.Elements = elems
	s.Index = index
	return s, nil
}

func (series Series) Get() []interface{} {
	ret := make([]interface{}, series.Elements.Len())
	for i, v := range series.Elements.Vals {
		ret[i] = v.Val
	}
	return ret
}

func (series* Series) Len() int {
	return len(series.Elements.Vals)
}

func (series* Series) Sum() (float64, error) {
	if !datatypes.IsNumeric(series.Type) {
		return float64(0), errors.New("ArithmeticError: can only add numeric types")
	}
	return series.Elements.Sum(), nil
}

func (series* Series) Product() (float64, error) {
	if !datatypes.IsNumeric(series.Type) {
		return float64(0), errors.New("ArithmeticError: can only add numeric types")
	}
	return series.Elements.Prod(), nil
}

func (series* Series) Max() (element.Element, error) {
	max, _ := series.Elements.Max()
	return max, nil
}

func (series* Series) Argmax() (int, error) {
	_, index := series.Elements.Max()
	return index, nil
}

func (series* Series) Min() (element.Element, error) {
	max, _ := series.Elements.Min()
	return max, nil
}

func (series* Series) Argmin() (int, error) {
	_, index := series.Elements.Min()
	return index, nil
}

func (series* Series) ResetIndex(inplace bool) *Series {
	index := make([]int, series.Len())
	for i := range index {
		index[i] = i
	}
	if inplace {
		series.Index = index
		return series
	} else {
		ret := Series{Elements: series.Elements, Type: series.Type, Name: series.Name, Index: index}
		return &ret
	}
}

func (series* Series) Append(e interface{}) {
	series.Elements.Append(e)
}

func (series *Series) Set(i int, e element.Element) {
	series.Elements.Vals[i] = e
}

func (series* Series) Cumsum() (Series, error) {
	sum := element.New(float64(0))
	s := Series{
		Name: series.Name,
		Elements: element.NewArr(series.Len()),
		Type: reflect.Float64,
	}
	fmt.Println(series.Elements)
	for i, v := range series.Elements.Vals {
		sum, _ = element.Add(sum, v)
		s.Set(i, sum)
	}
	s.ResetIndex(true)
	return s, nil
}