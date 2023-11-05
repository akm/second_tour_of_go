package main

import (
	"fmt"
	"strconv"
)

type Result interface {
	String() string
}

type Int int

var _ Result = Int(0)

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

type Calculation interface {
	Do(x, y int) Result
	String() string
}

type Addition struct{}

var _ Calculation = (*Addition)(nil)

func (*Addition) String() string {
	return "addition"
}

func (*Addition) Do(x, y int) Result {
	return Int(x + y)
}

type Subtraction struct{}

var _ Calculation = (*Subtraction)(nil)

func (*Subtraction) String() string {
	return "subtraction"
}

func (*Subtraction) Do(x, y int) Result {
	return Int(x - y)
}

type Multiplication struct{}

var _ Calculation = (*Multiplication)(nil)

func (*Multiplication) String() string {
	return "multiplication"
}

func (*Multiplication) Do(x, y int) Result {
	return Int(x * y)
}

type DivisionResult struct {
	Quotient  int
	Remainder int
}

var _ Result = (*DivisionResult)(nil)

func (d *DivisionResult) String() string {
	return fmt.Sprintf("quotient: %d, remainder: %d", d.Quotient, d.Remainder)
}

type InvalidDenominator struct{}

var _ Result = (*InvalidDenominator)(nil)

func (i *InvalidDenominator) String() string {
	return "Invalid denominator. It must be not zero"
}

type Division struct{}

var _ Calculation = (*Division)(nil)

func (*Division) String() string {
	return "division"
}

func (*Division) Do(x, y int) Result {
	if y == 0 {
		return &InvalidDenominator{}
	}
	q, r := x/y, x%y
	return &DivisionResult{q, r}
}
