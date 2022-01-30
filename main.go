package main

import (
	"fmt"
	"lai_zinx/tcp/linterface"
	"lai_zinx/tcp/lnet"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	Ch_byte = make(chan linterface.LRequest, 1)
	go Ks()
	Serve := lnet.NewServe("Lennon")
	Serve.AddRouter(1, &handle{})
	Serve.Serve()

}

var req linterface.LRequest

type handle struct {
	lnet.BaseRouter
}

func (this *handle) Handle(request linterface.LRequest) {
	Ch_byte <- request
	//先读取客户端的数据
	// fmt.Printf("recv from client : msgId=%d , data=%X", request.GetMsgID(), request.GetData())
	// if i == 10 {
	// 	request.GetConnection().GetTCPConnection().Close()
	// 	i = 0
	// }
	//再回写ping...ping...ping
	// err := request.GetConnection().SendBuffMsg(1, []byte("re"))
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

var Ch_byte chan linterface.LRequest

var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	defer fmt.Println("web end")
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()
	_, usermessage, err := conn.ReadMessage()
	if err != nil {
		log.Println("Error during message reading:", err)
		return
	}
	fmt.Println(string(usermessage))

	var req linterface.LRequest
	// The event loop
	for {
		req = <-Ch_byte
		fmt.Printf("%X", req.GetData())
		conn.WriteMessage(websocket.BinaryMessage, req.GetData())

	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index Page")
}

func Ks() {
	http.HandleFunc("/socket", SocketHandler)
	http.HandleFunc("/", Home)
	fmt.Println("http ws:", 8080)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
