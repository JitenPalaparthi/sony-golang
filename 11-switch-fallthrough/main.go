package main

import "fmt"

func main() {
	num := 24
	switch {
	case num%8 == 0:
		fmt.Println(num, "is divisible by 8")
		fallthrough
	case num%4 == 0:
		fmt.Println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		fmt.Println(num, "is divisible by 2")
		// fallthrough
		// case num%9 == 0:
		// 	fmt.Println(num, "is divisible by 9"// case num%9 == 0:
		// 	fmt.Println(num, "is divisible by 9"
	}
}

// do not use fallthrough when here are false negative statements
// use fallthrough when if the above case is satisfied and automatically the below case should also be satisfied use it
// think of scenarios in other programming switch cases as and where you dont use break statement(use fallthrough in go)
