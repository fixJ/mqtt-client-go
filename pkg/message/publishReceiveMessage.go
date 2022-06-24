package message

type PublishReceiveMessage struct {
	PacketID       int16
	fixHeader      []byte
	variableHeader []byte
	payload        []byte
}

func NewPublishReceiveMessage(packetID int16) *PublishReceiveMessage {
	return &PublishReceiveMessage{
		PacketID: packetID,
	}
}

func (m *PublishReceiveMessage) setFixHeader() {
	m.fixHeader = append(m.fixHeader, byte(80))
}

func (m *PublishReceiveMessage) setVariableHeader() {
	m.variableHeader = append(m.variableHeader, int16ToBytes(m.PacketID)...)
	m.variableHeader = append(m.variableHeader, byte(0))
}

func (m *PublishReceiveMessage) Build() []byte {
	var msg []byte
	m.setFixHeader()
	m.setVariableHeader()
	all := len(m.variableHeader) + len(m.payload)
	m.fixHeader = append(m.fixHeader, byte(all))
	msg = append(msg, m.fixHeader...)
	msg = append(msg, m.variableHeader...)
	msg = append(msg, m.payload...)
	return msg
}
