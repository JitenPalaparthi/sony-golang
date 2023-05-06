package main

import (
	"fmt"
	"os"
)

func main() {
	fn, ln := new(string), new(string)
	*fn = ""
	*ln = "Corp"

	if err := StoreFullname(fn, ln); err != nil {
		println(err)
	}

	*fn = "Sony"
	*ln = "Corp"

	if err := StoreFullname(fn, ln); err != nil {
		println(err)
	}

	Div(10, 0)

	fmt.Println("Hello Main-1")

	func() {
		fmt.Println("Hello Main-2")
	}()

}
func StoreFullname(fn, ln *string) error {
	defer OnPanic()
	f, err := os.Create("names.txt")
	defer f.Close()
	defer println("File is closed")

	if err != nil {
		return err
	}

	if fn == nil || *fn == "" {
		//str = "firstname error"
		panic("Firstname cannot be nil or empty")
		//log.Fatal("Firstname cannot be nil or empty")
	}

	if ln == nil || *ln == "" {
		//str = "lastname error"
		panic("Lastname cannot be nil or empty")
	}

	f.WriteString("Full Name:" + *fn + " " + *ln)

	return nil
}

func Div(a, b int) int {
	defer OnPanic()
	return a / b // what is div by zero?
}

func OnPanic() {
	if rec := recover(); rec != nil { // that means there is a panic
		fmt.Println("Recovered", rec)
		// if strings.Contains(rec.(string), "firstname") {
		// 	fmt.Println("issue is with firstname ")
		// }
	}
}

// defer,panic,recover
