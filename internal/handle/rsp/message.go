package rsp

import "github.com/3115826227/go-web-live/internal/constant"

type OriginLiveRoomMessageResponse struct {
	ID            int64                `json:"id"`
	MessageType   constant.MessageType `json:"message_type"`
	Send          LiveRoomOriginUser   `json:"send"`
	Content       string               `json:"content"`
	SendTimestamp int64                `json:"send_timestamp"`
}

type UserLiveRoomMessageResponse struct {
	ID            int64                `json:"id"`
	MessageType   constant.MessageType `json:"message_type"`
	Send          LiveRoomUser         `json:"send"`
	Content       string               `json:"content"`
	SendTimestamp int64                `json:"send_timestamp"`
}
