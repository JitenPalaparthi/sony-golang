package main

import (
	"fmt"
)

func main() {

	var (
		num1 int = 100

		num2 any = 200 // what is the type of any int

		num3 float32 = 300

		num4 uint8 = 10

		sum float64
	)

	// type assertion:  v.(t)
	sum = float64(num1) + float64(num2.(int)) + float64(num3) + float64(num4)

	fmt.Println(sum)

	var iface1 any = 10

	var num int = iface1.(int)

	//var float float64 = iface1.(float64) // type assertion is wrong as the underlining value is of type int

	//fmt.Println(num, float)

	fmt.Println(num)

	//var t1 reflect.Type = reflect.TypeOf(iface1)
	//var float float64 = float64(iface1.(t1))
	//fmt.Println(num, float)

	// var num8 float32
	// num8, ok := iface1.(int)
	// fmt.Println(num8, ok)
}
