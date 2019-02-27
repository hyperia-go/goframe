package frame

import (
  "reflect"
  "errors"
  "strings"
  // "fmt"
)
var _ = reflect.TypeOf(0)

// ------------------------------------------
// Structs and interfaces -------------------
// ------------------------------------------

type Series struct {
  Name string
  Values Elements
  Type reflect.Type
}

type SeriesInterface interface {

  // Get the values as an []interface{} object
  Get() []interface{}

  // Sum each element in the series
  Sum() float64
}

type Elements struct{
  Vals []Element
}

type Element struct {
  Val interface{}
}

// ------------------------------------------
// Helper Functions -------------------------
// ------------------------------------------
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

// ------------------------------------------
// Elements Methods -------------------------
// ------------------------------------------
func (elements Elements) Len() int {
  return len(elements.Vals)
}

// ------------------------------------------
// Series Methods ---------------------------
// ------------------------------------------
func SeriesInit(name string, data []interface{}, t reflect.Type) (Series, error) {
  if len(data) == 0 {
    return Series{Name: name}, errors.New("EmptyFrame")
  }
  s := Series{Name: name, Type: t}
  s_data := make([]Element, len(data))
  for i := range data {
    s_data[i] = Element{Val: data[i]}
  }
  elems := Elements{Vals: s_data}
  s.Values = elems
  return s, nil
}

func (series Series) Get() []interface{} {
  ret := make([]interface{}, series.Values.Len())
  for i, v := range series.Values.Vals {
    ret[i] = v.Val
  }
  return ret
}

func (series Series) Sum() (float64, error) {
  if !is_numeric(series.Type) {
      return float64(0), errors.New("ArithmeticError: can only add numeric types")
  }
  sum := float64(0)
  for _, x := range series.Get() {
    switch x.(type) {
    case uint8:
        sum += float64(x.(uint8))
    case int8:
        sum += float64(x.(uint8))
    case uint16:
        sum += float64(x.(uint16))
    case int16:
        sum += float64(x.(int16))
    case uint32:
        sum += float64(x.(uint32))
    case int32:
        sum += float64(x.(int32))
    case uint64:
        sum += float64(x.(uint64))
    case int64:
        sum += float64(x.(int64))
    case int:
        sum += float64(x.(int))
    case float32:
        sum += float64(x.(float32))
    case float64:
        sum += float64(x.(float64))
    default:
        return 0, errors.New("ArithmeticError: can only add numeric types")
     }
  }
  return sum, nil
}
