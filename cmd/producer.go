package main

import (
	"fmt"
	"mqtt-client-go/pkg/message"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1883")
	if err != nil {
		fmt.Println(err)
		return
	}
	cm := message.NewConnectMessage("producer", false, "", "")
	conn.Write(cm.Build())
	for i := 1; i <= 10; i++ {
		pm := message.NewPublishMessage("abc", fmt.Sprintf("this message is NO.%d", i))
		conn.Write(pm.Build())
	}
	conn.Close()
}
