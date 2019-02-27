package frame

import (
  "testing"
  "reflect"
)

var _ = reflect.TypeOf(0)

func Eq(a, b []interface{}) bool {

    // If one is nil, the other must also be nil.
    if (a == nil) != (b == nil) {
        return false;
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

func createSmallSeries(t *testing.T) (Series, error) {
  data := [...]int {1, 2, 3, 4, 5}
  s := make([]interface{}, len(data))
  for i, v := range data {
    s[i] = v
  }
  v, err := SeriesInit("test", s, reflect.TypeOf(0))
  if err != nil {
    t.Errorf("Series init failed. Expected <nil> error, got: %+v", err)
  }
  return v, err
}

func createEmptySeries(t *testing.T) (Series, error) {
  v, err := SeriesInit("empty", make([]interface{}, 0), reflect.TypeOf(nil))
  return v, err
}

func TestInit(t *testing.T) {
  createSmallSeries(t)
}

func TestEmpty(t *testing.T) {
  _, err := createEmptySeries(t)
  if err.Error() != "EmptyFrame" {
    t.Errorf("Exit status not correct %+v", err)
  }
}

func TestGet(t *testing.T) {
  series, _ := createSmallSeries(t)
  vals := series.Get()
  expected := [...]int{1, 2, 3, 4, 5}
  s := make([]interface{}, len(expected))
  for i, v := range expected {
    s[i] = v
  }
  if !Eq(s, vals) {
    t.Errorf("Expected %+v, got %+v", expected, vals)
  }

}
