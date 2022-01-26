package handle

import (
	"github.com/3115826227/go-web-live/internal/errors"
	"github.com/3115826227/go-web-live/internal/handle/rsp"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ErrCodeLoginInvalid   = 99
	CodeInvalidParams     = 400
	CodeRequiredLogin     = 401
	CodeRequiredForbidden = 403
	CodeNotFound          = 404
	CodeInternalError     = 500
	CodeServiceNotFound   = 502
	CodeUnVerifyForbidden = 600

	CodeNeedApplyAddFriend = 20010

	CodeNeedOriginAuditSession = 20021
	CodeNeedInviteJoinSession  = 20022

	CodeSessionOriginPermission = 20031

	CodePhoneInvalid = 20041
	CodePhoneEmpty   = 20042

	CodePhoneVerifyCodeTooBusy = 20011
	CodePhoneVerifyCodeError   = 20012
	CodePhoneVerifyCodeInvalid = 20013
	CodePhoneVerifyCodeExpire  = 20014
	CodePhoneVerifyCodeEmpty   = 20015

	CodeSelfVideoConflictError = 20101
	CodeUserVideoConflictError = 20102
	CodeUserOfflineError       = 20103

	CodeLoginNameExistError = 30001

	CodeUnOpenLiveRoomError    = 30010
	CodeOpenLiveRoomExistError = 30011

	CodeLiveRoomIdEmptyError    = 30101
	CodeLiveRoomIdNotExistError = 30102

	CodeUserNotInLiveRoomError    = 30201
	CodeUserNotExistLiveRoomError = 30202

	CodeUserPermissionSendMessageError     = 30301
	CodeLiveRoomPermissionSendMessageError = 30401
)

const (
	ErrCodeLoginInvalidMsg   = "登录名称不匹配或者密码错误"
	CodeInvalidParamsMsg     = "参数错误"
	CodeRequiredLoginMsg     = "请登录"
	CodeRequiredForbiddenMsg = "权限不够"
	CodeNotFoundMsg          = "未找到服务"
	CodeInternalErrorMsg     = "服务器错误"
	CodeServiceNotFoundMsg   = "服务不存在"
	CodeUnVerifyForbiddenMsg = "未认证用户无法访问"

	CodeNeedApplyAddFriendMsg = "对方已设置好友添加权限，请先申请添加好友"

	CodeNeedOriginAuditSessionMsg = "会话加入请求已发送，请耐心等待审核确认"
	CodeNeedInviteJoinSessionMsg  = "会话创建者已设置会话加入权限，请您联系会话创建者邀请您，才能加入该会话"

	CodeSessionOriginPermissionMsg = "只有会话创建才有该权限"

	// 手机号相关
	CodePhoneInvalidMsg = "无效的手机号码"
	CodePhoneEmptyMsg   = "手机号不能为空"

	// 验证码相关
	CodePhoneVerifyCodeTooBusyMsg = "验证码发送太频繁，请稍后再试"
	CodePhoneVerifyCodeErrorMsg   = "验证码申请失败，请重试"
	CodePhoneVerifyCodeInvalidMsg = "验证码无效，请填写正确的验证码"
	CodePhoneVerifyCodeExpireMsg  = "验证码已过期，请重新申请验证码"
	CodePhoneVerifyCodeEmptyMsg   = "验证码不能为空，请填写正确有效的验证码"

	CodeSelfVideoConflictErrorMsg = "您有未结束的通话，请先结束当前会话"
	CodeUserVideoConflictErrorMsg = "对方正在通话，请稍后再试"
	CodeUserOfflineErrorMsg       = "对方未上线，请稍后再试"

	CodeLoginNameExistErrorMsg = "登录名已存在"

	CodeUnOpenLiveRoomErrorMsg    = "您还未开通直播间"
	CodeOpenLiveRoomExistErrorMsg = "您已经开通了直播间"

	CodeLiveRoomIdEmptyErrorMsg      = "直播间号不能为空"
	CodeLiveRoomIdNotExistErrorMsg   = "直播间号不存在"
	CodeUserNotInLiveRoomErrorMsg    = "用户不在直播间内，请先进入该直播间"
	CodeUserNotExistLiveRoomErrorMsg = "用户不在直播间"

	CodeUserPermissionSendMessageErrorMsg     = "您已被主播禁言"
	CodeLiveRoomPermissionSendMessageErrorMsg = "直播间已被主播全员禁言"
)

const (
	InternalCodePhoneEmptyMsg   = "phone is empty"
	InternalCodePhoneInvalidMsg = "phone is invalid"
)

var ErrCodeM = map[int]string{
	ErrCodeLoginInvalid:                    ErrCodeLoginInvalidMsg,
	CodeInvalidParams:                      CodeInvalidParamsMsg,
	CodeInternalError:                      CodeInternalErrorMsg,
	CodeRequiredForbidden:                  CodeRequiredForbiddenMsg,
	CodePhoneInvalid:                       CodePhoneInvalidMsg,
	CodePhoneEmpty:                         CodePhoneEmptyMsg,
	CodePhoneVerifyCodeTooBusy:             CodePhoneVerifyCodeTooBusyMsg,
	CodePhoneVerifyCodeError:               CodePhoneVerifyCodeErrorMsg,
	CodePhoneVerifyCodeInvalid:             CodePhoneVerifyCodeInvalidMsg,
	CodePhoneVerifyCodeExpire:              CodePhoneVerifyCodeExpireMsg,
	CodePhoneVerifyCodeEmpty:               CodePhoneVerifyCodeEmptyMsg,
	CodeSelfVideoConflictError:             CodeSelfVideoConflictErrorMsg,
	CodeUserVideoConflictError:             CodeUserVideoConflictErrorMsg,
	CodeUserOfflineError:                   CodeUserOfflineErrorMsg,
	CodeLoginNameExistError:                CodeLoginNameExistErrorMsg,
	CodeUnOpenLiveRoomError:                CodeUnOpenLiveRoomErrorMsg,
	CodeOpenLiveRoomExistError:             CodeOpenLiveRoomExistErrorMsg,
	CodeLiveRoomIdEmptyError:               CodeLiveRoomIdEmptyErrorMsg,
	CodeLiveRoomIdNotExistError:            CodeLiveRoomIdNotExistErrorMsg,
	CodeUserNotInLiveRoomError:             CodeUserNotInLiveRoomErrorMsg,
	CodeUserNotExistLiveRoomError:          CodeUserNotExistLiveRoomErrorMsg,
	CodeUserPermissionSendMessageError:     CodeUserPermissionSendMessageErrorMsg,
	CodeLiveRoomPermissionSendMessageError: CodeLiveRoomPermissionSendMessageErrorMsg,
}

func SuccessResp(c *gin.Context, message string, data interface{}) {
	if data == nil {
		data = make(map[string]interface{})
	}
	c.JSON(http.StatusOK, rsp.CommonResp{Code: 0, Message: message, Data: data})
}

func SuccessListResp(c *gin.Context, message string, list []interface{}, total, page, pageSize int64) {
	if list == nil {
		list = make([]interface{}, 0)
	}
	c.JSON(http.StatusOK, rsp.CommonResp{Code: 0, Message: message, Data: rsp.CommonListResp{
		List:     list,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}})
}

func FailedResp(c *gin.Context, ec errors.ErrorCode) {
	c.JSON(http.StatusOK, rsp.CommonResp{Code: ec, Message: errors.GetZhErrorMsg(ec), Data: make(map[string]interface{})})
}
