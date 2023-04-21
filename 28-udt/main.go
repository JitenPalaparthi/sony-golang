package main

import "fmt"

func main() {

	var r1 Rect //declaring a  rect variable/object

	r2 := Rect{} // creating a rect variable/bject

	r3 := new(Rect) // creating a pointer rect variable/object

	r4 := &Rect{} // creating a pointer rect variable/objct

	fmt.Println(r1, r2, r3, r4)

	r5 := Rect{L: 10.23, B: 13.45} // creating a rect variable/object and assigning values

	var r6 Rect

	r6.L = 12.34
	r6.B = 13.45

	fmt.Println(r5, r6)

}

type Rect struct {
	L float32
	B float32
}
