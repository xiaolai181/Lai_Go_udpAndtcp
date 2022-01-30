package linterface

import (
	"context"
	"net"
)

type LConnection interface {
	Start()
	Stop()
	Context() context.Context

	GetTCPConnection() *net.TCPConn
	GetConnID() uint32
	RemoteAddr() net.Addr
	SendMsg(msgID uint32, data []byte) error
	SendBuffMsg(msgID uint32, data []byte) error
}
