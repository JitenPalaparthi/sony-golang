package main

import "fmt"

func main() {

	ch1 := Square()

	ch2 := Square()

	done := make(chan struct{}, 2)

	go func() {
		for v := range ch1 { // as long as the channel is not close the range loop tries/waits for values from channel
			fmt.Println("Receiving from chan1", (v))
		}
		done <- struct{}{}
	}()
	go func() {
		for v := range ch2 { // as long as the channel is not close the range loop tries/waits for values from channel
			fmt.Println("Receiving from chan2", (v))
		}
		done <- struct{}{}
	}()

	<-done
	<-done
}

func Square() chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i * i
		}
		close(ch)
	}()
	return ch
}
