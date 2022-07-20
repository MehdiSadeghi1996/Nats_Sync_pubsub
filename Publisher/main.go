package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

var subject = "Customer"

func main() {

	//wait := make(chan bool)
	nc, err := nats.Connect("nats://10.187.160.107:5009")

	if err != nil {
		fmt.Println(err)
	}

	msg, err := nc.Request(subject, []byte("help me"), 10*time.Millisecond)
	fmt.Println(string(msg.Data))

	//<-wait
}
