package Model

// Message 消息结构
type Message struct {
	Type     uint32 // 消息类型
	Length   uint32 // 消息体长度
	ClientId string
	Body     []byte // 消息体内容
}
