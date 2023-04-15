package main

import "fmt"

func main() {
	myFunc := func(x int) int {
		return x * x
	}
	var myFuncPtr *func(int) int = &myFunc
	result := (*myFuncPtr)(5)
	fmt.Println(result)
}
