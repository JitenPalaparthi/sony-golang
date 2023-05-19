package main

import (
	"fmt"
	"runtime"

	"github.com/nats-io/nats.go"
)

// Connect to a server
func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	//for {
	nc.Subscribe("contact.create", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
		// with this data you can perform what ever the kind of operation you have to do..
	})

	runtime.Goexit()
	//}
}
