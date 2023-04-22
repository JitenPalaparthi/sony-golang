package main

import (
	"fmt"
	"runtime"
	"time"
)

// 1. Main is also a go routine
// 2. By default no goroutine waits for other go routine to complete its execution
// 3. The order of execution goroutines cannot be same always
// goroutines are small in size.
// goroutines are scheduled by go runtime
// goroutines are multiplexed against OS threads
// If a goroutine is blocked on a OS thread, go runtime efficiently execute other goroutines by
// shifting to another thread.

func main() {

	go fmt.Println("Hello World")

	Greet("Sony")

	fmt.Println("End of Main")

	fmt.Println(runtime.NumCPU())

	fmt.Println(runtime.NumGoroutine())
	// var str string
	// fmt.Scanln(&str)
	time.Sleep(time.Second * 1)
}

func Greet(name string) {
	go func() {
		for {
			fmt.Println("Hello" + name + "!")
			go main()

		}
	}()
}
