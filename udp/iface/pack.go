package iface

type Pack interface {
	Pack(Message) ([]byte, error)
}
type Head interface {
	GetHeadlen() (uint32, uint32, uint32)
}
