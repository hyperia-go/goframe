package tests

import (
	"frame"
	"testing"
)

func TestIo_readCsv1(t *testing.T) {
	_, err := frame.ReadCsv("data/test1.csv")
	if err != nil {
		t.Errorf("%+v", err)
	}
}
