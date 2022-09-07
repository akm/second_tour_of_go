package main

import "testing"

func TestAddition(t *testing.T) {
	calc := &Addition{}
	if calc.String() != "addition" {
		t.FailNow()
	}
	if calc.Do(1, 2) != Int(3) {
		t.FailNow()
	}
	if calc.Do(1, -2) != Int(-1) {
		t.FailNow()
	}
}

func TestSubtraction(t *testing.T) {
	calc := &Subtraction{}
	if calc.String() != "subtraction" {
		t.FailNow()
	}
	if calc.Do(1, 2) != Int(-1) {
		t.FailNow()
	}
	if calc.Do(1, -2) != Int(3) {
		t.FailNow()
	}
}

func TestMultiplication(t *testing.T) {
	calc := &Multiplication{}
	if calc.String() != "multiplication" {
		t.FailNow()
	}
	if calc.Do(1, 2) != Int(2) {
		t.FailNow()
	}
	if calc.Do(3, 0) != Int(0) {
		t.FailNow()
	}
	if calc.Do(4, -1) != Int(-4) {
		t.FailNow()
	}
}
