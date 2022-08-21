package main

import (
	"testing"
)

func TestGolang(t *testing.T) {
	s := 1 + 2
	if s != 3 {
		t.FailNow()
	}
}
