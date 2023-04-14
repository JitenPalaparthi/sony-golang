package main

import "fmt"

// slice is a growable array.
// slice is a reference type
// you can check whether slice is instantiated or not by checking whether it is nil or not
// nil can be checked on slice,map,chan, any pointer type and also interface
// make builtin function is used to instantiate slice, map and chan
// when a slice is instantiated with a given length, type inference works

func main() {

	//var arr1 [5]int
	var slice1 []int // define a slice

	if slice1 == nil {
		fmt.Println("slice1 is not instantiated")
	}

	// to instantiate the slice
	slice1 = make([]int, 5, 10) // 40 bytes to be reserved and give handle to slice
	// slice header contains, pointer of first element, length and capacity

	fmt.Println(slice1)

	slice1[0] = 10
	slice1[1] = 20
	slice1[2] = 30
	slice1[3] = 40
	slice1[4] = 50
	//slice1[5] = 60
	fmt.Println(slice1)

	// by default for slice1 the length of the slice is 5 and the cap is 10
	slice1 = append(slice1, 60)
	slice1 = append(slice1, 70)
	slice1 = append(slice1, 80)
	slice1 = append(slice1, 90)
	slice1 = append(slice1, 100)

	fmt.Println(slice1, "Address:", &slice1[0], "Length:", len(slice1), "Cap:", cap(slice1))
	slice1 = append(slice1, 110) // the initial cap is 10 but there are 11 elements
	fmt.Println(slice1, "Address:", &slice1[0], "Length:", len(slice1), "Cap:", cap(slice1))
	slice1 = append(slice1, 120)
	fmt.Println(slice1, "Address:", &slice1[0], "Length:", len(slice1), "Cap:", cap(slice1))
	slice1 = append(slice1, 130, 140, 150, 160, 170, 180, 190, 200, 210)
	fmt.Println(slice1, "Address:", &slice1[0], "Length:", len(slice1), "Cap:", cap(slice1))

}
