// in go a type is not impllicitly converted to another type
// type casting to be used to converty from one type to another type
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var num1 int8 = 100 //1 byte

	//var num2 int = num1 //8 bytes on 64 bit arch

	var num2 int = int(num1)

	fmt.Println(num2)

	var f1 = 123.43 //float64

	var f2 float32 = float32(f1)

	fmt.Println(f2)

	// b1 := false

	// var num3 uint8 = uint8(b1)

	// fmt.Println(num3, b1)

	var str1 string = "Hello World"

	var num3 int = int(str1[0])

	fmt.Println(num3)

	var (
		num4  int     = 100
		f3    float32 = 102.3
		f4    float64 = 13.43
		num5  uint8   = 100
		sumi1 int     // num2+f3+f5+num5
		sumf1 float32 // num2+f3+f5+num5
		sumf2 float64 // num2+f3+f5+num5
	)

	sumi1 = num4 + int(f3) + int(f4) + int(num5)
	sumf1 = float32(num4) + f3 + float32(f4) + float32(num5)
	sumf2 = float64(num4) + float64(f3) + f4 + float64(num5)

	fmt.Println(sumf1, sumf2, sumi1)
	fmt.Printf("%0.2f,%0.2f %d", sumf1, sumf2, sumi1)

	var b1 byte = 123 // byte is nothing but uint8
	var b2 uint8 = b1

	fmt.Println(b1, b2)

	// var str2 string = "Sony"

	// var num6 int = int(str2)

	var str2 string = "10001"
	var str3 string = "Sony"

	var num6, num7 int

	num6, _ = strconv.Atoi(str2) // _ blank identifier
	num7, err := strconv.Atoi(str3)

	fmt.Println(num6, num7, err)

	str4 := strconv.Itoa(231312)
	fmt.Println(str4)

	var (
		num8 int     = 100
		f5   float64 = 123131.12321
		b3   bool    = false
	)

	str5 := fmt.Sprint(num8, f5, b3)

	fmt.Println("Value:", str5, "Type:", reflect.TypeOf(str5))
	// a, _, _ := 10, 11, 12
	// fmt.Println(a)

}
