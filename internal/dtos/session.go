package dtos

type Session struct {
	SessionID string
	// 直播间主播
	SessionOrigin string
	// 用户人数
	SessionUserTotal int64
}

type SessionDetail struct {
	Session
	// 是否全员禁言
	PermissionSendMessage bool
}

type SessionOrigin struct {
	SessionID string
	// 会话创建者
	SessionOrigin string
	// 用户人数
	SessionUserTotal int64
	// 举报次数
	SessionReports int64
}
