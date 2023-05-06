package main

import (
	"fmt"
	"time"
)

func main() {

	server1 := GetServer("Server-1")
	server2 := GetServer("Server-2")
	server3 := GetServer("Server-3")
	server4 := GetServer("Server-4")
	timeout := time.After(time.Millisecond * 5)

	select {
	case str := <-server1:
		fmt.Println(str)
	case str := <-server2:
		fmt.Println(str)
	case str := <-server3:
		fmt.Println(str)
	case str := <-server4:
		fmt.Println(str)
	case <-timeout:
		fmt.Println("timeout")
	}

}

func GetServer(str string) chan string {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Millisecond * 5)
		ch <- str
		close(ch)
	}()
	return ch

}
