package main

import (
	"fmt"
)

func main() {

	var num1 int = 100

	var ptr1 *int = &num1

	fmt.Println(num1, *ptr1)

	var ptr2 *int

	if ptr2 != nil {
		*ptr2 = 200
	} else {
		fmt.Println("ptr2 is a nil pointer")
	}

	ptr3 := new(int) // new creates some memory and assign that address to the pointer

	*ptr3 = 300

	//c1 := add(10, 20)

	var a1, b1 int
	var ptr4, ptr5 *int = &a1, &b1

	c2 := add(ptr4, ptr5)

	fmt.Println(*c2)

	var a2, b2 = 10, 20

	c3 := add(&a2, &b2)

	fmt.Println(*c3)

	a3, b3 := new(int), new(int) // *a3 ,*b3 contains 0 ,0

	c4 := add(a3, b3) // dur to type inference/default values to the given memory 0 and 0 are the values

	fmt.Println(*c4)

	*a3, *b3 = 100, 200 // dereferencing

	c5 := add(a3, b3)

	fmt.Println(*c5)

	var a4 = 100

	increment(a4) // call by value and hence another copy of variable is created in the function stack frame and incremented in that

	fmt.Println(a4)

	incrementPtr(&a4) // call by ref and hence the same varible address passed so the increment happens at the value in the address;so a4 gets incremented

	fmt.Println(a4)

}

func add(a, b *int) *int {

	// var c int

	// c = *a + *b

	// return &c

	c := new(int)

	*c = *a + *b

	return c

}

func increment(num int) {
	num++
}

// func incrementPtr(num *int) error {
// 	if num != nil {
// 		*num++
// 		return nil
// 	}

// 	return errors.New("invalid pointer or nil pointer as an argument")
// }

func incrementPtr(num *int) {
	*num++
}
