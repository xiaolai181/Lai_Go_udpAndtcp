package main

import (
	"fmt"
	"net"
)

// UDP 客户端
func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("连接UDP服务器失败，err: ", err)
		return
	}
	defer socket.Close()
	sendData := []byte("Hello Server")
	for i := 0; i < 10; i++ {
		_, err = socket.Write(sendData) // 发送数据
		if err != nil {
			fmt.Println("发送数据失败，err: ", err)
		}
	}

	data := make([]byte, 4096)

	n, _, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		fmt.Println("接收数据失败, err: ", err)
		return
	}
	fmt.Printf("recv:%v", string(data[:n]))
}
