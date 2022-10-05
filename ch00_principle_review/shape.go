package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	MoveBy(dx, dy int)
	BasePoint() *Point
}

type Shapes []Shape

func (s Shapes) Biggest() Shape {
	var biggest Shape
	for _, sh := range s {
		if biggest == nil || sh.Area() > biggest.Area() {
			biggest = sh
		}
	}
	return biggest
}

func (s Shapes) MoveBy(dx, dy int) {
	/**
	 * この実装だと、同じPointのインスタンスを共有している Rect あるいは Circle の Point に対して
	 * 何度も MoveBy が呼ばれてしまい、その都度 Point の値に加算し、想定していない結果になることがある。
	 */
	for _, sh := range s {
		sh.MoveBy(dx, dy)
	}
}

type Point struct {
	Px, Py int
}

func (p *Point) MoveBy(dx, dy int) {
	p.Px += dx
	p.Py += dy
}

func (p *Point) GoString() string {
	return fmt.Sprintf("Point{Px: %d, Py: %d}", p.Px, p.Py)
}

type Points []*Point

func (s Points) Include(p *Point) bool {
	for _, i := range s {
		if i == p {
			return true
		}
	}
	return false
}

type Circle struct {
	radius int
	Point  *Point
}

var _ Shape = (*Circle)(nil)

func NewCircle(p *Point, r int) *Circle {
	return &Circle{Point: p, radius: r}
}

func (c *Circle) BasePoint() *Point {
	return c.Point
}

func (c *Circle) Expand(dr int) {
	c.radius += dr
}

func (c *Circle) Area() float64 {
	return math.Pi * float64(c.radius*c.radius)
}

func (c *Circle) MoveBy(dx, dy int) {
	c.Point.MoveBy(dx, dy)
}

func (c *Circle) GoString() string {
	return fmt.Sprintf("%#v", *c)
}

type Rect struct {
	width, length int
	Point         *Point
}

var _ Shape = (*Rect)(nil)

func (c *Rect) BasePoint() *Point {
	return c.Point
}

func (r *Rect) Area() float64 {
	return float64(r.width * r.length)
}

func (r *Rect) MoveBy(dx, dy int) {
	r.Point.MoveBy(dx, dy)
}

func (r *Rect) GoString() string {
	return fmt.Sprintf("%#v", *r)
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
