package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {
	mu := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	wg.Add(2000)

	for i := 1; i <= 1000; i++ {
		go Increment(wg, mu)
		go Decrement(wg, mu)
	}

	//time.Sleep(time.Second + 1)
	wg.Wait()
	fmt.Println(counter)
}

func Increment(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	counter++
	mu.Unlock()
	wg.Done()
}

func Decrement(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	counter--
	mu.Unlock()
	wg.Done()
}
