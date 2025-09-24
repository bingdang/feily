package ClientHandleMessage

import (
	"feily/Model"
	"log"
	"net"
)

func HandleMessage(msg *Model.Message, conn net.Conn) {

	// 根据消息类型处理不同逻辑
	switch msg.Type {
	case 11:
		// 打印消息
		log.Printf("ReceivedBody: %s ReceivedType: %v ClientId: %s\n", string(msg.Body), msg.Type, msg.ClientId)
	default:
		log.Printf("Unknown message type: %d\n", msg.Type)
	}
}
