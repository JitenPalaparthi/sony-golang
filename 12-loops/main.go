package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ { // loop iterates until the condition is true
		fmt.Print(i, " ")
	}

	fmt.Println()
	i := 1
	for ; i <= 100; i++ { // loop iterates until the condition is true
		fmt.Print(i, " ")
	}

	fmt.Println()
	i = 1 // init
	for { // simlar to while loop
		if i > 100 { // condition // it breaks the loop when the condition is false
			break // break breaks the loop execution and takes the control out of the loop
		}
		fmt.Print(" ", i)
		i++ //post condition
	}
	fmt.Println("\neven numbers using continue")
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			continue // continue takes to the next iteration; that means it breaks only the ongoing iteration and takes to the next iteration
		}
		fmt.Print(i, " ")
	}

	fmt.Println("\nusing more than one inits")

	for i, j := 3, 5; i != j; i, j = i*4, j*3 {
		if i >= 1000 {
			break
		}
		fmt.Print("i:", i, "j:", j)

	}

	fmt.Println("\nnested loop")

outer:
	for i := 1; i <= 10; i++ {
		for j := 3; j <= 10; j++ {
			if i == j {
				break outer
			}
			fmt.Println("i->", i, "j->", j)
		}
	}
	fmt.Println("Above loops are done")
}
