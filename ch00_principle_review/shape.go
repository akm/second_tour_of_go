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

func (c *Circle) Expand(dr int) {
	c.radius += dr
}

type Rect struct {
	width, length int
	Point         *Point
}

func (r *Rect) Area() int {
	return r.width * r.length
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
