package main

type Point struct {
	Px, Py int
}

type Circle struct {
	radius int
	Point  *Point
}

func NewCircle(p *Point, r int) *Circle {
	return &Circle{Point: p, radius: r}
}

type Rect struct {
	width, length int
	Point         *Point
}
