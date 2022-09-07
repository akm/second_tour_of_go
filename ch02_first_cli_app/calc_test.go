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

func TestDivision(t *testing.T) {
	calc := &Division{}
	if calc.String() != "division" {
		t.FailNow()
	}
	if dr, ok := calc.Do(1, 2).(*DivisionResult); ok {
		if dr.Quotient != 0 || dr.Remainder != 1 {
			t.FailNow()
		}
	} else {
		t.FailNow()
	}
	if dr, ok := calc.Do(8, 3).(*DivisionResult); ok {
		if dr.Quotient != 2 || dr.Remainder != 2 {
			t.FailNow()
		}
	} else {
		t.FailNow()
	}
	if _, ok := calc.Do(7, 0).(*InvalidDenominator); !ok {
		t.FailNow()
	}
}
