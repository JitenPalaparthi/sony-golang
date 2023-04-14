package main

import (
	"fmt"
	"os"
)

func main() {

	// if age := 10; age >= 18 {
	// 	fmt.Println("eligible for vote", age)
	// 	return
	// } else {
	// 	fmt.Println("not eligble for vote", age)
	// }

	// if _, err := os.Open("abcd.txt"); err != nil {

	// } else {

	// }

	_, err := os.Open("abcd.txt")
	OnErr(err)

	// a, b := 10, 20

	// t := a
	// a = b
	// b = t

	// a, b = b, a
	var age uint8 = 18
	var country string = "India"
	if age >= 18 && (country == "India" || country == "Bangladesh") {
		// true && (true || false) -> true && true -> true
		fmt.Println("Eligible for vote because age is", age, "and country is", country)
	} else if age >= 16 && country == "Pakisthan" {
		fmt.Println("Eligible for vote because age is", age, "and country is", country)
	} else {
		fmt.Println("not eligible for vote")
	}
	fmt.Println("Can I be executed")

	if age >= 18 {
		if country == "India" {
			fmt.Println("Eligible for vote because age is", age, "and country is", country)
		} else if country == "Bangladesh" {
			fmt.Println("Eligible for vote because age is", age, "and country is", country)
		}
	}
}

func OnErr(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
