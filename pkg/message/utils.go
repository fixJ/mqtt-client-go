package message

import (
	"bytes"
	"encoding/binary"
)

func int16ToBytes(n int16) []byte {
	buffer := bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.BigEndian, n)
	return buffer.Bytes()
}

func BytesToPublishMessage(data []byte) ([]PublishMessage, error) {
	var res []PublishMessage
	current := 0
	var indexs []int
	indexs = append(indexs, 0)
	for {
		if current < len(data)-1 {
			current = current + int(data[current+1]) + 2
			indexs = append(indexs, current)
			continue
		} else {
			break
		}

	}
	for i := 0; i < len(indexs)-1; i++ {
		start := indexs[i]
		end := indexs[i+1]
		d := data[start:end]
		var pm PublishMessage
		var topicLength int16
		if d[0]&byte(240) != byte(48) {
			continue
		}
		topicLength = int16(binary.BigEndian.Uint16([]byte{d[2], d[3]}))
		pm.Topic = string(d[4 : topicLength+4])
		pm.SetPacketID(int16(binary.BigEndian.Uint16([]byte{d[topicLength+4], d[topicLength+5]})))
		var propertiesLength int
		propertiesLength = int(d[topicLength+6])
		pm.Data = string(d[topicLength+6+int16(propertiesLength):])
		res = append(res, pm)
	}
	return res, nil
}
