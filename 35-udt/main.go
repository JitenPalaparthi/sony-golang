package main

import (
	"fmt"
	"reflect"
)

type Employee struct{}

type myInt int

type Integer = int // alias

func main() {

	var mi1 myInt = 345

	str1 := mi1.ToString()

	fmt.Println("Value:", str1, "Type:", reflect.TypeOf(str1))

	var i1 int = int(mi1)

	fmt.Println(i1)

	i2 := 23456

	str2 := myInt(i2).ToString()

	fmt.Println("Value:", str2, "Type:", reflect.TypeOf(str2))

	var i3 Integer = i1

	fmt.Println(i3)

	//var r1 rune

}

func (mi myInt) ToString() string {
	return fmt.Sprint(mi)
}
