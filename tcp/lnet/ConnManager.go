package lnet

import (
	"errors"
	"fmt"
	"lai_zinx/tcp/linterface"
	"sync"
)

type ConnManager struct {
	connections map[uint32]linterface.LConnection
	connLock    sync.RWMutex
}

func newConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]linterface.LConnection),
	}

}

func (connMgr *ConnManager) Get(connID uint32) (linterface.LConnection, error) {
	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()

	if conn, ok := connMgr.connections[connID]; ok {
		return conn, nil
	}

	return nil, errors.New("connection not found")

}
func (cMgr *ConnManager) Add(conn linterface.LConnection) {
	cMgr.connLock.Lock()
	cMgr.connections[conn.GetConnID()] = conn
	cMgr.connLock.Unlock()
	fmt.Println("connection add to ConnManager successfully: conn num = ", cMgr.Len())
}

func (ConnManager *ConnManager) Remove(conn linterface.LConnection) {

	ConnManager.connLock.Lock()
	//删除连接信息
	delete(ConnManager.connections, conn.GetConnID())
	ConnManager.connLock.Unlock()
	fmt.Println("connection Remove ConnID=", conn.GetConnID(), " successfully: conn num = ", ConnManager.Len())
}
func (ConnManager *ConnManager) Len() int {
	ConnManager.connLock.RLock()
	length := len(ConnManager.connections)
	ConnManager.connLock.RUnlock()
	return length
}

//ClearConn 清除并停止所有连接
func (ConnManager *ConnManager) ClearConn() {
	ConnManager.connLock.Lock()

	//停止并删除全部的连接信息
	for connID, conn := range ConnManager.connections {
		//停止
		conn.Stop()
		//删除
		delete(ConnManager.connections, connID)
	}
	ConnManager.connLock.Unlock()
	fmt.Println("Clear All connections successfully: conn num = ", ConnManager.Len())
}
func (connMgr *ConnManager) ClearOneConn(connID uint32) {
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	connections := connMgr.connections
	if conn, ok := connections[connID]; ok {
		//停止
		conn.Stop()
		//删除
		delete(connections, connID)
		fmt.Println("Clear Connections ID:  ", connID, "succeed")
		return
	}

	fmt.Println("Clear Connections ID:  ", connID, "err")
	return
}
