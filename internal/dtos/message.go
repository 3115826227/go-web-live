package dtos

import "github.com/3115826227/go-web-live/internal/constant"

type Message struct {
	// 消息id 自增
	ID int64
	// 消息类型
	MessageType constant.MessageType
	// 发送者id
	Send string
	// 消息内容
	Content string
	// 消息发送时间
	SendTimestamp int64
}

type Messages []Message

func (messages Messages) Len() int {
	return len(messages)
}

func (messages Messages) Swap(i, j int) {
	messages[i], messages[j] = messages[j], messages[i]
}

func (messages Messages) Less(i, j int) bool {
	return messages[i].SendTimestamp < messages[j].SendTimestamp
}
