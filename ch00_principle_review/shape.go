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

func ExpandCircle(c *Circle, dr int) {
	c.radius += dr
}

type Rect struct {
	width, length int
	Point         *Point
}
