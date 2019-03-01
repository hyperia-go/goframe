package frame

import (
	"testing"
)

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
