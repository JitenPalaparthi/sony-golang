package main

import (
	"demo/shape"
	"demo/wrapper"

	"fmt"
)

func main() {
	r1 := new(shape.Rect)
	r1.L = shape.L1
	r1.B = shape.B1

	w := wrapper.New(r1)

	w.Area()
	w.Perimeter()

	// fmt.Println(r1.Area())

	//fmt.Println(shape.h1)
	Greet()
}

func init() {
	fmt.Println("Calling main init")
}
