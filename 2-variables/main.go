// variables

// numbers : int8,int16,int32,int64, uint4..uint64, int, float32,float64
// boolean: bool
// string: string

// this is how to write inline comment.

/* Multiline
comment*/
/// used documentation purpose

// in go when a variable is declared and not used is an error(not a warning)

// type inference: based on the type of a varialbe automatically a default value is given
// for all numbers the default value is 0
// for bool the default value is false
// for string the default value is empty string

package main

import (
	"fmt"
	"reflect"
)

type integer = int // aslias for int is integer

func main() {
	var num1 uint8 // 1byte 0-255
	var num2 int8  // -127 -- 127

	var num3 int = 324234 // assign value

	fmt.Println(num1, num2, num3)

	var num4 = 13 // automatically go creates a variable of the type which contains highest memory
	fmt.Println("Value:", num4, "type:", reflect.TypeOf(num4))

	var f1 float32 = 2323.223
	var f2 = 2334.343

	fmt.Println("Value of f1:", f1, "Type of f1:", reflect.TypeOf(f1))

	fmt.Printf("Value of f2:%f Type of f2 %T\n", f2, f2)

	var (
		num5, f3, bl1, str1 = 32434, 4343.34, true, "Hello"
	)

	fmt.Println(num5, f3, bl1, str1)

	// shorthand declaration

	num6 := 100 // based on the assigned value , variable num6 of type int is created

	f4 := 21312.34 // based on the assigned value, variable f4 of type float64 is created
	fmt.Println(num6, f4)

	// simple swap of two varialbles

	var a1, b1 int = 10, 20
	var t int = a1
	a1 = b1
	b1 = t

	var a2, b2 int = 10, 20
	a2, b2 = b2, a2 // simple swap

	var a3, b3, c3 = 10, 20, 30
	a3, b3, c3 = c3, a3, b3

	// +, - , *, / ,%, ++,+=

	fmt.Println((a3 + b3) / c3)

	fmt.Println(a1 == b1)

	var num7 integer = 1232131

	fmt.Println("Value:", num7, "Type:", reflect.TypeOf(num7))

	// rune and byte
	// rune is alias for int32
	// byte is alias for uint8

	str2 := "Hello World"

	char1 := 'H' // rune

	fmt.Println("str2", str2)

	fmt.Println("value", char1, "type", reflect.TypeOf(char1))

	var num8 int32 = 'H' // internally it gives unicode value of the char
	fmt.Println("value", num8, "type", reflect.TypeOf(num8))

	// complex numbers contain real and imaginary part
	// there are two types of complex numbers : complex64 and complex128
	c1 := complex(12, 12.34)

	fmt.Println("value", c1, "type", reflect.TypeOf(c1))

	var r1, i1 float32 = 12, 12.34

	c2 := complex(r1, i1)

	fmt.Println("value", c2, "type", reflect.TypeOf(c2))

	c4 := 12 + 12.34i // another way of declaring complex number
	fmt.Println("value", c4, "type", reflect.TypeOf(c4))

	fmt.Println("Complex add", c1+c4)
	fmt.Println("Complex sub", c1-c4)
	fmt.Println("Complex mul", c1*c4)

	// empty interface

	// there is an alias for empty interface that is "any"

	//var iface1 interface{} = 100
	var iface1 any = 100

	fmt.Println("Value:", iface1, "Type:", reflect.TypeOf(iface1))

	iface1 = false
	fmt.Println("Value:", iface1, "Type:", reflect.TypeOf(iface1))

	iface1 = "Hello Sony"
	fmt.Println("Value:", iface1, "Type:", reflect.TypeOf(iface1))

	iface1 = c1 //c1 is complex128
	fmt.Println("Value:", iface1, "Type:", reflect.TypeOf(iface1))

}
