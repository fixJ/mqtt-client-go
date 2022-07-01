package main

import (
	"mqtt-client-go/pkg/client"
)

func main() {
	consumer, err := client.NewConsumer("127.0.0.1", "1883", "jx", false, "", "")
	if err != nil {
		return
	}
	consumer.Subscribe("abc")
	consumer.Get()
}
