package gotest

import "testing"

func TestSumFile(t *testing.T) {
	if i, e := SumFile("sample.txt"); e != nil {
		t.Error("fffff")
	}
}