// Another straightforward package to get use to interfaces implementation
package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func main() {
	sq := square{
		sideLength: 4,
	}
	tr := triangle{
		height: 3,
		base:   7,
	}

	printArea(sq)
	printArea(tr)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	area := s.sideLength * s.sideLength

	return area
}

func (t triangle) getArea() float64 {
	area := 0.5 * t.base * t.height

	return area
}
