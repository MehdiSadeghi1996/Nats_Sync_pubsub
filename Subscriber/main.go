package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

var subject = "Customer"

func main() {

	wait := make(chan bool)

	nc, err := nats.Connect("nats://10.187.160.107:5009")
	if err != nil {
		fmt.Println(err)
	}

	nc.Subscribe(subject, func(m *nats.Msg) {
		fmt.Println(string(m.Data))
		nc.Publish(m.Reply, []byte("I can help!"))
	})

	fmt.Println("Subscribed to", subject)

	<-wait
}
