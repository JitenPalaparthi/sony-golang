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

	mul := Cal(10, 30, func(i1, i2 int) int {
		return i1 * i2
	})

	fmt.Println("Mul", mul)

	di := Cal(10, 20, div)
	fmt.Println("Div", di)

	add1 := func(a int, b int) int {
		return a + b
	}

	var ptr *func(int, int) int = &add1
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

// func main() {
// 	      var pfunc func()
// 		   pfunc = Greet
// 		   pfunc()
// 		    var pfunc2 func(string)
// 			  pfunc2 = Greet2
// 			   pfunc2("LNR")  }
// 			    func Greet() {
// 					fmt.Println("Greet")  }
// 					 func Greet2(name string) {
// 						     fmt.Println(
