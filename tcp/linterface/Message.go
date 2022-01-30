package linterface

type LMessage interface {
	GetDatalen() uint32
	GetMsgID() uint32
	GetData() []byte

	SetDatalen(uint32)
	SetMsgID(uint32)
	SetData([]byte)
}
