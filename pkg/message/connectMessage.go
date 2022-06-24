package message

type ConnectMessage struct {
	ClientID       string
	NeedAuth       bool
	Username       string
	Password       string
	fixHeader      []byte
	variableHeader []byte
	payload        []byte
}

func NewConnectMessage(clientID string, needAuth bool, username string, password string) *ConnectMessage {
	return &ConnectMessage{
		ClientID: clientID,
		NeedAuth: needAuth,
		Username: username,
		Password: password,
	}
}

func (m *ConnectMessage) setFixHeader() {
	m.fixHeader = append(m.fixHeader, byte(16))
}

func (m *ConnectMessage) setVariableHeader() {
	m.variableHeader = append(m.variableHeader, byte(0), byte(4), byte(77), byte(81), byte(84), byte(84))
	// 设置version 5.0
	m.variableHeader = append(m.variableHeader, byte(5))
	//设置connect flags
	cleanStart := 1 << 1
	willFlag := 0 << 2
	willQOS := 0<<3 | 0<<4
	willRetain := 0 << 5
	usernameFlag := 0 << 7
	passwordFlag := 0 << 6
	if m.NeedAuth {
		usernameFlag = 1 << 7
		passwordFlag = 1 << 6
	}
	flags := cleanStart | willFlag | willQOS | willRetain | usernameFlag | passwordFlag
	m.variableHeader = append(m.variableHeader, byte(flags))
	//设置keep-alive
	m.variableHeader = append(m.variableHeader, int16ToBytes(10000)...)
	//设置properties
	m.variableHeader = append(m.variableHeader, byte(5), byte(17), byte(0), byte(0), byte(0), byte(10))
}

func (m *ConnectMessage) setPayload() {
	var length int16
	//设置client ID
	length = int16(len(m.ClientID))
	m.payload = append(m.payload, int16ToBytes(length)...)
	m.payload = append(m.payload, []byte(m.ClientID)...)

	if m.NeedAuth {
		//设置username
		usernameBytes := []byte(m.Username)
		length = int16(len(usernameBytes))
		m.payload = append(m.payload, int16ToBytes(length)...)
		m.payload = append(m.payload, usernameBytes...)

		//设置password
		passwordBytes := []byte(m.Password)
		length = int16(len(passwordBytes))
		m.payload = append(m.payload, int16ToBytes(length)...)
		m.payload = append(m.payload, passwordBytes...)
	}
}

func (m *ConnectMessage) Build() []byte {
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
