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

func TestSubtract(t *testing.T) {
	if s := Subtract(1, 2); s != -1 {
		t.FailNow()
	}
	if s := Subtract(0, 3); s != -3 {
		t.FailNow()
	}
}
