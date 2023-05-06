package main

import "fmt"

type Message struct {
	Data   []byte
	Status int
}

func main() {
	ch := make(chan int, 2)
	ch <- 100
	ch <- 200
	fmt.Println(<-ch, <-ch)
	ch <- 300
	ch <- 400
	fmt.Println(<-ch, <-ch)

	strch := make(chan Message, 10)
	go func() {
		strch <- Message{Data: []byte("Sending to logger"), Status: 1}
		strch <- Message{Data: []byte("Sending to Email Server"), Status: 2}
		strch <- Message{Data: []byte("Sending to Messaging server"), Status: 3}
		strch <- Message{Data: []byte("Sending to XYZ server"), Status: 4}
		strch <- Message{Data: []byte("Sending to XYZ server"), Status: 5}
		strch <- Message{Data: []byte("Sending to logger"), Status: 6}
		strch <- Message{Data: []byte("Sending to Email Server"), Status: 7}
		strch <- Message{Data: []byte("Sending to Messaging server"), Status: 8}
		strch <- Message{Data: []byte("Sending to XYZ server"), Status: 9}
		strch <- Message{Data: []byte("Sending to XYZ server"), Status: 10}
		strch <- Message{Data: []byte("Sending to logger"), Status: 11}
		strch <- Message{Data: []byte("Sending to Email Server"), Status: 12}
		strch <- Message{Data: []byte("Sending to Messaging server"), Status: 13}
		strch <- Message{Data: []byte("Sending to XYZ server"), Status: 14}
		strch <- Message{Data: []byte("Sending to XYZ server"), Status: 15}
		strch <- Message{Data: []byte("Sending to logger"), Status: 16}
		strch <- Message{Data: []byte("Sending to Email Server"), Status: 17}
		strch <- Message{Data: []byte("Sending to Messaging server"), Status: 18}
		strch <- Message{Data: []byte("Sending to XYZ server"), Status: 19}
		strch <- Message{Data: []byte("Sending to XYZ server"), Status: 20}
		strch <- Message{Data: []byte("Sending to logger"), Status: 21}
		strch <- Message{Data: []byte("Sending to Email Server"), Status: 22}
		strch <- Message{Data: []byte("Sending to Messaging server"), Status: 23}
		strch <- Message{Data: []byte("Sending to XYZ server"), Status: 24}
		strch <- Message{Data: []byte("Sending to XYZ server"), Status: 25}
		close(strch)
	}()

	go func() {
		for v := range strch {
			fmt.Println("Received from -1", v)
		}
	}()
	go func() {
		for v := range strch {
			fmt.Println("Received from -3", v)
		}
	}()
	for v := range strch {
		fmt.Println("Received from -2", v)
	}

}
