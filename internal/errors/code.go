package errors

type ErrorCode uint64

const (
	// 基本错误
	CodeInternalError         ErrorCode = 500 // 服务器错误
	CodeInvalidParamError               = 502 // 参数错误
	CodeUnLoginError                    = 401 // 用户未登录
	CodeResourceNotExistError           = 503 // 资源不存在

	// 用户错误
	CodeLoginNameNotExistError = 30001 // 用户登录账号不存在
	CodeLoginNameExistError    = 30002 // 用户登录账号重复
	CodePasswordFaultError     = 30003 // 密码错误
	CodeUnOpenLiveRoomError    = 30010 // 用户未开通直播间
	CodeOpenLiveRoomExistError = 30011 // 用户已开通直播间

	// 直播间错误
	CodeLiveRoomIdEmptyError      = 30101 // 直播间号为空
	CodeLiveRoomIdNotExistError   = 30102 // 直播间不存在
	CodeUserNotInLiveRoomError    = 30201 // 用户不在直播间中
	CodeUserNotExistLiveRoomError = 30202 // 直播间不存在该用户

	CodeUserPermissionSendMessageError     = 30301 // 用户已被禁言，无法发送消息
	CodeLiveRoomPermissionSendMessageError = 30401 // 直播间被禁言，用户无法发送消息
)
