package linterface

type LConnManager interface {
	Add(conn LConnection)
	Remove(conn LConnection)
	Get(ConnID uint32) (LConnection, error)
	Len() int   //统计链接数量
	ClearConn() //清除所有链接
}
