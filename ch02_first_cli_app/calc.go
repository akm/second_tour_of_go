package main

type Calculation interface {
	Do(x, y int) int
	String() string
}

type Addition struct{}

func (*Addition) String() string {
	return "addition"
}

func (*Addition) Do(x, y int) int {
	return x + y
}

type Subtraction struct{}

func (*Subtraction) String() string {
	return "subtraction"
}

func (*Subtraction) Do(x, y int) int {
	return x - y
}
