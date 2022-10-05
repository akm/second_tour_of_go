package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	MoveBy(dx, dy int)
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
