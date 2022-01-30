package linterface

type LRequest interface {
	GetConnection() LConnection
	GetData() []byte
	GetMsgID() uint32
}
