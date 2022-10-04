package main

import "math"

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

func (c *Circle) Expand(dr int) {
	c.radius += dr
}

func (c *Circle) Area() float64 {
	return math.Pi * float64(c.radius*c.radius)
}

type Rect struct {
	width, length int
	Point         *Point
}

func (r *Rect) Area() float64 {
	return float64(r.width * r.length)
}

type RectList []*Rect

func (s RectList) Biggest() *Rect {
	var biggest *Rect
	for _, r := range s {
		if biggest == nil || r.Area() > biggest.Area() {
			biggest = r
		}
	}
	return biggest
}
