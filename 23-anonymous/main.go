package main

import "fmt"

func main() {

	func() {
		fmt.Println("Hello WOrld")
	}()

	f1 := func() {
		fmt.Println("Hello World")
	}

	var f2 func() = func() {
		fmt.Println("Hello World")
	}

	f1()
	f2()

	f3 := func(a, b int) int {
		return a + b
	}
	c := f3(10, 20)
	fmt.Println("addition", c)

	c1 := func(a, b int) int {
		return a + b
	}(10, 20)
	fmt.Println(c1)

	add := Cal(10, 30, func(i1, i2 int) int {
		return i1 + i2
	})

	fmt.Println("Add", add)

	sub := Cal(10, 30, func(i1, i2 int) int {
		return i1 - i2
	})

	fmt.Println("Sub", sub)

	mul := Cal(10, 30, func(i1, i2 int) int { // function pointer as the last parameter is a function and passed as an argument
		return i1 * i2
	})

	fmt.Println("Mul", mul)

	di := Cal(10, 20, div)
	fmt.Println("Div", di)

	add1 := func(a int, b int) int {
		return a + b
	}

	var ptr *func(int, int) int = &add1 // pointers as functions
	a4 := (*ptr)(10, 20)

	fmt.Println(a4)

	d1 := div

	ptr = &d1

	fmt.Println((*ptr)(10, 20))

}

func div(a, b int) int {
	return a / b
}

// clousure
func Cal(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func sort(arr []int, f func([]int) []int) []int {
	return f(arr)
}

// task

//cal(a,b any, func(any,any)float64)(float64,error)

// if a and b are int , for addition it should be a+b return type if float64,nil
// if a and b are float32 , for addition it should be a+b return type if float64,nil
// if a and b are uint8-uint64 and int8-int64 or float32 or float64 it should return a+b -> float64 ,nil
// if a or b are other than uint8-uint64 and int8-int64 or float32 or float64 return a+b -> 0, invalid type provided
