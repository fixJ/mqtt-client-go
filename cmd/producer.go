package main

import (
	"fmt"
	"mqtt-client-go/pkg/client"
	"mqtt-client-go/pkg/message"
)

func main() {
	producer, err := client.NewProducer("127.0.0.1", "1883", "producer", false, "", "")
	if err != nil {
		return
	}
	for i := 1; i <= 10; i++ {
		pm := message.NewPublishMessage("abc", fmt.Sprintf("this message is NO.%d", i))
		producer.Publish(pm)
	}
	producer.Close()
}
