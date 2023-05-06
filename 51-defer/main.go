package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("This is end of main")

	// func() {
	// 	defer fmt.Println("call -1")
	// 	fmt.Println("call -2")
	// 	fmt.Println("call-3")
	// }()

	fmt.Println("call-4")
	fmt.Println("Start of main")

	f, err := os.Create("test.log")
	defer f.Close() // this is called 100%
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString("Hello Sony Corp. This is a simple example")

	a := 100

	b := new((int))
	*b = 200

	defer fmt.Println("\nValue of a and b with defer", a, *b)

	a += 100
	*b += 200

	fmt.Println("\nValue of a and b before defer", a, *b)

	// defer inside a loop

	str := "Sony Corp"

	for _, s := range str {
		defer fmt.Print(string(s))
	}
}

// defer : deferes the execution towards the end of the cfunction call
// defer: maintains its own callstack
// defer: all defer executions are stacked. That means the first defer is executed at last
