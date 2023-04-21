package main

import "fmt"

func main() {

	c1 := ColorCode{10, "Red", true}
	c2 := ColorCode{int: 11, string: "Blue", bool: true}

	fmt.Println(c1, c2)

	c3 := struct {
		string
		int
	}{
		string: "Jiten",
		int:    10001,
	}
	fmt.Println(c3)

	e1 := Empty{}
	e1.Who()
}

// fields without names but only with types
type ColorCode struct {
	int
	string
	bool
}

type Empty struct{}

func (e Empty) Who() {
	fmt.Println("I am empty struct")
}
