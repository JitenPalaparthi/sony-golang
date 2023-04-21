package shape

import "fmt"

type IShape interface {
	Area() float32
	Perimeter() float32
}

func init() {
	fmt.Println("Calling init-1")
}

func init() {
	fmt.Println("Calling init-2")
}

func init() {
	fmt.Println("Calling init-3")
	L1 = 12.34
	B1 = 13.45
	h1 = 14.56
}

var L1, B1, h1 float32

type Rect struct {
	L, B float32
}

func (r *Rect) Area() float32 {
	return r.L * r.B
}

func (r *Rect) Perimeter() float32 {
	return 2 * (r.L * r.B)
}

func (r *Rect) Display() {
	fmt.Println(r)
}

type Square float32

func (s Square) Area() float32 {
	return float32(s * s)
}

func (s Square) Perimeter() float32 {
	return 4 * float32(s)
}

type Cube struct {
	L, B, H float32
}

func (c *Cube) Area() float32 {
	return c.L * c.B * c.H
}

func (c *Cube) Perimeter() float32 {
	return 4 * (c.L + c.B + c.H)
}

type Circle struct {
	R float32
}

func (t *Circle) Area() float32 {

	return 3.14 * t.R * t.R
}

func (t *Circle) Perimeter() float32 {

	return 2 * 3.14 * t.R
}
