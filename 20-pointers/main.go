package main

import "fmt"

var num uint8 // global variable

func main() {

	var num1 uint8 = 100
	var ptr1 *uint8 // this is a pointer by default it is nil
	ptr1 = &num1    // & address of
	fmt.Println(num1, *ptr1)

	fmt.Println(num1, *ptr1, *(&num1))

	var ptr4 **uint8
	ptr4 = &ptr1

	var ptr5 ***uint8 = &ptr4

	fmt.Println(ptr4, ptr5)

	fmt.Println(***ptr5)

	changeVal(num)
	changeVal(num1)
	fmt.Println("Global num", num)
	fmt.Println("Local num", num1)

	changeValPtr(&num)
	changeValPtr(&num1)
	fmt.Println("Global num", num)
	fmt.Println("Local num", num1)

	// dereferencing an uninitilised pointer
	// var ptr2 *int
	// *ptr2 = 100

	// dangling pointer or nil pointer dereference
	// var ptr2 *uint8 = &num1
	// ptr2 = nil
	// *ptr2 = 101

	// Memory leaks

	//n1 := Square(100)
	// for i := 1; i <= 10000000; i++ {
	// 	ptr3 := Squareptr(100)
	// 	fmt.Println("Address out side of function", ptr3)
	// }

	var ptr6 *func() = &Greet

	(*ptr6)()

}

func Greet() {
	fmt.Println("Hello World")
}

func changeVal(n uint8) {
	n = n + 1
	//fmt.Println(n)
}

func changeValPtr(n *uint8) {
	*n = *n + 1
}

// in c or c++ this could cause a memory leak
// creating a pointer inside a func and return that pointer outside of that func
func Squareptr(n int) *int {
	n = n * n
	var nptr *int = &n
	fmt.Println("Address inside of function", nptr)
	return nptr
}

func Square(n int) int {
	return n * n
	// var nptr *int = &n
	// return *nptr
	//return n
}
