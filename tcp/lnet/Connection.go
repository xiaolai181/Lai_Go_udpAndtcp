package lnet

import (
	"context"
	"errors"
	"fmt"
	"io"
	"lai_zinx/tcp/linterface"
	"net"
	"sync"
	"time"
)

var WorkPoolSize = 10

type Connection struct {
	TCPserver linterface.Lserver

	Conn *net.TCPConn

	ConnID uint32

	MsgHandler linterface.LMsgHandle
	ctx        context.Context
	cancal     context.CancelFunc

	msgChan     chan []byte
	msgBuffChan chan []byte
	sync.RWMutex
	property map[string]interface{}
	////保护当前property的锁
	propertyLock sync.Mutex
	isClosed     bool
}

func NewConnection(server linterface.Lserver, conn *net.TCPConn, connID uint32, msghandle linterface.LMsgHandle) *Connection {
	c := &Connection{
		TCPserver:   server,
		Conn:        conn,
		ConnID:      connID,
		isClosed:    false,
		MsgHandler:  msghandle,
		msgChan:     make(chan []byte),
		msgBuffChan: make(chan []byte, 1024),
	}
	c.TCPserver.GetConnMgr().Add(c)
	return c
}

// func (c *Connection) StartWiter() {
// 	fmt.Println("[Witer Goroutine is Running]")
// 	defer fmt.Println(c.RemoteAddr().String(), "[conn Writer exit!]")
// 	for {
// 		select {
// 		case data := <-c.msgChan:
// 			if _, err := c.Conn.Write(data); err != nil {
// 				fmt.Println("send data err:", err)
// 				return
// 			}
// 		case data, ok := <-c.msgBuffChan:
// 			if ok {
// 				if _, err := c.Conn.Write(data); err != nil {
// 					fmt.Println("send buff data err:", err)
// 					return
// 				}
// 			} else {
// 				fmt.Println("msgBuffChan is Closed")
// 				break
// 			}

// 		case <-c.ctx.Done():
// 			return
// 		}
// 	}
// }

// func (c *Connection) StartReader() {
// 	fmt.Println("[Reader Goroutine is running]")
// 	defer fmt.Println(c.RemoteAddr().String(), "[conn Reader exit!]")
// 	defer c.Stop()
// 	for {
// 		select {
// 		case <-c.ctx.Done():
// 			return
// 		default:
// 			headData := make([]byte, c.TCPserver.Packet().GetHeadlen())
// 			if _, err := io.ReadFull(c.Conn, headData); err != nil {
// 				fmt.Println("read msg head error ", err)
// 				return
// 			}
// 			msg, err := c.TCPserver.Packet().Unpack(headData)
// 			if err != nil {
// 				fmt.Println("unpack error ", err)
// 				return
// 			}
// 			var data []byte
// 			if msg.GetDatalen() > 0 {
// 				data = make([]byte, msg.GetDatalen())
// 				if _, err := io.ReadFull(c.Conn, data); err != nil {
// 					fmt.Println("read msg data error ", err)
// 					return
// 				}
// 				msg.SetData(data)
// 				req := Request{
// 					conn: c,
// 					msg:  msg,
// 				}
// 				if WorkPoolSize > 0 {
// 					c.MsgHandler.SendMsTaskQueue(&req)
// 				} else {
// 					go c.MsgHandler.DoMsgHandler(&req)
// 				}
// 			}
// 		}
// 	}
// }
func (c *Connection) StartReader() {
	fmt.Println("[Reader Goroutine is running]")
	defer fmt.Println(c.RemoteAddr().String(), "[conn Reader exit!]")
	defer c.Stop()

	// 创建拆包解包的对象
	for {
		select {
		case <-c.ctx.Done():
			return
		default:

			//读取客户端的Msg head
			headData := make([]byte, c.TCPserver.Packet().GetHeadlen())
			if _, err := io.ReadFull(c.Conn, headData); err != nil {
				fmt.Println("read msg head error ", err)
				return
			}
			//fmt.Printf("read headData %+v\n", headData)

			//拆包，得到msgID 和 datalen 放在msg中
			msg, err := c.TCPserver.Packet().Unpack(headData)
			if err != nil {
				fmt.Println("unpack error ", err)
				return
			}

			//根据 dataLen 读取 data，放在msg.Data中
			var data []byte
			if msg.GetDatalen() > 0 {
				data = make([]byte, msg.GetDatalen())
				if _, err := io.ReadFull(c.Conn, data); err != nil {
					fmt.Println("read msg data error ", err)
					return
				}
			}
			msg.SetData(data)

			//得到当前客户端请求的Request数据
			req := Request{
				conn: c,
				msg:  msg,
			}

			if WorkPoolSize > 0 {
				//已经启动工作池机制，将消息交给Worker处理
				c.MsgHandler.SendMsTaskQueue(&req)
			} else {
				//从绑定好的消息和对应的处理方法中执行对应的Handle方法
				go c.MsgHandler.DoMsgHandler(&req)
			}
		}
	}
}

func (c *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is running]")
	defer fmt.Println(c.RemoteAddr().String(), "[conn Writer exit!]")

	for {
		select {
		case data := <-c.msgChan:
			//有数据要写给客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send Data error:, ", err, " Conn Writer exit")
				return
			}
			//fmt.Printf("Send data succ! data = %+v\n", data)
		case data, ok := <-c.msgBuffChan:
			if ok {
				//有数据要写给客户端
				if _, err := c.Conn.Write(data); err != nil {
					fmt.Println("Send Buff Data error:, ", err, " Conn Writer exit")
					return
				}
			} else {
				fmt.Println("msgBuffChan is Closed")
				break
			}
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Connection) Start() {
	c.ctx, c.cancal = context.WithCancel(context.Background())
	go c.StartReader()
	go c.StartWriter()
	c.TCPserver.CallOnConnStart(c)
	select {
	case <-c.ctx.Done():
		return
	}
}
func (c *Connection) Stop() {
	c.cancal()
}
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) SendMsg(msgID uint32, data []byte) error {
	c.RLock()
	defer c.RUnlock()
	if c.isClosed == true {
		return errors.New("connection colsed when send msg")
	}
	dp := c.TCPserver.Packet()
	msg, err := dp.Pack(NewMsgPackage(msgID, data))
	if err != nil {
		fmt.Println("Pack error msg ID = ", msgID)
		return errors.New("Pack error msg ")
	}
	c.msgChan <- msg
	return nil
}
func (c *Connection) SendBuffMsg(msgID uint32, data []byte) error {
	c.RLock()
	defer c.RUnlock()
	idleTimeout := time.NewTimer(5 * time.Millisecond)
	defer idleTimeout.Stop()
	if c.isClosed == true {
		return errors.New("Connection closed when send buff msg")
	}
	dp := c.TCPserver.Packet()
	msg, err := dp.Pack(NewMsgPackage(msgID, data))
	if err != nil {
		fmt.Println("Pack error msg ID = ", msgID)
		return errors.New("Pack error msg ")
	}

	// 发送超时
	select {
	case <-idleTimeout.C:
		return errors.New("send buff msg timeout")
	case c.msgBuffChan <- msg:
		return nil
	}
	//写回客户端
	//c.msgBuffChan <- msg
	return nil
}
func (c *Connection) Context() context.Context {
	return c.ctx
}
