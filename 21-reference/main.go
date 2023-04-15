package main

import (
	"fmt"
	"math/rand"
)

func main() {

	slice1 := make([]int, 1, 5) // Elm > ptr ,len, cap &slice1 , len:1, cap:5
	slice1[0] = 100
	fmt.Println(slice1, len(slice1), cap(slice1))
	fmt.Println("Address of slice1", &slice1[0])
	slice1 = AddRandomElem(slice1, 4)
	fmt.Println("Address of slice1", &slice1[0])
	fmt.Println(slice1, len(slice1), cap(slice1))

	slice2 := []int{10, 11, 12, 13, 14} // Elm ptr:= &slice2[0], len:5, cap:5
	fmt.Println(slice2)

	UpdateRandomElem(slice2)
	fmt.Println(slice2)
}

func AddRandomElem(slice []int, count int) []int {
	for i := 0; i < count; i++ {
		slice = append(slice, rand.Intn(1000000))
	}
	//fmt.Println(slice)[]int
	return slice
}

func UpdateRandomElem(slice []int) {
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(1000000)
	}
}
