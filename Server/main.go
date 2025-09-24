package main

import (
	"feily/Encode"
	ServerHandleMessage "feily/Server/handleMessage"
	SessionManager "feily/Server/sessionManager"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	manager := SessionManager.NewClientManager()

	for {
		// 解码消息
		msg, err := Encode.Decode(conn)
		if err != nil {
			log.Println("Decode error:", err)
			return
		}
		manager.AddClient(conn, msg.ClientId)

		// 处理消息
		ServerHandleMessage.HandleMessage(msg, conn)

	}
}

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Listen error:", err)
	}
	defer listener.Close()

	log.Println("Server started on :8080")

	for {
		conn, err := listener.Accept()
		log.Printf("%s < %s", conn.LocalAddr(), conn.RemoteAddr())
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}

		go handleConnection(conn)
	}
}
