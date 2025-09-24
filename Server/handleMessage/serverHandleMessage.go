package ServerHandleMessage

import (
	"feily/Encode"
	"feily/Model"
	"log"
	"net"
	"strings"
)

func HandleMessage(msg *Model.Message, conn net.Conn) {
	log.Printf("ReceivedBody: %s ReceivedType: %v ClientId: %s\n", string(msg.Body), msg.Type, msg.ClientId)
	// 根据消息类型处理不同逻辑
	switch msg.Type {
	case 1:
		// 回显消息
		response := &Model.Message{
			Type:     11,
			ClientId: HandleClientId(msg.ClientId),
			Body:     []byte("Hello Client"),
		}
		encoded := Encode.Encode(response)
		conn.Write(encoded)
	default:
		log.Printf("Unknown message type: %d\n", msg.Type)
	}
}

// 处理客户端id
func HandleClientId(ClientId string) (ServerId string) {
	ServerIdSplit := strings.Split(ClientId, "-")
	ServerIdSplit[3] = "55AA"
	return strings.Join(ServerIdSplit, "-")
}
