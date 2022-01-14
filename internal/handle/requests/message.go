package requests

import "github.com/3115826227/go-web-live/internal/constant"

type AddLiveRoomMessageReq struct {
	LiveRoomId  string               `json:"live_room_id"`
	MessageType constant.MessageType `json:"message_type"`
	Content     string               `json:"content"`
}
