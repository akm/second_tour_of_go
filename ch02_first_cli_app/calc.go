package main

type Calculation interface {
	Do(x, y int) int
	String() string
}

type Addition struct{}

var _ Calculation = (*Addition)(nil)

func (*Addition) String() string {
	return "addition"
}

func (*Addition) Do(x, y int) int {
	return x + y
}

type Subtraction struct{}

var _ Calculation = (*Subtraction)(nil)

func (*Subtraction) String() string {
	return "subtraction"
}

func (*Subtraction) Do(x, y int) int {
	return x - y
}
