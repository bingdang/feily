package Encode

import (
	"encoding/binary"
	"feily/Model"
	"io"
)

// Encode 编码消息
func Encode(msg *Model.Message) []byte {
	// 计算总长度: 类型(4字节) + 长度字段(4字节) + 消息体长度
	buffer := make([]byte, 44+len(msg.Body))

	// 写入消息类型
	binary.BigEndian.PutUint32(buffer[0:4], msg.Type)

	// 写入消息体长度
	binary.BigEndian.PutUint32(buffer[4:8], uint32(len(msg.Body)))

	//写入uid
	copy(buffer[8:44], []byte(msg.ClientId))

	// 写入消息体
	copy(buffer[44:], msg.Body)

	return buffer
}

// Decode 解码消息
func Decode(reader io.Reader) (*Model.Message, error) {
	// 读取消息头(8字节) + 机器uid 36字节
	header := make([]byte, 44)
	_, err := io.ReadFull(reader, header)
	if err != nil {
		return nil, err
	}

	// 解析消息类型和长度
	msgType := binary.BigEndian.Uint32(header[0:4])
	bodyLength := binary.BigEndian.Uint32(header[4:8])
	ClientId := header[8:44]

	// 读取消息体
	body := make([]byte, bodyLength)
	_, err = io.ReadFull(reader, body)
	if err != nil {
		return nil, err
	}

	return &Model.Message{
		Type:     msgType,
		Length:   bodyLength,
		ClientId: string(ClientId),
		Body:     body,
	}, nil
}
