package client

import (
	"fmt"
	"mqtt-client-go/pkg/message"
	"net"
)

type Consumer struct {
	ClientID string
	NeedAuth bool
	Username string
	Password string
	conn     net.Conn
}

func NewConsumer(ip string, port string, clientID string, needAuth bool, username string, password string) (*Consumer, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	c := &Consumer{
		ClientID: clientID,
		NeedAuth: needAuth,
		Username: username,
		Password: password,
		conn:     conn,
	}
	c.connect()
	return c, nil
}

func (c *Consumer) connect() {
	cm := message.NewConnectMessage(c.ClientID, c.NeedAuth, c.Username, c.Password)
	c.conn.Write(cm.Build())
}

func (c *Consumer) Subscribe(topic string) {
	sm := message.NewSubscribeMessage(topic)
	c.conn.Write(sm.Build())
}

func (c *Consumer) Get() {
	for {
		buffer := make([]byte, 1024*10)
		n, err := c.conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		pms, err := message.BytesToPublishMessage(buffer[:n])
		for _, pm := range pms {
			prm := message.NewPublishReceiveMessage(pm.GetPacketID())
			c.conn.Write(prm.Build())
			fmt.Printf("receive message, topic is %s, message is %s\n", pm.Topic, pm.Data)
		}
	}
}

func (c *Consumer) Close() {
	c.conn.Close()
}
