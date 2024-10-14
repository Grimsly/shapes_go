package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	// If using pointers for structs fields, you can check for nil
	// if t.base == nil {
	// 	return 4 * t.height * 0.5
	// }
	return t.base * t.height * 0.5
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	// var base float64
	// base = 2
	t := triangle{height: 2}
	s := square{4}

	printArea(t)
	printArea(s)
}
