package iface

type Unpack interface {
	Unpack([]byte) (Message, error)
}
