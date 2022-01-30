package iface

type Message interface {
	GetMsgID() uint32
	GetDataLen() uint32
	GetKeyLen() uint32
	GetKey() []byte
	GetMsgData() []byte

	SetMsgID(uint32)
	SetDataLen(uint32)
	SetKeyLen(uint32)
	SetKey([]byte)
	SetMsgData([]byte)
}
