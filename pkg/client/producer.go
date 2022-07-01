package client

import (
	"fmt"
	"mqtt-client-go/pkg/message"
	"net"
)

type Producer struct {
	ClientID string
	NeedAuth bool
	Username string
	Password string
	conn     net.Conn
}

func NewProducer(ip string, port string, clientID string, needAuth bool, username string, password string) (*Producer, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	c := &Producer{
		ClientID: clientID,
		NeedAuth: needAuth,
		Username: username,
		Password: password,
		conn:     conn,
	}
	c.connect()
	return c, nil
}

func (p *Producer) connect() {
	cm := message.NewConnectMessage(p.ClientID, p.NeedAuth, p.Username, p.Password)
	p.conn.Write(cm.Build())
}

func (p *Producer) Publish(msg *message.PublishMessage) {
	p.conn.Write(msg.Build())
}

func (p *Producer) Close() {
	p.conn.Close()
}
