package dtos

import "github.com/3115826227/go-web-live/internal/constant"

type Live struct {
	LiveRoomID string
	// 直播间主播
	LiveRoomOrigin string
	// 直播间状态
	LiveRoomStatus constant.LiveRoomStatus
	// 用户人数
	LiveRoomUserTotal int64
}

type LiveDetail struct {
	Live
	// 主播直播时长
	OnlineTime int64
	// 是否全员禁言
	PermissionSendMessage bool
}

type LiveOrigin struct {
	LiveRoomID string
	// 直播间主播
	LiveRoomOrigin string
	// 直播间状态
	LiveRoomStatus constant.LiveRoomStatus
	// 用户人数
	LiveRoomUserTotal int64
	// 最高用户人数
	LiveRoomMaxUserTotal int64
	// 举报次数
	LiveRoomReports int64
}

type LiveUser struct {
	// 直播房间id
	LiveRoomID string
	// 用户id
	AccountID string
	// 是否被禁言
	PermissionSendMessage bool
	// 加入房间的时间点
	JoinTimestamp int64
	// 是否为访客
	Visitor bool
}
