package main

import "fmt"

func main() {

	fmt.Println("1-100 prime numbers")
	fmt.Print(1, " ", 2, " ", 3, " ")
	num := 4
loop:
	var count uint8 = 0
	for i := 2; i <= num-1; i++ {
		if num%i == 0 {
			count++
			break
		}
	}
	if count == 1 {
		//fmt.Println(num, "is not primer number")
	} else {
		fmt.Print(num, " ")
	}
	num++
	if num <= 100 {
		goto loop
	}

	fmt.Println("\n1-100 numbers")
	i := 1
loop1:
	fmt.Print(i, " ")
	i++
	if i <= 100 {
		goto loop1
	}
}
