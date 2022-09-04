package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	s := Add(1, 2)
	if s != 3 {
		t.FailNow()
	}
}

func TestSubstract(t *testing.T) {
	if s := Substract(1, 2); s != -1 {
		t.FailNow()
	}
	if s := Substract(0, 3); s != -3 {
		t.FailNow()
	}
}
