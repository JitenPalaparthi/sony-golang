package main

import (
	"fmt"
)

type IArea interface {
	Area() float32
}

type IPerimeter interface {
	Perimeter() float32
}

type IShape interface {
	IArea
	IPerimeter
}

func main() {
	r1 := Rect{L: 12.34, B: 13.45}
	c1 := Cube{L: 10.34, B: 8.95, H: 10.45}
	var s1 Square = 23.34
	//var s2 Square = 45.54

	slice := make([]IShape, 4)
	slice[0] = &r1
	slice[1] = &c1
	slice[2] = s1
	slice[3] = &Circle{R: 23.45}

	for _, v := range slice {
		AreaAndPerimeter(v)
	}

	var iface1 IShape = &Rect{L: 123.34, B: 343.34}

	fmt.Println("Area:", iface1.Area(), "Perimeter:", iface1.Perimeter())

	//	iface1.Display()
}

func AreaAndPerimeter(iface IShape) {
	fmt.Println("Area:", iface.Area(), "Perimeter:", iface.Perimeter())
}

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
