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
	"time"
)

// 直播间消息主播查询
func LiveMessageOriginQueryHandle(c *gin.Context) {
	userMeta := GetUserMeta(c)
	id, exist, err := application.GetLiveRoomIdByAccountId(c, userMeta.AccountId)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, err.Code())
		return
	}
	if !exist {
		err1 := fmt.Errorf("user %v hadn't open live room", userMeta.AccountId)
		log.Logger.Error(err1.Error())
		FailedResp(c, errors.CodeUnOpenLiveRoomError)
		return
	}
	var reqPage requests.PageCommonReq
	reqPage, err = PageHandle(c)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, err.Code())
		return
	}
	var messages []dtos.Message
	if messages, err = application.OriginGetMessages(c, reqPage, id, constant.LiveRoomBizType); err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, err.Code())
		return
	}
	var ids = make([]string, 0)
	for _, msg := range messages {
		ids = append(ids, msg.Send)
	}
	var mp map[string]dtos.User
	mp, err = application.GetUserByIds(c, ids)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, err.Code())
		return
	}
	var list = make([]interface{}, 0)
	for _, msg := range messages {
		var u = mp[msg.Send]
		var message = rsp.OriginLiveRoomMessageResponse{
			ID:          msg.ID,
			MessageType: msg.MessageType,
			Send: rsp.LiveRoomOriginUser{
				AccountID:             u.AccountId,
				Username:              u.Username,
				Visitor:               u.Visitor,
				PermissionSendMessage: u.PermissionSendMessage,
			},
			Content:       msg.Content,
			SendTimestamp: msg.SendTimestamp,
		}
		list = append(list, message)
	}
	SuccessListResp(c, "", list, 0, reqPage.Page, reqPage.PageSize)
}

// 直播间消息游客查询
func MessageVisitorQueryHandle(c *gin.Context) {
	bizId := c.Query("biz_id")
	bizType, err := strconv.Atoi(c.Query("biz_type"))
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, errors.CodeInvalidParamError)
		return
	}
	var exist bool
	switch constant.BizType(bizType) {
	case constant.LiveRoomBizType:
		_, exist, err = application.QueryLiveDetail(c, bizId)
		if err != nil {
			log.Logger.Error(err.Error())
			FailedResp(c, CodeInternalError)
			return
		}
		if !exist {
			err = fmt.Errorf("live room id %v isn't exist", bizId)
			log.Logger.Errorf(err.Error())
			FailedResp(c, errors.CodeLiveRoomIdNotExistError)
			return
		}
	case constant.SessionBizType:
	}
	var accountId = GetUserMeta(c).AccountId
	var reqPage requests.PageCommonReq
	if reqPage, err = PageHandle(c); err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	var lr dtos.LiveUser
	lr, exist, err = application.GetUserRelation(c, bizId, constant.BizType(bizType), accountId)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	if !exist {
		err = fmt.Errorf("user %v is not in live room or session %v", accountId, bizId)
		log.Logger.Error(err.Error())
		FailedResp(c, CodeUserNotInLiveRoomError)
		return
	}
	var messages []dtos.Message
	messages, err = application.UserGetMessages(c, reqPage, bizId, constant.BizType(bizType), lr.JoinTimestamp)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	var ids = make([]string, 0)
	for _, msg := range messages {
		ids = append(ids, msg.Send)
	}
	var mp map[string]dtos.User
	mp, err = application.GetUserByIds(c, ids)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	var list = make([]interface{}, 0)
	for _, msg := range messages {
		var u = mp[msg.Send]
		var message = rsp.UserLiveRoomMessageResponse{
			ID:          msg.ID,
			MessageType: msg.MessageType,
			Send: rsp.LiveRoomUser{
				Username: u.Username,
				Visitor:  u.Visitor,
			},
			Content:       msg.Content,
			SendTimestamp: msg.SendTimestamp,
		}
		list = append(list, message)
	}
	SuccessListResp(c, "", list, 0, reqPage.Page, reqPage.PageSize)
}

// 消息用户查询
func MessageUserQueryHandle(c *gin.Context) {
	bizId := c.Query("biz_id")
	bizType, err := strconv.Atoi(c.Query("biz_type"))
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInvalidParams)
		return
	}
	var exist bool
	switch constant.BizType(bizType) {
	case constant.LiveRoomBizType:
		_, exist, err = application.QueryLiveDetail(c, bizId)
		if err != nil {
			log.Logger.Error(err.Error())
			FailedResp(c, CodeInternalError)
			return
		}
		if !exist {
			err = fmt.Errorf("live room id %v isn't exist", bizId)
			log.Logger.Errorf(err.Error())
			FailedResp(c, CodeLiveRoomIdNotExistError)
			return
		}
	case constant.SessionBizType:
	}
	userMeta := GetUserMeta(c)
	reqPage, err := PageHandle(c)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	var lr dtos.LiveUser
	lr, exist, err = application.GetUserRelation(c, bizId, constant.BizType(bizType), userMeta.AccountId)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	if !exist {
		err = fmt.Errorf("user %v is not in live room or session %v", userMeta.AccountId, bizId)
		log.Logger.Error(err.Error())
		FailedResp(c, CodeUserNotInLiveRoomError)
		return
	}
	var messages []dtos.Message
	messages, err = application.UserGetMessages(c, reqPage, bizId, constant.BizType(bizType), lr.JoinTimestamp)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	var ids = make([]string, 0)
	for _, msg := range messages {
		ids = append(ids, msg.Send)
	}
	var mp map[string]dtos.User
	mp, err = application.GetUserByIds(c, ids)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	var list = make([]interface{}, 0)
	for _, msg := range messages {
		var u = mp[msg.Send]
		var message = rsp.UserLiveRoomMessageResponse{
			ID:          msg.ID,
			MessageType: msg.MessageType,
			Send: rsp.LiveRoomUser{
				Username: u.Username,
				Visitor:  u.Visitor,
			},
			Content:       msg.Content,
			SendTimestamp: msg.SendTimestamp,
		}
		list = append(list, message)
	}
	SuccessListResp(c, "", list, 0, reqPage.Page, reqPage.PageSize)
}

// 直播间消息发送
func SendLiveMessageHandle(c *gin.Context) {
	var userMeta = GetUserMeta(c)
	var req requests.AddLiveRoomMessageReq
	req.LiveRoomId = c.PostForm("live_room_id")
	msgType, err := strconv.Atoi(c.PostForm("message_type"))
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInvalidParams)
		return
	}
	req.MessageType = constant.MessageType(msgType)
	req.Content = c.PostForm("content")
	if err = c.Bind(&req); err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInvalidParams)
		return
	}
	live, exist, err := application.QueryLiveDetail(c, req.LiveRoomId)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	if !exist {
		err = fmt.Errorf("live room id %v isn't exist", req.LiveRoomId)
		log.Logger.Errorf(err.Error())
		FailedResp(c, CodeLiveRoomIdNotExistError)
		return
	}
	if live.PermissionSendMessage {
		err = fmt.Errorf("live room id %v is permission send message", req.LiveRoomId)
		log.Logger.Errorf(err.Error())
		FailedResp(c, CodeLiveRoomPermissionSendMessageError)
		return
	}
	if live.LiveRoomOrigin != userMeta.AccountId {
		var lv dtos.LiveUser
		lv, exist, err = application.GetUserRelation(c, req.LiveRoomId, constant.LiveRoomBizType, userMeta.AccountId)
		if err != nil {
			log.Logger.Error(err.Error())
			FailedResp(c, CodeInternalError)
			return
		}
		if !exist {
			err = fmt.Errorf("user %v is not in live room %v", userMeta.AccountId, req.LiveRoomId)
			log.Logger.Error(err.Error())
			FailedResp(c, CodeUserNotInLiveRoomError)
			return
		}
		if lv.PermissionSendMessage {
			err = fmt.Errorf("user %v is permission send message in live room %v", userMeta.AccountId, req.LiveRoomId)
			log.Logger.Error(err.Error())
			FailedResp(c, CodeUserPermissionSendMessageError)
			return
		}
	}
	var msg = dtos.Message{
		MessageType:   req.MessageType,
		Send:          userMeta.AccountId,
		Content:       req.Content,
		SendTimestamp: time.Now().Unix(),
	}
	if err = application.AddMessage(c, msg, req.LiveRoomId, constant.LiveRoomBizType); err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	SuccessResp(c, "", nil)
}
