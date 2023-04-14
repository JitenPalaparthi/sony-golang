package main

import "fmt"

func main() {

	var iface any

	iface = 10
	iface = false
	iface = "Hello World"
	iface = 12.45678

	var f2 float32 = 12.45678

	iface = f2

	switch iface.(type) {
	case int:
		fmt.Println("Value of iface is", iface, "and type is int")

	case float64:
		fmt.Println("Value of iface is", iface, "and type is float64")

	case bool:
		fmt.Println("Value of iface is", iface, "and type is bool")

	case string:
		fmt.Println("Value of iface is", iface, "and type is string")

	default:
		fmt.Println("Value of iface is", iface, "and type is some other type")
	}

}
