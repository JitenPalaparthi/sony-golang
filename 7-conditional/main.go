package main

import (
	"fmt"
)

func main() {

	var age uint8 = 16
	var country string = "Pakisthan"

	if age >= 18 && (country == "India" || country == "Bangladesh") {
		// true && (true || false) -> true && true -> true
		fmt.Println("Eligible for vote because age is", age, "and country is", country)
	} else if age >= 16 {
		fmt.Println("Eligible for vote because age is", age, "and country is", country)
	} else if age >= 16 && country == "Pakisthan" {
		fmt.Println("Eligible for vote because age is", age, "and country is", country)
	} else {
		fmt.Println("not eligible for vote")
	}

	// ok := true

	// if !ok {

	// }

	//fmt.Println("Not eligible for vote")
	// os.Exit(0) // it  successfully executed
	// os.Exit(1) // if it is non zero then it failed to execute.i.e there is an error
}

// > < >= <= && || != & |
// if else else if switch
