package constant

const (
	TimeLayout = "2006-01-02 15:04:05"
	DateLayout = "2006-01-02"
)

const (
	PrivateMessageIDDefaultLength = 12
	DefaultPageSize               = 10
	DefaultPage                   = 1

	DefaultUserEncryMd5 = "md5"
)

type LiveRoomStatus int64

const (
	// 直播未开播
	LiveClosed = 0
	// 直播间主播离线
	LiveOriginOffline = 11
	// 直播主播在线
	LiveOriginOnline = 200
)

type LiveOperator int64

const (
	// 主播上线
	LiveRoomOnline LiveOperator = 1
	// 主播下线
	LiveRoomOffline = 2
	// 全员禁言
	NoPermissionSendMessage = 201
	// 全员解除禁言
	PermissionSendMessage = 202
)

type OriginOperator int64

const (
	// 将用户从直播间/会话移除
	RemoveUser OriginOperator = 1
	// 将用户禁言
	NoUserPermissionSendMessage = 101
	// 解除用户禁言
	UserPermissionSendMessage = 102
)

type UserOperator int64

const (
	// 用户进入直播间/会话操作
	UserEnter UserOperator = 1
	// 用户离开直播间/会话操作
	UserLeave = 2
)

type MessageType int64

const (
	// 默认文字消息
	DefaultTextMessage MessageType = 0
	// 图片消息
	ImageMessage = 1
	// 文件消息
	FileMessage = 2
)

type BizType int64

const (
	// 直播间会话
	LiveRoomBizType BizType = 1
	// 群聊会话
	SessionBizType = 2
)
