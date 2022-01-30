package iudp

import (
	"bytes"
	"encoding/binary"
	"lai_zinx/udp/iface"
)

type Unpack struct {
}

func (up *Unpack) Unpack(binaryData []byte) (iface.Message, error) {
	dataBuff := bytes.NewReader(binaryData)
	msg := &Message{}
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.ID); err != nil {
	}
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.KeyLen); err != nil {
	}
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
	}
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Key); err != nil {
	}
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Data); err != nil {
	}
	return msg, nil
}
