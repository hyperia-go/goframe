package frame

import (
	"reflect"
	"testing"
)

func createSmallTestFrame(t *testing.T) (Frame, error) {
	c1_data := [...]int{1, 2, 3, 4, 5}
	name1 := "Integers"
	c1 := make([]interface{}, len(c1_data))
	for i, c := range c1_data {
		c1[i] = c
	}

	c2_data := [...]float64{0.1, 0.2, 0.3, 0.4, 0.5}
	name2 := "Floats"
	c2 := make([]interface{}, len(c1_data))
	for i, c := range c2_data {
		c2[i] = c
	}

	c3_data := [...]string{"s1", "s2", "s3", "s4", "s5"}
	name3 := "Strings"
	c3 := make([]interface{}, len(c1_data))
	for i, c := range c3_data {
		c3[i] = c
	}

	series1, err1 := GoSeries(name1, c1, reflect.TypeOf(c1_data[0]))
	series2, err2 := GoSeries(name2, c2, reflect.TypeOf(c2_data[0]))
	series3, err3 := GoSeries(name3, c3, reflect.TypeOf(c3_data[0]))

	if err1 != nil || err2 != nil || err3 != nil {
		t.Errorf("%+v %+v %+v", err1, err2, err3)
	}

	s := [...]Series{series1, series2, series3}
	series := make([]interface{}, 3)
	for i, c := range s {
		series[i] = c
	}

	f, err := GoFrame(series)
	return f, err

}

func createSmallTestFrame2(t *testing.T) (Frame, error) {
	c1_data := [...]int{1, 2, 3, 4, 5}
	c1 := make([]interface{}, len(c1_data))
	for i, c := range c1_data {
		c1[i] = c
	}

	c2_data := [...]float64{0.1, 0.2, 0.3, 0.4, 0.5}
	c2 := make([]interface{}, len(c1_data))
	for i, c := range c2_data {
		c2[i] = c
	}

	c3_data := [...]string{"s1", "s2", "s3", "s4", "s5"}
	c3 := make([]interface{}, len(c1_data))
	for i, c := range c3_data {
		c3[i] = c
	}

	data := [...]interface{}{c1, c2, c3}
	series := make([]interface{}, 3)
	for i, c := range data {
		series[i] = c
	}

	f, err := GoFrame(series)
	if err != nil {
		t.Errorf("%+v", err)
	}
	return f, err
}

func TestFrameInitEmpty(t *testing.T) {
	_, err := GoFrame(nil)
	if err != nil {
		t.Errorf("%+v", err)
	}
}

func TestFrameInit(t *testing.T) {
	_, err := createSmallTestFrame(t)
	if err != nil {
		t.Errorf("%+v", err)
	}
}

func TestFrameColumnSelection(t *testing.T) {
	f, err := createSmallTestFrame(t)
	if err != nil {
		t.Errorf("%+v", err)
	}
	col := f.Get("Integers")
	expected := [...]int{1, 2, 3, 4, 5}
	s := make([]interface{}, len(expected))
	for i, v := range expected {
		s[i] = v
	}
	vals := col.Get()
	if !Eq(vals, s) {
		t.Errorf("Expected %+v, got %+v", expected, vals)
	}
}

func TestFrameInitGeneric(t *testing.T) {
	_, err := createSmallTestFrame2(t)
	if err != nil {
		t.Errorf("%+v", err)
	}
}
