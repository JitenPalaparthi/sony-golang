package main

import "fmt"

func main() {

	r1 := Rect{L: 12.45, B: 13.46}

	// a1, p1 := Area(r1), Perimeter(r1)

	// fmt.Println("Area of r1", a1, "Perimeter of r1", p1)

	a2, p2 := r1.Area(), r1.Perimeter()

	fmt.Println("Area of r1", a2, "Perimeter of r1", p2)

	fmt.Println("Area of r1 in r1", r1.a, "Perimeter of r1 in r1", r1.p)

}

type Rect struct {
	L, B float32
	a, p float32
}

func (r *Rect) Area() float32 {
	r.a = r.L * r.B
	return r.a
}

func (r *Rect) Perimeter() float32 {
	r.p = 2 * (r.L + r.B)
	return r.p
}

func Area(r Rect) float32 {
	return r.L * r.B
}

func Perimeter(r Rect) float32 {
	return 2 * (r.L + r.B)
}

// class Rect{
// 	float L
// 	float B

// 	public float Area()
// }
