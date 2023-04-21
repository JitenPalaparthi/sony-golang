package main

import (
	"fmt"
	"shape/rect"
)

func main() {

	r1 := new(rect.Rect)

	r1.L = 12.34
	r1.B = 13.45

	a1, p1 := r1.Area(), r1.Perimeter()

	fmt.Println("Area of r1:", a1, "Perimeter of r1:", p1)

	fmt.Println("Area of r1 inside r1:", r1.A)
	fmt.Println("Perimeter of r1 inside r1:", r1.P)

	s1 := new(rect.Square)
	s1.S = 25.25

	fmt.Println("Area of square:", s1.Area())

	// fmt.Println("Area of square in square:", s1.a) // does not work becase s1.a is unexported field

}
