// channel
// do not communicate by sharing memory;instead share memory by communicating
// 1- There are two types of channels buffered and unbuffered.
// 2- To instantiate a channel use make built in function.
// 3- when the channel is declared by not instantiated then it is a nil channel
// 4- To send a value to a channel is ch <- 100
// 5- To receive a value from a channel is val := <- ch
// 6- The sender gorotuine is blocked until the receiver goroutine receives a value (For unbuffered channels)
// 7- The receiver goroutine is blocked until the sender sends a value to the channel (For unbuffered channels)
// 8- if sender or receiver is blocked then the runtime identifies the declare it as a deadlock.

package main

import (
	"fmt"
	"time"
)

func main() {

	var ch chan int     // declaration of the channel
	ch = make(chan int) // This int an unbuffered channel
	go func() {
		time.Sleep(time.Second * 5)
		ch <- 100 // sending a value to the channel
	}()
	num := <-ch // receive a value from the channel

	fmt.Println(num)

}
