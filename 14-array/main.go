package main

import (
	"fmt"
	"reflect"
)

// arrays in golang are fixed length elements of same type.The length of the array must be known to the compiler
// arrays are immutable. It does not mean that we cannot change value of array.
// It means we cannot change the length of an array once it is declared
// The type of an array contains its length as well

func main() {

	// var num int
	// const cst int
	var arr1 [2]int
	arr1[0] = 10
	arr1[1] = 20
	// arr1[2] = 30 not ok becase the length is 2

	arr1[0] = 100 // ok because values are mutable. only the length is not mutable

	arr2 := [3]int{10, 20, 30} // shorthand declaration

	arr3 := [...]int{10, 20, 30, 40} // shorthand declaration with eclipse symbol

	fmt.Println("Values of arrays", arr1, arr2, arr3)
	fmt.Println("Type of arrays", reflect.TypeOf(arr1), reflect.TypeOf(arr2), reflect.TypeOf(arr3))

	var arr4 [4]int
	arr4 = arr3

	// var arr5 [5]int
	// arr5 = arr4 //not ok becase both arrays are different as lengths are different

	fmt.Println("arr4", arr4)
	fmt.Println("Length of arr1", len(arr1))
	fmt.Println("Cap of arr1", cap(arr1)) // capacity does not make sense in arrays

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		fmt.Println("Index:", i, "Value:", v)
	}

	arr6 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("2D array", arr6)

	arr7 := [2][2][2]int{{{1, 2}, {3, 4}}, {{1, 2}, {3, 4}}}
	fmt.Println(arr7)

}
