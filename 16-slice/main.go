package main

import "fmt"

func main() {

	slice1 := []int{10, 20, 30, 40}
	fmt.Println(slice1, "Length:", len(slice1), "Cap:", cap(slice1))
	slice1 = append(slice1, 50)
	fmt.Println(slice1, "Length:", len(slice1), "Cap:", cap(slice1))

	slice2 := []int{} // --> make([]int,0) declare and instantiate
	if slice2 == nil {
		fmt.Println("nil slice")
	}
	fmt.Println(slice2, "Length:", len(slice2), "Cap:", cap(slice2))

	var slice3 []int // declaration
	if slice3 == nil {
		fmt.Println("nil slice")
	}

	fmt.Println(slice3, "Length:", len(slice3), "Cap:", cap(slice3))

	slice3 = make([]int, 0)

	slice4 := make([]int, 5) // shorthand for creating and instantiating a slice
	fmt.Println(slice4, "Length:", len(slice4), "Cap:", cap(slice4))

	arr1 := [5]int{10, 11, 12, 13, 14}

	slice5 := arr1[:] // an array can be converted to slice with this notation [:]
	fmt.Println("address of arr1:", &arr1[0], "address of slice5:", &slice5[0])
	slice5[0] = 100
	arr1[1] = 110

	fmt.Println("arr1", arr1)
	fmt.Println("slice5", slice5)

	//slice5 = append(slice5, 15)
	fmt.Println("address of arr1:", &arr1[0], "address of slice5:", &slice5[0])
	arr1[1] = 1110
	slice5[0] = 1000
	fmt.Println("arr1", arr1)
	fmt.Println("slice5", slice5)

	slice6 := slice5[:2]
	slice7 := slice5[2:5]
	slice8 := slice5[3:]
	slice9 := slice5[:]
	slice10 := slice5
	slice5 = append(slice5, 15)
	fmt.Println(slice5, slice6, slice7, slice8, slice9, slice10)
	fmt.Println("address of slice5", &slice5[0])
	fmt.Println("address of slice6", &slice6[0])
	fmt.Println("address of slice9", &slice9[0])
	fmt.Println("address of slice10", &slice10[0])

}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// slice := []int{1, 2, 3, 4}
// 	slice := make([]int, 4, 5)
// 	fmt.Println("slice = ", slice, len(slice), cap(slice))
// 	slice = append(slice, 2, 3, 8, 8, 9, 9, 10)
// 	fmt.Println("slice = ", slice, len(slice), cap(slice))
// 	slice = append(slice, 2)
// 	slice = append(slice, 3)
// 	fmt.Println("slice = ", slice, len(slice), cap(slice))
// }

// // 1-
