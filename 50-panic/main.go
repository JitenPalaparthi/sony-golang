// general errors
// runtime exception
// deadlocks
// panic
// panic by runtime
// panic by user (userdefined panic)
// fatals

// userdefined errors

// ex
// stackoverflow?( Oen of the exceptions types)
// panic: runtime error: index out of range [0] with length 0
// panic: runtime error: integer divide by zero
// fatal error: all goroutines are asleep - deadlock!

// by default panic panics the entire application. That means the application is caught with panic and it is a runtime error.Hence application is crashed

// error,panic,fatal

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	// fmt.Println("checking error, panic,deadlock")
	// v := 0
	// fmt.Println(100 / v)

	// ch := make(chan int)

	// ch <- 100

	// var slice1 []int // slice is declared but not initialised

	// slice1 = make([]int, 1)

	// fmt.Println("slice declared")

	// slice1[0] = 100

	// fmt.Println(slice1)
	// <-ch

	var fn, ln string = "Sony", " Corp"

	fun := FullnameF(nil, &ln)

	fun = Fullname(&fn, &ln)

	fmt.Println(*fun)

	fun, err := FullnameE(nil, &ln)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*fun)

}

func Fullname(fn, ln *string) *string {
	if fn == nil || *fn == "" {
		panic("first name cannot be nil or empty")
	}
	if ln == nil || *ln == "" {
		panic("last name cannot be nil or empty")
	}
	str := *fn + *ln
	return &str
}

var EMPTY_FIRST_NAME = errors.New("empty first name")

func FullnameE(fn, ln *string) (*string, error) {
	if fn == nil {
		return nil, errors.New("first name cannot be nil or empty")
	}

	if *fn == "" {
		return nil, EMPTY_FIRST_NAME
	}

	if ln == nil || *ln == "" {
		//return nil, errors.New("last name cannot be nil or empty")
		return nil, fmt.Errorf("last name cannot be nil or empty.Type of the variable %T", ln)
	}
	str := *fn + *ln
	return &str, nil
}

func FullnameF(fn, ln *string) *string {
	if fn == nil || *fn == "" {
		fmt.Println("Fatal:first name cannot be nil or empty")
		os.Exit(1)
		//log.Fatal("first name cannot be nil or empty")
	}
	if ln == nil || *ln == "" {
		fmt.Println("Fatal: last name cannot be nil or empty")
		os.Exit(1)
	}
	str := *fn + *ln
	return &str
}
