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

	fmt.Println("Hello Main-1")

	func() {
		fmt.Println("Hello Main-2")
	}()

}
func StoreFullname(fn, ln *string) error {
	f, err := os.Create("names.txt")
	defer f.Close()
	defer println("File is closed")

	if err != nil {
		return err
	}

	if fn == nil || *fn == "" {
		panic("Firstname cannot be nil or empty")
		//log.Fatal("Firstname cannot be nil or empty")

	}

	if ln == nil || *ln == "" {
		panic("Lastname cannot be nil or empty")
	}

	f.WriteString("Full Name:" + *fn + " " + *ln)

	return nil
}
