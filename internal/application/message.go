package application

import (
	"context"
	"github.com/3115826227/go-web-live/internal/constant"
	"github.com/3115826227/go-web-live/internal/db/infrastructure/dbclient"
	"github.com/3115826227/go-web-live/internal/db/tables"
	"github.com/3115826227/go-web-live/internal/dtos"
	"github.com/3115826227/go-web-live/internal/handle/requests"
	"gorm.io/gorm"
	"sort"
	"time"
)

func AddMessage(ctx context.Context, message dtos.Message, bizId string, bizType constant.BizType) error {
	var msg = tables.Message{
		BizID:         bizId,
		BizType:       bizType,
		MessageType:   message.MessageType,
		Send:          message.Send,
		Content:       message.Content,
		SendTimestamp: message.SendTimestamp,
	}
	return dbclient.GetDBClient().AddMessage(msg)
}

func OriginGetMessages(ctx context.Context, req requests.PageCommonReq, bizId string, bizType constant.BizType) ([]dtos.Message, error) {
	var list = make([]dtos.Message, 0)
	messages, err := dbclient.GetDBClient().GetMessagesByBiz(bizId, bizType, 0, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	for _, msg := range messages {
		list = append(list, dtos.Message{
			ID:            msg.ID,
			MessageType:   msg.MessageType,
			Send:          msg.Send,
			Content:       msg.Content,
			SendTimestamp: msg.SendTimestamp,
		})
	}
	sort.Sort(dtos.Messages(list))
	return list, nil
}

func UserGetMessages(ctx context.Context, req requests.PageCommonReq, bizId string, bizType constant.BizType, joinTimestamp int64) ([]dtos.Message, error) {
	var list = make([]dtos.Message, 0)
	messages, err := dbclient.GetDBClient().GetMessagesByBiz(bizId, bizType, joinTimestamp, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	for _, msg := range messages {
		list = append(list, dtos.Message{
			ID:            msg.ID,
			MessageType:   msg.MessageType,
			Send:          msg.Send,
			Content:       msg.Content,
			SendTimestamp: msg.SendTimestamp,
		})
	}
	sort.Sort(dtos.Messages(list))
	return list, nil
}

func AddUserRelation(ctx context.Context, bizId string, bizType constant.BizType, accountId string, visitor bool) error {
	var relation = tables.UserRelation{
		BizID:         bizId,
		BizType:       bizType,
		AccountID:     accountId,
		JoinTimestamp: time.Now().Unix(),
		Visitor:       visitor,
	}
	return dbclient.GetDBClient().AddUserRelation(relation)
}

func QueryUserRelations(ctx context.Context, req requests.PageCommonReq, bizId string, bizType constant.BizType) ([]dtos.User, int64, error) {
	relations, total, err := dbclient.GetDBClient().GetUserRelations(bizId, bizType, req.Page, req.PageSize)
	if err != nil {
		return nil, 0, err
	}
	var visitorIds = make([]string, 0)
	var ids = make([]string, 0)
	for _, rel := range relations {
		if rel.Visitor {
			visitorIds = append(visitorIds, rel.AccountID)
		} else {
			ids = append(ids, rel.AccountID)
		}
	}
	var mp map[string]dtos.User
	mp, err = GetUserByIds(ctx, ids)
	if err != nil {
		return nil, 0, err
	}
	var visitorMap map[string]dtos.User
	visitorMap, err = GetVisitorByIds(ctx, visitorIds)
	if err != nil {
		return nil, 0, err
	}
	var list = make([]dtos.User, 0)
	for _, rel := range relations {
		var u dtos.User
		if rel.Visitor {
			u = visitorMap[rel.AccountID]
		} else {
			u = mp[rel.AccountID]
		}
		u.PermissionSendMessage = rel.PermissionSendMessage
		list = append(list, u)
	}
	return list, total, nil
}

func GetUserRelation(ctx context.Context, bizId string, bizType constant.BizType, accountId string) (dtos.LiveUser, bool, error) {
	rel, err := dbclient.GetDBClient().GetUserRelation(bizId, bizType, accountId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return dtos.LiveUser{}, false, nil
		}
		return dtos.LiveUser{}, false, err
	}
	var lu = dtos.LiveUser{
		LiveRoomID:            rel.BizID,
		AccountID:             rel.AccountID,
		PermissionSendMessage: rel.PermissionSendMessage,
		JoinTimestamp:         rel.JoinTimestamp,
		Visitor:               rel.Visitor,
	}
	return lu, true, nil
}

func DeleteUserRelation(ctx context.Context, bizId string, bizType constant.BizType, accountId string) error {
	return dbclient.GetDBClient().DeleteUserRelation(bizId, bizType, accountId)
}

func UpdateLiveRoomUserSendMessagePermission(ctx context.Context, bizId string, bizType constant.BizType, accountId string, permission bool) error {
	return dbclient.GetDBClient().UpdateUserSendMessagePermission(bizId, bizType, accountId, permission)
}
