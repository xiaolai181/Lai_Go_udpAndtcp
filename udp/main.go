package main

import (
	"fmt"
	"net"
	"os"
)

// 限制goroutine数量
var limitChan = make(chan bool, 1)

// UDP goroutine 实现并发读取UDP数据
func udpProcess(conn *net.UDPConn, i int) {
	fmt.Println("goroutine is ", i, " running")
	// 最大读取数据大小
	data := make([]byte, 1024)
	n, remote, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("failed read udp msg, error: " + err.Error())
	}
	str := string(data[:n])
	fmt.Println("receive from client, data:"+str, "in ", i, " gotoutine")
	n, err = conn.WriteToUDP([]byte("hello"), remote)
	if err != nil {
		fmt.Println("failed [wirte] udp msg, error: " + err.Error())
	}
	<-limitChan
}

func udpServer(address string) {

	udpAddr, err := net.ResolveUDPAddr("udp", address)
	conn, err := net.ListenUDP("udp", udpAddr)
	defer conn.Close()
	if err != nil {
		fmt.Println("read from connect failed, err:" + err.Error())
		os.Exit(1)
	}
	i := 0
	for {
		limitChan <- true
		go udpProcess(conn, i)
		i++
	}
}

func main() {
	address := "0.0.0.0:8080"
	udpServer(address)
}