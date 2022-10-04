package main

type Point struct {
	Px, Py int
}

type Circle struct {
	radius int
	Point  *Point
}

type Rect struct {
	width, length int
	Point         *Point
}
