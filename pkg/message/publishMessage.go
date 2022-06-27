package message

import (
	"math/rand"
)

type PublishMessage struct {
	Topic          string
	Data           string
	packetID       int16
	fixHeader      []byte
	variableHeader []byte
	payload        []byte
}

func NewPublishMessage(topic, data string) *PublishMessage {
	return &PublishMessage{
		Topic: topic,
		Data:  data,
	}
}

func (m *PublishMessage) setFixHeader() {
	m.fixHeader = append(m.fixHeader, byte(60))
}

func (m *PublishMessage) setVariableHeader() {
	var length int16
	topicBytes := []byte(m.Topic)
	length = int16(len(topicBytes))
	packetID := rand.Int()
	m.SetPacketID(int16(packetID))
	m.variableHeader = append(m.variableHeader, int16ToBytes(length)...)
	m.variableHeader = append(m.variableHeader, topicBytes...)
	m.variableHeader = append(m.variableHeader, int16ToBytes(m.GetPacketID())...)
	m.variableHeader = append(m.variableHeader, byte(0))
}

func (m *PublishMessage) setPayload() {
	m.payload = append(m.payload, []byte(m.Data)...)
}

func (m *PublishMessage) GetPacketID() int16 {
	return m.packetID
}

func (m *PublishMessage) SetPacketID(pid int16) {
	m.packetID = pid
}

func (m *PublishMessage) Build() []byte {
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
