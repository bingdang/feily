package main

import (
	ClientHandleMessage "feily/Client/handleMessage"
	"feily/Encode"
	"feily/Model"
	"github.com/denisbrodbeck/machineid"
	"log"
	"net"
)

func main() {
	MyUid, err := machineid.ID()

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	//发送消息
	msg := &Model.Message{
		Type:     1,
		ClientId: MyUid,
		Body:     []byte("Hello Server"),
	}

	encoded := Encode.Encode(msg)
	conn.Write(encoded)
	log.Printf("Sent message: %s\n", string(msg.Body))

	for {
		// 接收消息
		response, err := Encode.Decode(conn)
		if err != nil {
			log.Println("Decode response error:", err)
			return
		}
		ClientHandleMessage.HandleMessage(response, conn)
	}

}
