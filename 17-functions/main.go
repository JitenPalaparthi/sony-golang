package main

import "fmt"

func main() {
	Greet()
	a1 := AreaOfRect(10, 12.5)
	p1 := PerimeterOfRect(10, 12.5)
	fmt.Println("Area:", a1, "Perimeter:", p1)

	a2, p2 := AreaAndPerimeter(10, 12.5)
	fmt.Println("Area:", a2, "Perimeter:", p2)

	a3, _ := AreaAndPerimeter(10.7, 12.5)
	fmt.Println("Area:", a3)
	_, p3 := AreaAndPerimeter(10.7, 12.5)
	fmt.Println("Perimeter:", p3)

	arr1 := [...]int{1, 2, 3}

	fmt.Println(SumOfArr(arr1))

	arr2 := [4]int{1, 2, 3, 4}
	//fmt.Println(SumOfArr(arr2))

	fmt.Println("SumOf", SumOf(arr1[:]))
	fmt.Println("SumOf", SumOf(arr2[:]))

	fmt.Println("Sum:", Sum())
	fmt.Println("Sum:", Sum(1))
	fmt.Println("Sum:", Sum(1, 2, 3, 4, 5, 6))
	fmt.Println("Sum:", Sum(1, 234, 234, 324, 24, 234, 324, 234, 23423, 423, 423, 423, 423, 4234342342424, 234234234))
	fmt.Println("Sum:", Sum(arr1[:]...)) // slice as an variadic argument to the function
	fmt.Println("Sum:", Sum(arr2[:]...)) // slice as an variadic argument to the function
}

func Greet() {
	fmt.Println("hello World")
}

func AreaOfRect(l, b float32) float32 {
	return l * b
}

func PerimeterOfRect(l float32, b float32) (f float32) {
	f = 2 * (l + b)
	return f
}

func AreaAndPerimeter(l, b float32) (a float32, p float32) {
	a, p = l*b, 2*(l+b)
	return a, p
}

func SumOfArr(arr [3]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumOf(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// here nums is a variadic parameter
// variadic parameter must be the last parameter in the function
// from 0-any number of values(as long as the stack supports) can be passed as an argument to the function
func Sum(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// func SumM(nums ...int,msg string) int { Not OK as variadic param must be the last param
func SumM(msg string, nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
