package main

import (
	"fmt"
	"mqtt-client-go/pkg/message"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1883")
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	cm := message.NewConnectMessage("jx", false, "", "")
	conn.Write(cm.Build())
	sm := message.NewSubscribeMessage("abc")
	conn.Write(sm.Build())
	for {
		buffer := make([]byte, 1024*10)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		pms, err := message.BytesToPublishMessage(buffer[:n])
		for _, pm := range pms {
			prm := message.NewPublishReceiveMessage(pm.GetPacketID())
			conn.Write(prm.Build())
			fmt.Printf("receive message, topic is %s, message is %s\n", pm.Topic, pm.Data)
		}
	}
}
