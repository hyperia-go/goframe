package frame

import (
  "reflect"
  "errors"
)

var _ = reflect.TypeOf(0)

type Series struct {
  Name string
  Values Elements
  Type reflect.Type
}

type Elements struct{
  Vals []Element
}

type Element struct {
  Val interface{}
}

func (elements Elements) Len() int {
  return len(elements.Vals)
}

func (series Series) Get() []interface{} {
  ret := make([]interface{}, series.Values.Len())
  for i, v := range series.Values.Vals {
    ret[i] = v.Val
  }
  return ret
}

func SeriesInit(name string, data []interface{}, t reflect.Type) (Series, error) {
  if len(data) == 0 {
    return Series{}, errors.New("EmptyFrame")
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
