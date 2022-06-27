package message

import (
	"math/rand"
)

type SubscribeMessage struct {
	Topic          string
	fixHeader      []byte
	variableHeader []byte
	payload        []byte
}

func NewSubscribeMessage(topic string) *SubscribeMessage {
	return &SubscribeMessage{
		Topic: topic,
	}
}

func (m *SubscribeMessage) setFixHeader() {
	m.fixHeader = append(m.fixHeader, byte(130))
}

func (m *SubscribeMessage) setVariableHeader() {
	packetID := int16(rand.Int())
	m.variableHeader = append(m.variableHeader, int16ToBytes(packetID)...)
	m.variableHeader = append(m.variableHeader, byte(0))
}

func (m *SubscribeMessage) setPayload() {
	var length int16
	topicBytes := []byte(m.Topic)
	length = int16(len(topicBytes))
	m.payload = append(m.payload, int16ToBytes(length)...)
	m.payload = append(m.payload, topicBytes...)
	m.payload = append(m.payload, byte(2))
}

func (m *SubscribeMessage) Build() []byte {
	var msg []byte
	m.setFixHeader()
	m.setVariableHeader()
	m.setPayload()
	all := len(m.variableHeader) + len(m.payload)
	m.fixHeader = append(m.fixHeader, byte(all))
	msg = append(msg, m.fixHeader...)
	msg = append(msg, m.variableHeader...)
	msg = append(msg, m.payload...)
	return msg
}
