package main

import "strconv"

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
