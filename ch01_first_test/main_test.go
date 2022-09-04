package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	s := 1 + 2
	if s != 3 {
		t.FailNow()
	}
}
