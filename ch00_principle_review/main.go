package main

import "fmt"

func main() {
	p1 := new(Point)
	p1.Px = 100
	p1.Py = 50
	fmt.Printf("%#v\n", p1)

	c1 := new(Circle)
	c1.Point = p1
	c1.radius = 30
	fmt.Printf("%#v\n", c1)

	r1 := new(Rect)
	r1.Point = p1
	r1.width = 20
	r1.length = 10
	fmt.Printf("%#v\n", r1)
}
