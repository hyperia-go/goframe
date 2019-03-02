package tests

import (
	"frame/element"
	"testing"
)

func TestSeries_Init(t *testing.T) {
	_, err := createSmallNumericIntSeries(t)
	if err != nil {
		t.Errorf("%+v", err)
	}
}

func TestSeries_InitEmpty(t *testing.T) {
	_, err := createEmptySeries(t)
	if err.Error() != "EmptyFrame" {
		t.Errorf("Exit status not correct %+v", err)
	}
}

func TestSeries_Get(t *testing.T) {
	series, _ := createSmallNumericIntSeries(t)
	vals := series.Get()
	s := make([]interface{}, len(SmallDataInt))
	for i, v := range SmallDataInt {
		s[i] = v
	}
	if !Eq(s, vals) {
		t.Errorf("Expected %+v, got %+v", SmallDataInt, vals)
	}
}

func TestSeries_Get2(t *testing.T) {
	series, err := createLargeNumericIntSeries(t)
	vals := series.Get()
	if err != nil {
		t.Errorf("%+v", err)
	}
	s := make([]interface{}, len(LargeDataInt))
	for i, v := range LargeDataInt {
		s[i] = v
	}
	if !Eq(s, vals) {
		t.Errorf("Failed. Large arrays did not match")
	}
}

func TestSeries_Index(t *testing.T) {
	series, err := createLargeNumericIntSeries(t)
	if err != nil {
		t.Errorf("%+v", err)
	}
	index := series.Index
	expected := make([]interface{}, series.Len())
	got := make([]interface{}, series.Len())
	for i, v := range index {
		expected[i] = i
		got[i] = v
	}
	if !Eq(got, expected) {
		t.Error("Indexes do not match")
	}
}

func TestSeries_ResetIndex(t *testing.T) {
	series, err := createLargeNumericIntSeries(t)
	if err != nil {
		t.Errorf("%+v", err)
	}
	series.Index[100] = 12345654321
	index := series.Index
	expected := make([]interface{}, series.Len())
	got := make([]interface{}, series.Len())
	for i, v := range index {
		expected[i] = i
		got[i] = v
	}
	if Eq(got, expected) {
		t.Error("Indexes should not match")
	}

	// Test not inplace index reset
	series.ResetIndex(false)
	index = series.Index
	expected = make([]interface{}, series.Len())
	got = make([]interface{}, series.Len())
	for i, v := range index {
		expected[i] = i
		got[i] = v
	}
	if Eq(got, expected) {
		t.Error("Indexes should not match during non inplace reset of index")
	}


	// Test inplace index reset
	series.ResetIndex(true)
	index = series.Index
	expected = make([]interface{}, series.Len())
	got = make([]interface{}, series.Len())
	for i, v := range index {
		expected[i] = i
		got[i] = v
	}
	if !Eq(got, expected) {
		t.Error("Indexes do not match")
	}

}

func TestSeries_Sum(t *testing.T) {

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
	if err_float != nil || !FloatEq(sum_float, expected_float) {
		t.Errorf("Error: %+v (<nil> expected). Expected sum %f got %f", err_float, expected_float, sum_float)
	}
}

func TestSeries_Sum2(t *testing.T) {

	// Test sumInt with large ints
	s_int, _ := createLargeNumericIntSeries(t)
	sumInt, errInt := s_int.Sum()

	sumIntExpected := float64(0)
	for _, v := range LargeDataInt {
		sumIntExpected += float64(v)
	}

	if errInt != nil || float64(sumIntExpected) != sumInt {
		t.Errorf("Error: %+v (<nil> expected). Expected sum %f got %f", errInt, sumIntExpected, sumInt)
	}

	// Test sumInt with large floats
	s_float, _ := createLargeNumericFloatSeries(t)
	sumFloat, errFloat := s_float.Sum()

	sumFloatExpected := float64(0)
	for _, v := range LargeDataFloat {
		sumFloatExpected += float64(v)
	}

	if errFloat != nil || !FloatEq(sumFloat, sumFloatExpected) {
		t.Errorf("Error: %+v (<nil> expected). Expected sum %f got %f", errInt, sumFloatExpected, sumFloat)
	}
}

func TestSeries_Product(t *testing.T) {

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

	if err_float != nil || !FloatEq(prod_float, expected_float) {
		t.Errorf("Error: %+v (<nil> expected). Expected product %f got %f", err_float, expected_float, prod_float)
	}
}

func TestSeries_Product2(t *testing.T) {

	// Test prodInt with large ints
	s_int, _ := createLargeNumericIntSeries(t)
	prodInt, errInt := s_int.Product()

	prodIntExpected := float64(0)
	for _, v := range LargeDataInt {
		prodIntExpected *= float64(v)
	}

	if errInt != nil || float64(prodIntExpected) != prodInt {
		t.Errorf("Error: %+v (<nil> expected). Expected prod %f got %f", errInt, prodIntExpected, prodInt)
	}

	// Test prodInt with large floats
	s_float, _ := createLargeNumericFloatSeries(t)
	prodFloat, errFloat := s_float.Product()

	prodFloatExpected := float64(0)
	for _, v := range LargeDataFloat {
		prodFloatExpected *= float64(v)
	}

	if errFloat != nil || !FloatEq(prodFloat, prodFloatExpected) {
		t.Errorf("Error: %+v (<nil> expected). Expected prod %f got %f", errInt, prodFloatExpected, prodFloat)
	}
}

func TestSeries_Max(t *testing.T) {
	s, _ := createSmallNumericIntSeries(t)
	max, _ := s.Max()
	res, _ := element.Eq(max, SmallMax)
	if !res.Val.(bool) {
		t.Errorf("Expected 5.0, got %f", max.Val.(float64))
	}
}

func TestSeries_Argmax(t *testing.T) {
	s, _ := createSmallNumericIntSeries(t)
	index, _ := s.Argmax()
	if index != 4 {
		t.Errorf("Expected index 4, got %d", index)
	}
}

func TestSeries_Min(t *testing.T) {
	s, _ := createSmallNumericIntSeries(t)
	min, _ := s.Min()
	res, _ := element.Eq(min, SmallMax)
	if res.Val.(bool) {
		t.Errorf("Expected 1, got %f", min.Val.(float64))
	}
}

func TestSeries_Argmin(t *testing.T) {
	s, _ := createSmallNumericIntSeries(t)
	index, _ := s.Argmin()
	if index != 0 {
		t.Errorf("Expected 0, got %d", index)
	}
}

func TestSeries_CumSum(t *testing.T) {
	s, _ := createSmallNumericIntSeries(t)
	cumsum, err := s.Cumsum()
	got := cumsum.Get()
	if err != nil {
		t.Errorf("%+v", err)
	}
	expected := make([]interface{}, len(SmallDataInt))
	expected_int := [...]float64{1, 3, 6, 10, 15}
	for i, v := range expected_int {
		expected[i] = v
	}
	if !Eq(got, expected) {
		t.Errorf("Cumsum failed. Expected %+v, got %+v", expected, got)
	}
}

func TestSeries_CumSum2(t *testing.T) {
	s, _ := createSmallNumericFloatSeries(t)
	cumsum, err := s.Cumsum()
	got := cumsum.Get()
	if err != nil {
		t.Errorf("%+v", err)
	}
	expected := make([]interface{}, len(SmallDataInt))
	expected_int := [...]float64{0.1, 0.3, 0.6, 1.0, 1.5}
	for i, v := range expected_int {
		expected[i] = v
	}
	if !Eq(got, expected) {
		t.Errorf("Cumsum failed. Expected %+v, got %+v", expected, got)
	}
}
