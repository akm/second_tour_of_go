package main

import (
	"testing"
)

func TestGolang(t *testing.T) {
	s := Add(1, 2)
	if s != 3 {
		t.FailNow()
	}
}

func TestSubstract(t *testing.T) {
	s := Substract(6, 2)
	if s != 4 {
		t.FailNow()
	}
}
