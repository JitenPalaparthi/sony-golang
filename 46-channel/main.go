// package main

// import (
// 	"fmt"
// )

// var counter int

// func main() {
// 	// var ch chan int
// 	// declaration of the channel
// 	ch := make(chan int) // int an unbuffered channel
// 	done := make(chan struct{})
// 	done2 := make(chan struct{})
// 	go func() {
// 		for i := 1; i <= 100; i++ {
// 			go Increment(&ch, &done)
// 			go Decrement(&ch, &done)
// 		}
// 		done2 <- struct{}{}
// 	}()

// 	<-done2
// 	fmt.Println(counter)
// }
// func Increment(ch *chan int, done *chan struct{}) {
// 	counter++
// 	*done <- struct{}{}
// 	//<-*done
// }
// func Decrement(ch *chan int, done *chan struct{}) {
// 	<-*done
// 	counter = counter - 1
// 	//*done <- struct{}{}
// }
//package main

// import (
// 	"fmt"
// )

// var counter int

// func main() {
// 	// var ch chan int
// 	// declaration of the channel
// 	ch := make(chan int) // int an unbuffered channel
// 	done := make(chan struct{})
// 	done2 := make(chan struct{})
// 	go func() {
// 		for i := 1; i <= 100; i++ {
// 			go Increment(&ch, &done)
// 			go Decrement(&ch, &done)
// 		}
// 		done2 <- struct{}{}
// 	}()

// 	<-done2
// 	fmt.Println(counter)
// }
// func Increment(ch *chan int, done *chan struct{}) {
// 	counter++
// 	*done <- struct{}{}
// 	//<-*done
// }
// func Decrement(ch *chan int, done *chan struct{}) {
// 	<-*done
// 	counter = counter - 1
// 	//*done <- struct{}{}
// }

// instead of using channel using mutex (mutual exclusion) -provide atomocity to operation.until that operation finished that global variable is protected.
package main

import (
	"fmt"
	"time"
)

var counter int

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	//done := make(chan int)
	go func() {
		ch1 <- struct{}{}
	}()

	go func() {
		for i := 1; i <= 1000; i++ {
			go func(a int) {
				<-ch1
				counter++
				fmt.Println("Incre->", a, counter)
				ch2 <- struct{}{}

			}(i)
			go func(a int) {
				<-ch2
				counter--
				fmt.Println("Decr->", a, counter)
				ch1 <- struct{}{}
			}(i)
			//if i == 1000 {
			//	done <- struct{}{}
			//}
		}
		//done <- 10
		//close(done)
		//done <- struct{}{}
		// only sender can close the channel
		// when the channel is closed it sends a notification so that the receiver can receive a value
	}()

	// v, ok := <-done

	// fmt.Println(v, ok)

	time.Sleep(time.Second * 5)
	fmt.Println(counter)
}
