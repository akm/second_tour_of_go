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

	r2 := &Rect{Point: p1, width: 20, length: 5}
	r3 := &Rect{Point: p1, width: 50, length: 5}
	r4 := &Rect{Point: p1, width: 10, length: 10}
	fmt.Printf("Biggest: %#v\n", RectList([]*Rect{r1, r2, r3, r4}).Biggest())
}
