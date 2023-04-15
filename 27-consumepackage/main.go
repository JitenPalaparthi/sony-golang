package main

import (
	"fmt"

	shape "github.com/JitenPalaparthi/shapes/shape"
	"github.com/JitenPalaparthi/shapes/shape/square"
)

func main() {

	a1 := shape.Area(10.123, 12.343)
	fmt.Printf("Area of Rect:%0.2f\n", a1)

	p1 := shape.Perimeter(10.123, 12.35)
	fmt.Printf("Perimeter of Rect:%0.2f\n", p1)

	a2, p2 := square.Area(25.12), square.Perimeter(25.12)

	fmt.Println("Area of Square:", a2, "Perimeter of Square:", p2)

}
