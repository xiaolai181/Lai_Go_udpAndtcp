package linterface

type Packet interface {
	Unpack(binaryData []byte) (LMessage, error)
	Pack(msg LMessage) ([]byte, error)
	GetHeadlen() uint32
}
