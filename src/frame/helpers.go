package frame

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"testing"
)

// Test variables
var SmallDataInt = [...]int{1, 2, 3, 4, 5}
var SmallDataFloat = [...]float64{0.1, 0.2, 0.3, 0.4, 0.5}
var SmallDataString = [...]string{"s1", "s2", "s3", "s4", "s5"}

var LargeDataInt = intArange(0, 1000000, 4)
var LargeDataFloat = floatArange(0, 500000, 2)
var LargeDataString = strRepeat("s", 250000, true)

// ------------------------------------------
// Helper Functions -------------------------
// ------------------------------------------
func floatArange(start, stop, step float64) []float64 {
	N := int(math.Ceil((stop - start) / step))
	rnge := make([]float64, N, N)
	i := 0
	for x := start; x < stop; x += step {
		rnge[i] = x
		i += 1
	}
	return rnge
}

func intArange(start, stop, step float64) []int {
	N := int(math.Ceil((stop - start) / step))
	rnge := make([]int, N, N)
	i := 0
	for x := start; x < stop; x += step {
		rnge[i] = int(math.Floor(x))
		i += 1
	}
	return rnge
}

func strRepeat(s string, n int, unique bool) []string {
	rnge := make([]string, n)
	str := ""
	for x := 0; x < n; x++ {
		if unique {
			str = fmt.Sprintf("%s%d", s, x)
		} else {
			str = s
		}
		rnge[x] = str
	}
	return rnge
}

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

func FloatEq(a, b float64) bool {
	if a == 0 && b == 0 {
		return true
	}
	res := math.Abs(1 - a/b)
	return res < 0.00000001
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

// --------------------------------------------
// Helpers for Series -------------------------
// --------------------------------------------
func createSmallNumericIntSeries(t *testing.T) (Series, error) {
	s := make([]interface{}, len(SmallDataInt))
	for i, v := range SmallDataInt {
		s[i] = v
	}
	v, err := GoSeries("test", s, reflect.TypeOf(0))
	if err != nil {
		t.Errorf("Series init failed. Expected <nil> error, got: %+v", err)
	} else if v.Name != "test" {
		t.Errorf("Series init failed. Expected 'test' name, got: '%+v'", v.Name)
	}
	return v, err
}

func createSmallNumericFloatSeries(t *testing.T) (Series, error) {
	s := make([]interface{}, len(SmallDataFloat))
	for i, v := range SmallDataFloat {
		s[i] = v
	}
	v, err := GoSeries("test", s, reflect.TypeOf(0))
	if err != nil {
		t.Errorf("Series init failed. Expected <nil> error, got: %+v", err)
	}
	return v, err
}

func createSmallStringSeries(t *testing.T) (Series, error) {
	s := make([]interface{}, len(SmallDataString))
	for i, v := range SmallDataString {
		s[i] = v
	}
	v, err := GoSeries("test", s, reflect.TypeOf(0))
	if err != nil {
		t.Errorf("Series init failed. Expected <nil> error, got: %+v", err)
	}
	return v, err
}

func createEmptySeries(t *testing.T) (Series, error) {
	v, err := GoSeries("empty", make([]interface{}, 0), reflect.TypeOf(nil))
	if err.Error() != "EmptyFrame" {
		t.Errorf("%+v", err)
	}
	return v, err
}

func createLargeNumericIntSeries(t *testing.T) (Series, error) {
	s := make([]interface{}, len(LargeDataInt))
	for i, v := range LargeDataInt {
		s[i] = v
	}
	v, err := GoSeries("test", s, reflect.TypeOf(0))
	if err != nil {
		t.Errorf("Series init failed. Expected <nil> error, got: %+v", err)
	} else if v.Name != "test" {
		t.Errorf("Series init failed. Expected 'test' name, got: '%+v'", v.Name)
	}
	return v, err
}

func createLargeNumericFloatSeries(t *testing.T) (Series, error) {
	s := make([]interface{}, len(LargeDataFloat))
	for i, v := range LargeDataFloat {
		s[i] = v
	}
	v, err := GoSeries("test", s, reflect.TypeOf(0))
	if err != nil {
		t.Errorf("Series init failed. Expected <nil> error, got: %+v", err)
	} else if v.Name != "test" {
		t.Errorf("Series init failed. Expected 'test' name, got: '%+v'", v.Name)
	}
	return v, err
}

func createLargeStringSeries(t *testing.T) (Series, error) {
	s := make([]interface{}, len(LargeDataString))
	for i, v := range LargeDataString {
		s[i] = v
	}
	v, err := GoSeries("test", s, reflect.TypeOf(0))
	if err != nil {
		t.Errorf("Series init failed. Expected <nil> error, got: %+v", err)
	} else if v.Name != "test" {
		t.Errorf("Series init failed. Expected 'test' name, got: '%+v'", v.Name)
	}
	return v, err
}

// -------------------------------------------
// Helpers for Frame -------------------------
// -------------------------------------------
func createSmallTestFrame(t *testing.T) (Frame, error) {
	name1 := "Integers"
	c1 := make([]interface{}, len(SmallDataInt))
	for i, c := range SmallDataInt {
		c1[i] = c
	}

	name2 := "Floats"
	c2 := make([]interface{}, len(SmallDataFloat))
	for i, c := range SmallDataFloat {
		c2[i] = c
	}

	name3 := "Strings"
	c3 := make([]interface{}, len(SmallDataString))
	for i, c := range SmallDataString {
		c3[i] = c
	}

	series1, err1 := GoSeries(name1, c1, reflect.TypeOf(SmallDataInt[0]))
	series2, err2 := GoSeries(name2, c2, reflect.TypeOf(SmallDataFloat[0]))
	series3, err3 := GoSeries(name3, c3, reflect.TypeOf(SmallDataString[0]))

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
