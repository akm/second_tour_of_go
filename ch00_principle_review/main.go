package main

import "fmt"

func main() {
	p1 := &Point{Px: 100, Py: 50}
	fmt.Printf("%#v\n", p1)

	c1 := &Circle{Point: p1, radius: 30}
	fmt.Printf("%#v\n", c1)

	r1 := &Rect{Point: p1, width: 20, length: 10}
	fmt.Printf("%#v\n", r1)
}
