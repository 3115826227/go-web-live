package rsp

import "github.com/3115826227/go-web-live/internal/constant"

type LiveResponse struct {
	LiveRoomID string `json:"live_room_id"`
	// 直播间主播
	LiveRoomOrigin User `json:"live_room_origin"`
	// 直播间状态
	LiveRoomStatus constant.LiveRoomStatus `json:"live_room_status"`
	// 用户人数
	LiveRoomUserTotal int64 `json:"live_room_user_total"`
}

type LiveDetailResponse struct {
	LiveResponse
	// 主播直播时长
	OnlineTime int64 `json:"online_time"`
	// 是否全员禁言
	PermissionSendMessage bool `json:"permission_send_message"`
	// 用户是否被禁言
	UserPermissionSendMessage bool `json:"user_permission_send_message"`
}

type LiveOriginResponse struct {
	LiveRoomID string `json:"live_room_id"`
	// 直播间主播 
	LiveRoomOrigin User `json:"live_room_origin"`
	// 直播间状态
	LiveRoomStatus constant.LiveRoomStatus `json:"live_room_status"`
	// 用户人数
	LiveRoomUserTotal int64 `json:"live_room_user_total"`
	// 最高用户人数
	LiveRoomMaxUserTotal int64 `json:"live_room_max_user_total"`
	// 举报次数
	LiveRoomReports int64 `json:"live_room_reports"`
}
