package handle

import (
	"fmt"
	"github.com/3115826227/go-web-live/internal/application"
	"github.com/3115826227/go-web-live/internal/constant"
	"github.com/3115826227/go-web-live/internal/dtos"
	"github.com/3115826227/go-web-live/internal/errors"
	"github.com/3115826227/go-web-live/internal/handle/requests"
	"github.com/3115826227/go-web-live/internal/handle/rsp"
	"github.com/3115826227/go-web-live/internal/log"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 查询自己直播间
func LiveOriginHandle(c *gin.Context) {
	userMeta := GetUserMeta(c)
	detail, exist, err := application.GetLiveRoomByAccountId(c, userMeta.AccountId)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, err.Code())
		return
	}
	if !exist {
		err1 := fmt.Errorf("user %v didn't open live room", userMeta.AccountId)
		log.Logger.Error(err1.Error())
		FailedResp(c, errors.CodeUnOpenLiveRoomError)
		return
	}
	var user dtos.User
	user, err = application.GetUserById(c, detail.LiveRoomOrigin)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	var response = rsp.LiveOriginResponse{
		LiveRoomID: detail.LiveRoomID,
		LiveRoomOrigin: rsp.User{
			AccountID: user.AccountId,
			Username:  user.Username,
		},
		LiveRoomStatus:       detail.LiveRoomStatus,
		LiveRoomUserTotal:    detail.LiveRoomUserTotal,
		LiveRoomMaxUserTotal: detail.LiveRoomMaxUserTotal,
		LiveRoomReports:      detail.LiveRoomReports,
	}
	SuccessResp(c, "", response)
}

// 开通直播间
func OpenLiveHandle(c *gin.Context) {
	userMeta := GetUserMeta(c)
	_, exist, err := application.GetLiveRoomByAccountId(c, userMeta.AccountId)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	if exist {
		err1 := fmt.Errorf("user %v had open live room", userMeta.AccountId)
		log.Logger.Errorf(err1.Error())
		FailedResp(c, CodeOpenLiveRoomExistError)
		return
	}
	if err = application.OpenLiveRoom(c, userMeta.AccountId); err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, err.Code())
		return
	}
	SuccessResp(c, "", nil)
}

// 直播间列表查询
func LiveHandle(c *gin.Context) {
	_, exist := c.Get(GinContextKeyUserMeta)
	var accountId string
	if exist {
		userMeta := GetUserMeta(c)
		accountId = userMeta.AccountId
	}
	reqPage, err := PageHandle(c)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, err.Code())
		return
	}
	lives, total, err := application.QueryLive(c, reqPage, accountId)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, err.Code())
		return
	}
	var accountIds = make([]string, 0)
	for _, live := range lives {
		accountIds = append(accountIds, live.LiveRoomOrigin)
	}
	var userMap map[string]dtos.User
	userMap, err = application.GetUserByIds(c, accountIds)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, err.Code())
		return
	}
	var list = make([]interface{}, 0)
	for _, live := range lives {
		var user = userMap[live.LiveRoomOrigin]
		list = append(list, rsp.LiveResponse{
			LiveRoomID: live.LiveRoomID,
			LiveRoomOrigin: rsp.User{
				AccountID: user.AccountId,
				Username:  user.Username,
			},
			LiveRoomStatus:    live.LiveRoomStatus,
			LiveRoomUserTotal: live.LiveRoomUserTotal,
		})
	}
	SuccessListResp(c, "", list, total, reqPage.Page, reqPage.PageSize)
}

// 直播间访客查询
func LiveDetailVisitorQueryHandle(c *gin.Context) {
	id := c.Query("live_room_id")
	if id == "" {
		err := fmt.Errorf("live room id isn't empty")
		log.Logger.Error(err.Error())
		FailedResp(c, CodeLiveRoomIdEmptyError)
		return
	}
	detail, exist, err := application.QueryLiveDetail(c, id)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	if !exist {
		err = fmt.Errorf("live room id %v isn't exist", id)
		log.Logger.Errorf(err.Error())
		FailedResp(c, CodeLiveRoomIdNotExistError)
		return
	}
	var accountId = GetUserMeta(c).AccountId
	var lr dtos.LiveUser
	lr, exist, err = application.GetUserRelation(c, detail.LiveRoomID, constant.LiveRoomBizType, accountId)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	if !exist {
		err = fmt.Errorf("user %v is not in live room %v", accountId, detail.LiveRoomID)
		log.Logger.Error(err.Error())
		FailedResp(c, CodeUserNotInLiveRoomError)
		return
	}
	var user dtos.User
	user, err = application.GetUserById(c, detail.LiveRoomOrigin)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	var response = rsp.LiveDetailResponse{
		LiveResponse: rsp.LiveResponse{
			LiveRoomID: detail.LiveRoomID,
			LiveRoomOrigin: rsp.User{
				AccountID: user.AccountId,
				Username:  user.Username,
			},
			LiveRoomStatus:    detail.LiveRoomStatus,
			LiveRoomUserTotal: detail.LiveRoomUserTotal,
		},
		OnlineTime:                detail.OnlineTime,
		PermissionSendMessage:     detail.PermissionSendMessage,
		UserPermissionSendMessage: lr.PermissionSendMessage,
	}
	SuccessResp(c, "", response)
}

// 直播间详情查询
func LiveDetailUserQueryHandle(c *gin.Context) {
	id := c.Query("live_room_id")
	if id == "" {
		err := fmt.Errorf("live room id isn't empty")
		log.Logger.Error(err.Error())
		FailedResp(c, CodeLiveRoomIdEmptyError)
		return
	}
	detail, exist, err := application.QueryLiveDetail(c, id)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	if !exist {
		err = fmt.Errorf("live room id %v isn't exist", id)
		log.Logger.Errorf(err.Error())
		FailedResp(c, CodeLiveRoomIdNotExistError)
		return
	}
	userMeta := GetUserMeta(c)
	var lr dtos.LiveUser
	lr, exist, err = application.GetUserRelation(c, detail.LiveRoomID, constant.LiveRoomBizType, userMeta.AccountId)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	if detail.LiveRoomOrigin != userMeta.AccountId && !exist {
		err = fmt.Errorf("user %v is not in live room %v", userMeta.AccountId, detail.LiveRoomID)
		log.Logger.Error(err.Error())
		FailedResp(c, CodeUserNotInLiveRoomError)
		return
	}
	var user dtos.User
	user, err = application.GetUserById(c, detail.LiveRoomOrigin)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	var response = rsp.LiveDetailResponse{
		LiveResponse: rsp.LiveResponse{
			LiveRoomID: detail.LiveRoomID,
			LiveRoomOrigin: rsp.User{
				AccountID: user.AccountId,
				Username:  user.Username,
			},
			LiveRoomStatus:    detail.LiveRoomStatus,
			LiveRoomUserTotal: detail.LiveRoomUserTotal,
		},
		OnlineTime:                detail.OnlineTime,
		PermissionSendMessage:     detail.PermissionSendMessage,
		UserPermissionSendMessage: lr.PermissionSendMessage,
	}
	SuccessResp(c, "", response)
}

// 直播间用户列表查询（仅主播可见）
func LiveUserHandle(c *gin.Context) {
	userMeta := GetUserMeta(c)
	id, _, err := application.GetLiveRoomIdByAccountId(c, userMeta.AccountId)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, err.Code())
		return
	}
	var reqPage requests.PageCommonReq
	reqPage, err = PageHandle(c)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	users, total, err := application.QueryUserRelations(c, reqPage, id, constant.LiveRoomBizType)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	var list = make([]interface{}, 0)
	for _, u := range users {
		list = append(list, rsp.LiveRoomOriginUser{
			AccountID:             u.AccountId,
			Username:              u.Username,
			Visitor:               u.Visitor,
			PermissionSendMessage: u.PermissionSendMessage,
		})
	}
	SuccessListResp(c, "", list, total, reqPage.Page, reqPage.PageSize)
}

// 直播间主播的操作（全员禁言/下线/上线）
func LiveOriginOperatorHandle(c *gin.Context) {
	userMeta := GetUserMeta(c)
	_, exist, err := application.GetLiveRoomIdByAccountId(c, userMeta.AccountId)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	if !exist {
		err1 := fmt.Errorf("user %v hadn't open live room", userMeta.AccountId)
		log.Logger.Error(err1.Error())
		FailedResp(c, CodeUnOpenLiveRoomError)
		return
	}
	opt, err1 := strconv.Atoi(c.PostForm("live_operator"))
	if err1 != nil {
		log.Logger.Error(err1.Error())
		FailedResp(c, errors.CodeInternalError)
		return
	}
	var req requests.LiveOperatorReq
	req.LiveOperator = constant.LiveOperator(opt)
	if err1 = c.Bind(&req); err1 != nil {
		log.Logger.Error(err1.Error())
		FailedResp(c, errors.CodeInvalidParamError)
		return
	}
	switch req.LiveOperator {
	case constant.LiveRoomOnline:
		if err = application.UpdateLiveRoomStatus(c, userMeta.AccountId, constant.LiveOriginOnline); err != nil {
			log.Logger.Error(err.Error())
			FailedResp(c, CodeInternalError)
			return
		}
	case constant.LiveRoomOffline:
		if err = application.UpdateLiveRoomStatus(c, userMeta.AccountId, constant.LiveOriginOffline); err != nil {
			log.Logger.Error(err.Error())
			FailedResp(c, CodeInternalError)
			return
		}
	case constant.PermissionSendMessage:
		if err = application.UpdateLiveRoomSendMessagePermission(c, userMeta.AccountId, false); err != nil {
			log.Logger.Error(err.Error())
			FailedResp(c, CodeInternalError)
			return
		}
	case constant.NoPermissionSendMessage:
		if err = application.UpdateLiveRoomSendMessagePermission(c, userMeta.AccountId, true); err != nil {
			log.Logger.Error(err.Error())
			FailedResp(c, CodeInternalError)
			return
		}
	default:
		FailedResp(c, CodeInvalidParams)
		return
	}
	SuccessResp(c, "", nil)
}

// 对用户的操作（移除/禁言/解除禁言）
func OperatorUserHandle(c *gin.Context) {
	userMeta := GetUserMeta(c)
	var req requests.LiveOriginOperatorReq
	req.AccountId = c.PostForm("account_id")
	bizType, err := strconv.Atoi(c.PostForm("biz_type"))
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInvalidParams)
		return
	}
	req.BizType = constant.BizType(bizType)
	var opt int
	opt, err = strconv.Atoi(c.PostForm("live_origin_operator"))
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInvalidParams)
		return
	}
	req.OriginOperator = constant.OriginOperator(opt)
	if err = c.Bind(&req); err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInvalidParams)
		return
	}
	var bizId string
	var exist bool
	switch req.BizType {
	case constant.LiveRoomBizType:
		bizId, exist, err = application.GetLiveRoomIdByAccountId(c, userMeta.AccountId)
		if err != nil {
			log.Logger.Error(err.Error())
			FailedResp(c, CodeInternalError)
			return
		}
		if !exist {
			err = fmt.Errorf("user %v hadn't open live room", userMeta.AccountId)
			log.Logger.Error(err.Error())
			FailedResp(c, CodeUnOpenLiveRoomError)
			return
		}
	case constant.SessionBizType:
	}
	var lv dtos.LiveUser
	lv, exist, err = application.GetUserRelation(c, bizId, req.BizType, req.AccountId)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	if !exist {
		err = fmt.Errorf("user %v is not in live room or session %v", req.AccountId, bizId)
		log.Logger.Error(err.Error())
		FailedResp(c, CodeUserNotExistLiveRoomError)
		return
	}
	switch req.OriginOperator {
	case constant.RemoveUser:
		if err = application.DeleteUserRelation(c, bizId, req.BizType, req.AccountId); err != nil {
			log.Logger.Error(err.Error())
			FailedResp(c, CodeInternalError)
			return
		}
	case constant.NoUserPermissionSendMessage:
		if !lv.Visitor {
			if err = application.UpdateLiveRoomUserSendMessagePermission(c, bizId, req.BizType, req.AccountId, true); err != nil {
				log.Logger.Error(err.Error())
				FailedResp(c, CodeInternalError)
				return
			}
		}
	case constant.UserPermissionSendMessage:
		if !lv.Visitor {
			if err = application.UpdateLiveRoomUserSendMessagePermission(c, bizId, req.BizType, req.AccountId, false); err != nil {
				log.Logger.Error(err.Error())
				FailedResp(c, CodeInternalError)
				return
			}
		}
	default:
		FailedResp(c, CodeInvalidParams)
		return
	}
	SuccessResp(c, "", nil)
}

func VisitorOperatorHandle(c *gin.Context) {
	var req requests.LiveUserOperatorReq
	req.BizId = c.PostForm("biz_id")
	bizType, err := strconv.Atoi(c.PostForm("biz_type"))
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInvalidParams)
		return
	}
	req.BizType = constant.BizType(bizType)
	var opt int
	opt, err = strconv.Atoi(c.PostForm("live_user_operator"))
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInvalidParams)
		return
	}
	req.UserOperator = constant.UserOperator(opt)
	if err = c.Bind(&req); err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInvalidParams)
		return
	}
	switch req.BizType {
	case constant.LiveRoomBizType:
		var exist bool
		_, exist, err = application.QueryLiveDetail(c, req.BizId)
		if err != nil {
			log.Logger.Error(err.Error())
			FailedResp(c, CodeInternalError)
			return
		}
		if !exist {
			err = fmt.Errorf("live room id %v isn't exist", req.BizId)
			log.Logger.Errorf(err.Error())
			FailedResp(c, CodeLiveRoomIdNotExistError)
			return
		}
	case constant.SessionBizType:
	}
	var accountId = GetUserMeta(c).AccountId
	switch req.UserOperator {
	case constant.UserEnter:
		if err = application.AddUserRelation(c, req.BizId, req.BizType, accountId, true); err != nil {
			log.Logger.Error(err.Error())
			FailedResp(c, CodeInternalError)
			return
		}
	case constant.UserLeave:
		if err = application.DeleteUserRelation(c, req.BizId, req.BizType, accountId); err != nil {
			log.Logger.Error(err.Error())
			FailedResp(c, CodeInternalError)
			return
		}
	default:
		FailedResp(c, CodeInvalidParams)
		return
	}
	SuccessResp(c, "", nil)
}

// 用户进入/退出直播间
func UserOperatorHandle(c *gin.Context) {
	var userMeta = GetUserMeta(c)
	var accountId = userMeta.AccountId
	var req requests.LiveUserOperatorReq
	req.BizId = c.PostForm("biz_id")
	bizType, err := strconv.Atoi(c.PostForm("biz_type"))
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInvalidParams)
		return
	}
	req.BizType = constant.BizType(bizType)
	var opt int
	opt, err = strconv.Atoi(c.PostForm("live_user_operator"))
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInvalidParams)
		return
	}
	req.UserOperator = constant.UserOperator(opt)
	if err = c.Bind(&req); err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInvalidParams)
		return
	}
	switch req.BizType {
	case constant.LiveRoomBizType:
		var live dtos.LiveDetail
		var exist bool
		live, exist, err = application.QueryLiveDetail(c, req.BizId)
		if err != nil {
			log.Logger.Error(err.Error())
			FailedResp(c, CodeInternalError)
			return
		}
		if !exist {
			err = fmt.Errorf("live room id %v isn't exist", req.BizId)
			log.Logger.Errorf(err.Error())
			FailedResp(c, CodeLiveRoomIdNotExistError)
			return
		}
		if live.LiveRoomOrigin == userMeta.AccountId {
			SuccessResp(c, "", nil)
			return
		}
	case constant.SessionBizType:
	}
	switch req.UserOperator {
	case constant.UserEnter:
		if err = application.AddUserRelation(c, req.BizId, req.BizType, accountId, false); err != nil {
			log.Logger.Error(err.Error())
			FailedResp(c, CodeInternalError)
			return
		}
	case constant.UserLeave:
		if err = application.DeleteUserRelation(c, req.BizId, req.BizType, accountId); err != nil {
			log.Logger.Error(err.Error())
			FailedResp(c, CodeInternalError)
			return
		}
	default:
		FailedResp(c, CodeInvalidParams)
		return
	}
	SuccessResp(c, "", nil)
}
