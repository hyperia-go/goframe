package tests

import (
	"frame"
	"math"
	"testing"
)

func TestElement_Add(t *testing.T) {
	a := frame.Element{Val: 100}
	b := frame.Element{Val: 200}
	got, err := frame.Add(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	expected := frame.Element{Val: a.Val.(int) + b.Val.(int)}
	res, err2 := frame.Eq(got, expected)
	if err2 != nil {
		t.Errorf("%+v", err2)
	}
	if !res.Val.(bool) {
		t.Errorf("Expected %+v, got %+v", expected.Val, got.Val)
	}
}

func TestElement_Add2(t *testing.T) {
	a := frame.Element{Val: float64(123.456)}
	b := frame.Element{Val: float64(-1234.56)}
	got, err := frame.Add(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	expected := frame.Element{Val: a.Val.(float64) + b.Val.(float64)}
	res, err2 := frame.Eq(got, expected)
	if err2 != nil {
		t.Errorf("%+v", err2)
	}
	if !res.Val.(bool) {
		t.Errorf("Expected %+v, got %+v", expected.Val, got.Val)
	}
}

func TestElement_Sub(t *testing.T) {
	a := frame.Element{Val: 100}
	b := frame.Element{Val: 200}
	got, err := frame.Diff(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	expected := frame.Element{Val: a.Val.(int) - b.Val.(int)}
	res, err2 := frame.Eq(got, expected)
	if err2 != nil {
		t.Errorf("%+v", err2)
	}
	if !res.Val.(bool) {
		t.Errorf("Expected %+v, got %+v", expected.Val, got.Val)
	}
}

func TestElement_Sub2(t *testing.T) {
	a := frame.Element{Val: float64(123.456)}
	b := frame.Element{Val: float64(-1234.56)}
	got, err := frame.Diff(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	expected := frame.Element{Val: a.Val.(float64) - b.Val.(float64)}
	res, err2 := frame.Eq(got, expected)
	if err2 != nil {
		t.Errorf("%+v", err2)
	}
	if !res.Val.(bool) {
		t.Errorf("Expected %+v, got %+v", expected.Val, got.Val)
	}
}

func TestElement_Divide(t *testing.T) {
	a := frame.Element{Val: 100}
	b := frame.Element{Val: 200}
	got, err := frame.Quot(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}

	// Performs float division by design
	expected := frame.Element{Val: 0.5}
	res, err2 := frame.Eq(got, expected)
	if err2 != nil {
		t.Errorf("%+v", err2)
	}
	if !res.Val.(bool) {
		t.Errorf("Expected %+v, got %+v", expected.Val, got.Val)
	}
}

func TestElement_Divide2(t *testing.T) {
	a := frame.Element{Val: float64(123.456)}
	b := frame.Element{Val: float64(-1234.56)}
	got, err := frame.Quot(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	expected := frame.Element{Val: a.Val.(float64) / b.Val.(float64)}
	res, err2 := frame.Eq(got, expected)
	if err2 != nil {
		t.Errorf("%+v", err2)
	}
	if !res.Val.(bool) {
		t.Errorf("Expected %+v, got %+v", expected.Val, got.Val)
	}
}

func TestElement_Mod(t *testing.T) {
	a := frame.Element{Val: float64(100)}
	b := frame.Element{Val: float64(200)}
	got, err := frame.Mod(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	expected := frame.Element{Val: math.Mod(a.Val.(float64), b.Val.(float64))}
	res, err2 := frame.Eq(got, expected)
	if err2 != nil {
		t.Errorf("%+v", err2)
	}
	if !res.Val.(bool) {
		t.Errorf("Expected %+v, got %+v", expected.Val, got.Val)
	}
}

func TestElement_Mod2(t *testing.T) {
	a := frame.Element{Val: float64(123.456)}
	b := frame.Element{Val: float64(-1234.56)}
	got, err := frame.Mod(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	expected := frame.Element{Val: math.Mod(a.Val.(float64), b.Val.(float64))}
	res, err2 := frame.Eq(got, expected)
	if err2 != nil {
		t.Errorf("%+v", err2)
	}
	if !res.Val.(bool) {
		t.Errorf("Expected %+v, got %+v", expected.Val, got.Val)
	}
}

func TestElement_Equals(t *testing.T) {
	a := frame.Element{Val: 100}
	b := frame.Element{Val: 200}
	got, err := frame.Eq(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	if got.Val.(bool) {
		t.Errorf("Comparing %+v and %+v. Expected false, was true", a.Val, b.Val)
	}
}

func TestElement_Equals2(t *testing.T) {
	a := frame.Element{Val: float64(123.456)}
	b := frame.Element{Val: float64(123.456)}
	got, err := frame.Eq(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	if !got.Val.(bool) {
		t.Errorf("Comparing %+v and %+v. Expected true, was false", a.Val, b.Val)
	}
}

func TestElement_EqualsStrings(t *testing.T) {
	a := frame.Element{Val: "yes"}
	b := frame.Element{Val: "no"}
	got, err := frame.Eq(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	if got.Val.(bool) {
		t.Errorf("Comparing '%+v' and '%+v'. Expected false, was true", a.Val, b.Val)
	}
}

func TestElement_EqualsStrings2(t *testing.T) {
	a := frame.Element{Val: "yes"}
	b := frame.Element{Val: "yes"}
	got, err := frame.Eq(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	if !got.Val.(bool) {
		t.Errorf("Comparing '%+v' and '%+v'. Expected true, was false", a.Val, b.Val)
	}
}

func TestElement_EqualsStrings3(t *testing.T) {
	a := frame.Element{Val: ""}
	b := frame.Element{Val: ""}
	got, err := frame.Eq(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	if !got.Val.(bool) {
		t.Errorf("Comparing '%+v' and '%+v'. Expected true, was false", a.Val, b.Val)
	}
}

func TestElement_EqualsStrings4(t *testing.T) {
	a := frame.Element{Val: "no"}
	b := frame.Element{Val: ""}
	got, err := frame.Eq(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	if got.Val.(bool) {
		t.Errorf("Comparing '%+v' and '%+v'. Expected false, was true", a.Val, b.Val)
	}
}