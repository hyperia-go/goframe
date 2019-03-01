package frame

import (
	"reflect"
	"testing"
)

func createSmallNumericIntSeries(t *testing.T) (Series, error) {
	data := [...]int{1, 2, 3, 4, 5}
	s := make([]interface{}, len(data))
	for i, v := range data {
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
	data := [...]float64{0.1, 0.2, 0.3, 0.4, 0.5}
	s := make([]interface{}, len(data))
	for i, v := range data {
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
	if err != nil {
		t.Errorf("%+v", err)
	}
	return v, err
}

func TestSeriesInit(t *testing.T) {
	_, err := createSmallNumericIntSeries(t)
	if err != nil {
		t.Errorf("%+v", err)
	}
}

func TestSeriesInitEmpty(t *testing.T) {
	_, err := createEmptySeries(t)
	if err.Error() != "EmptyFrame" {
		t.Errorf("Exit status not correct %+v", err)
	}
}

func TestSeriesGet(t *testing.T) {
	series, _ := createSmallNumericIntSeries(t)
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

func TestSeriesSum(t *testing.T) {

	// Test sum with ints
	s_int, _ := createSmallNumericIntSeries(t)
	sum, err := s_int.Sum()
	expected := float64(1 + 2 + 3 + 4 + 5)
	if err != nil || sum != expected {
		t.Errorf("Error: %+v (<nil> expected). Expected sum %f got %f", err, expected, sum)
	}

	// Test sum with floats
	s_float, _ := createSmallNumericFloatSeries(t)
	sum_float, err_float := s_float.Sum()
	expected_float := float64(0.1 + 0.2 + 0.3 + 0.4 + 0.5)
	if err_float != nil || sum_float != expected_float {
		t.Errorf("Error: %+v (<nil> expected). Expected sum %f got %f", err_float, expected_float, sum_float)
	}
}

func TestSeriesProduct(t *testing.T) {

	// Test product with ints
	s_int, _ := createSmallNumericIntSeries(t)
	prod, err := s_int.Product()
	expected := float64(1 * 2 * 3 * 4 * 5)
	if err != nil || prod != expected {
		t.Errorf("Error: %+v (<nil> expected). Expected product %f got %f", err, expected, prod)
	}

	// Test product with floats
	s_float, _ := createSmallNumericFloatSeries(t)
	prod_float, err_float := s_float.Product()
	expected_float := float64(0.1 * 0.2 * 0.3 * 0.4 * 0.5)

	// TODO: deal with floating point precision problems.
	if err_float != nil || prod_float-expected_float > 0.0001 {
		t.Errorf("Error: %+v (<nil> expected). Expected product %f got %f", err_float, expected_float, prod_float)
	}
}
