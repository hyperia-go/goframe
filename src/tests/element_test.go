package tests

import (
	"frame/element"
	"math"
	"testing"
)

func TestElement_Add(t *testing.T) {
	a := element.Element{Val: 100}
	b := element.Element{Val: 200}
	got, err := element.Add(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	expected := element.Element{Val: a.Val.(int) + b.Val.(int)}
	res, err2 := element.Eq(got, expected)
	if err2 != nil {
		t.Errorf("%+v", err2)
	}
	if !res.Val.(bool) {
		t.Errorf("Expected %+v, got %+v", expected.Val, got.Val)
	}
}

func TestElement_Add2(t *testing.T) {
	a := element.Element{Val: float64(123.456)}
	b := element.Element{Val: float64(-1234.56)}
	got, err := element.Add(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	expected := element.Element{Val: a.Val.(float64) + b.Val.(float64)}
	res, err2 := element.Eq(got, expected)
	if err2 != nil {
		t.Errorf("%+v", err2)
	}
	if !res.Val.(bool) {
		t.Errorf("Expected %+v, got %+v", expected.Val, got.Val)
	}
}

func TestElement_Sub(t *testing.T) {
	a := element.Element{Val: 100}
	b := element.Element{Val: 200}
	got, err := element.Diff(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	expected := element.Element{Val: a.Val.(int) - b.Val.(int)}
	res, err2 := element.Eq(got, expected)
	if err2 != nil {
		t.Errorf("%+v", err2)
	}
	if !res.Val.(bool) {
		t.Errorf("Expected %+v, got %+v", expected.Val, got.Val)
	}
}

func TestElement_Sub2(t *testing.T) {
	a := element.Element{Val: float64(123.456)}
	b := element.Element{Val: float64(-1234.56)}
	got, err := element.Diff(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	expected := element.Element{Val: a.Val.(float64) - b.Val.(float64)}
	res, err2 := element.Eq(got, expected)
	if err2 != nil {
		t.Errorf("%+v", err2)
	}
	if !res.Val.(bool) {
		t.Errorf("Expected %+v, got %+v", expected.Val, got.Val)
	}
}

func TestElement_Divide(t *testing.T) {
	a := element.Element{Val: 100}
	b := element.Element{Val: 200}
	got, err := element.Quot(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}

	// Performs float division by design
	expected := element.Element{Val: 0.5}
	res, err2 := element.Eq(got, expected)
	if err2 != nil {
		t.Errorf("%+v", err2)
	}
	if !res.Val.(bool) {
		t.Errorf("Expected %+v, got %+v", expected.Val, got.Val)
	}
}

func TestElement_Divide2(t *testing.T) {
	a := element.Element{Val: float64(123.456)}
	b := element.Element{Val: float64(-1234.56)}
	got, err := element.Quot(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	expected := element.Element{Val: a.Val.(float64) / b.Val.(float64)}
	res, err2 := element.Eq(got, expected)
	if err2 != nil {
		t.Errorf("%+v", err2)
	}
	if !res.Val.(bool) {
		t.Errorf("Expected %+v, got %+v", expected.Val, got.Val)
	}
}

func TestElement_Mod(t *testing.T) {
	a := element.Element{Val: float64(100)}
	b := element.Element{Val: float64(200)}
	got, err := element.Mod(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	expected := element.Element{Val: math.Mod(a.Val.(float64), b.Val.(float64))}
	res, err2 := element.Eq(got, expected)
	if err2 != nil {
		t.Errorf("%+v", err2)
	}
	if !res.Val.(bool) {
		t.Errorf("Expected %+v, got %+v", expected.Val, got.Val)
	}
}

func TestElement_Mod2(t *testing.T) {
	a := element.Element{Val: float64(123.456)}
	b := element.Element{Val: float64(-1234.56)}
	got, err := element.Mod(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	expected := element.Element{Val: math.Mod(a.Val.(float64), b.Val.(float64))}
	res, err2 := element.Eq(got, expected)
	if err2 != nil {
		t.Errorf("%+v", err2)
	}
	if !res.Val.(bool) {
		t.Errorf("Expected %+v, got %+v", expected.Val, got.Val)
	}
}

func TestElement_Equals(t *testing.T) {
	a := element.Element{Val: 100}
	b := element.Element{Val: 200}
	got, err := element.Eq(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	if got.Val.(bool) {
		t.Errorf("Comparing %+v and %+v. Expected false, was true", a.Val, b.Val)
	}
}

func TestElement_Equals2(t *testing.T) {
	a := element.Element{Val: float64(123.456)}
	b := element.Element{Val: float64(123.456)}
	got, err := element.Eq(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	if !got.Val.(bool) {
		t.Errorf("Comparing %+v and %+v. Expected true, was false", a.Val, b.Val)
	}
}

func TestElement_EqualsStrings(t *testing.T) {
	a := element.Element{Val: "yes"}
	b := element.Element{Val: "no"}
	got, err := element.Eq(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	if got.Val.(bool) {
		t.Errorf("Comparing '%+v' and '%+v'. Expected false, was true", a.Val, b.Val)
	}
}

func TestElement_EqualsStrings2(t *testing.T) {
	a := element.Element{Val: "yes"}
	b := element.Element{Val: "yes"}
	got, err := element.Eq(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	if !got.Val.(bool) {
		t.Errorf("Comparing '%+v' and '%+v'. Expected true, was false", a.Val, b.Val)
	}
}

func TestElement_EqualsStrings3(t *testing.T) {
	a := element.Element{Val: ""}
	b := element.Element{Val: ""}
	got, err := element.Eq(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	if !got.Val.(bool) {
		t.Errorf("Comparing '%+v' and '%+v'. Expected true, was false", a.Val, b.Val)
	}
}

func TestElement_EqualsStrings4(t *testing.T) {
	a := element.Element{Val: "no"}
	b := element.Element{Val: ""}
	got, err := element.Eq(a, b)
	if err != nil {
		t.Errorf("%+v", err)
	}
	if got.Val.(bool) {
		t.Errorf("Comparing '%+v' and '%+v'. Expected false, was true", a.Val, b.Val)
	}
}