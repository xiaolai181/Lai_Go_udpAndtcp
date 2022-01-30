package lnet

import (
	"fmt"
	"lai_zinx/tcp/linterface"
	"net"
)

type Serve struct {
	Name        string
	IPVersion   string
	IP          string
	Port        int
	ConnMgr     linterface.LConnManager
	Msghandle   linterface.LMsgHandle
	OnConnStart func(conn linterface.LConnection)
	OnConnStop  func(conn linterface.LConnection)
	packet      linterface.Packet
}

func NewServe(name string) linterface.Lserver {
	s := &Serve{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
		ConnMgr:   newConnManager(),
		Msghandle: NewMsgHandle(),
		packet:    NewDataPack(),
	}
	return s
}

func (s *Serve) Start() {
	fmt.Printf("[START] Serve name: %s,listenner at IP: %s, Port %d is starting\n", s.Name, s.IP, s.Port)
	go func() {
		s.Msghandle.StartWorkerPool()
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			panic(err)
		}
		fmt.Println("strat Linx TCP Serve", s.Name, "suncceful ,now listenning... ")
		var cID uint32
		cID = 0
		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}
			fmt.Println("Get conn remote addr = ", conn.RemoteAddr().String())

			//3.2 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
			if s.ConnMgr.Len() >= 10 {
				conn.Close()
				continue
			}

			//3.3 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
			dealConn := NewConnection(s, conn, cID, s.Msghandle)

			cID++

			//3.4 启动当前链接的处理业务
			go dealConn.Start()
		}

	}()

}

func (s *Serve) Stop() {
	fmt.Println("[STOP] Linx Serve , name ", s.Name)

	//将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	s.ConnMgr.ClearConn()

}
func (s *Serve) Serve() {
	s.Start()

	//TODO Serve.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加

	//阻塞,否则主Go退出， listenner的go将会退出
	select {}
}
func (s *Serve) AddRouter(msgID uint32, router linterface.LRouter) {
	s.Msghandle.AddRouter(msgID, router)
}

//GetConnMgr 得到链接管理
func (s *Serve) GetConnMgr() linterface.LConnManager {
	return s.ConnMgr
}

//SetOnConnStart 设置该Serve的连接创建时Hook函数
func (s *Serve) SetOnConnStart(hookFunc func(linterface.LConnection)) {
	s.OnConnStart = hookFunc
}

//SetOnConnStop 设置该Serve的连接断开时的Hook函数
func (s *Serve) SetOnConnStop(hookFunc func(linterface.LConnection)) {
	s.OnConnStop = hookFunc
}

//CallOnConnStart 调用连接OnConnStart Hook函数
func (s *Serve) CallOnConnStart(conn linterface.LConnection) {
	if s.OnConnStart != nil {
		fmt.Println("---> CallOnConnStart....")
		s.OnConnStart(conn)
	}
}

//CallOnConnStop 调用连接OnConnStop Hook函数
func (s *Serve) CallOnConnStop(conn linterface.LConnection) {
	if s.OnConnStop != nil {
		fmt.Println("---> CallOnConnStop....")
		s.OnConnStop(conn)
	}
}

func (s *Serve) Packet() linterface.Packet {
	return s.packet
}

func printLogo() {
	fmt.Printf("[Zinx] Version: %s, MaxConn: %d",
		"0.0",
		10000,
	)
}
