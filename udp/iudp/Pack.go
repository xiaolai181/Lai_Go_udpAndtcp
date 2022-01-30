package iudp

import (
	"bytes"
	"encoding/binary"
	"lai_zinx/udp/iface"
)

type Pack struct {
}

func (p *Pack) Pack(msg iface.Message) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetKeyLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetKey()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

type Head_Len struct {
	Head_ID      uint32
	Head_Key_Len uint32
}

var Head = &Head_Len{
	Head_ID:      4,
	Head_Key_Len: 4,
}

func (h *Head_Len) GetHeadlen() *Head_Len {
	return Head
}
