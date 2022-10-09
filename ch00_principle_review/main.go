package main

import "fmt"

func main() {
	p1 := &Point{Px: 100, Py: 50}
	fmt.Printf("%#v\n", p1)

	c1 := NewCircle(p1, 30)
	fmt.Printf("%#v\n", c1)

	c1.Expand(10)
	fmt.Printf("%#v\n", c1)

	r1 := &Rect{Point: p1, width: 20, length: 10}
	fmt.Printf("%#v\n", r1)

	shapes := Shapes{
		r1,
		c1,
		&Rect{Point: p1, width: 20, length: 5},
		&Rect{Point: p1, width: 50, length: 5},
		&Circle{Point: p1, radius: 20},
		&Rect{Point: p1, width: 10, length: 10},
	}
	fmt.Printf("shapes: %#v\n", shapes)
	fmt.Printf("Biggest: %#v\n", shapes.Biggest())
	shapes.MoveBy(10, 10)
	fmt.Printf("shapes: %#v\n", shapes)
}
