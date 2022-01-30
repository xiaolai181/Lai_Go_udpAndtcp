package linterface

type Lserver interface {
	Start()
	Stop()
	Serve()
	AddRouter(msgId uint32, route LRouter)
	GetConnMgr() LConnManager
	SetOnConnStart(func(LConnection)) //设置该Server的连接创建时Hook函数
	SetOnConnStop(func(LConnection))  //设置该Server的连接断开时的Hook函数
	CallOnConnStart(conn LConnection) //调用连接OnConnStart Hook函数
	CallOnConnStop(conn LConnection)  //调用连接OnConnStop Hook函数
	Packet() Packet
}
