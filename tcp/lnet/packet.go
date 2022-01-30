package lnet

import (
	"bytes"
	"encoding/binary"
	"lai_zinx/tcp/linterface"
)

var defaultHeadrLen uint32 = 8

type Datapack struct{}

func NewDataPack() linterface.Packet {
	return &Datapack{}
}
func (dp *Datapack) GetHeadlen() uint32 {
	//ID uint32(4字节) +  DataLen uint32(4字节)
	return defaultHeadrLen
}

func (dp *Datapack) Pack(msg linterface.LMessage) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDatalen()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

func (dp *Datapack) Unpack(binaryData []byte) (linterface.LMessage, error) {
	dataBuff := bytes.NewReader(binaryData)
	msg := &Message{}
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Datalen); err != nil {
	}
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.ID); err != nil {
	}
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Data); err != nil {
	}
	return msg, nil
}
