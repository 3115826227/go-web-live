package dbclient

import (
	"github.com/3115826227/go-web-live/internal/constant"
	"github.com/3115826227/go-web-live/internal/db/tables"
	"github.com/3115826227/go-web-live/internal/errors"
	"gorm.io/gorm"
)

func addMessage(c *Client, msg tables.Message) error {
	return c.client.CreateObject(&msg)
}

func getMessages(c *Client, bizId string, bizType constant.BizType, timestamp int64, page, pageSize int64) ([]tables.Message, errors.Error) {
	var (
		messages []tables.Message
		offset   = int((page - 1) * pageSize)
		limit    = int(pageSize)
	)
	var template = c.client.GetDB().Model(&tables.Message{}).Where("biz_id = ? and biz_type = ? and send_timestamp > ?", bizId, bizType, timestamp)
	err := template.Offset(offset).Limit(limit).Order("id desc").Find(&messages).Error
	if err != nil {
		return nil, errors.NewCommonError(errors.CodeInternalError)
	}
	return messages, nil
}

func getUserRelationTotal(c *Client, bizId string, bizType constant.BizType) (int64, errors.Error) {
	var template = c.client.GetDB().Model(&tables.UserRelation{}).
		Where("biz_id = ? and biz_type = ?", bizId, bizType)
	var total int64
	if err := template.Count(&total).Error; err != nil {
		return 0, errors.NewCommonError(errors.CodeInternalError)
	}
	return total, nil
}

func getUserRelations(c *Client, bizId string, bizType constant.BizType, page, pageSize int64) ([]tables.UserRelation, int64, errors.Error) {
	var (
		relations []tables.UserRelation
		total     int64
		err       error
		offset    = int((page - 1) * pageSize)
		limit     = int(pageSize)
	)
	var template = c.client.GetDB().Model(&tables.UserRelation{}).
		Where("biz_id = ? and biz_type", bizId, bizType)
	if err = template.Count(&total).Error; err != nil {
		return nil, 0, errors.NewCommonError(errors.CodeInternalError)
	}
	if err = template.Offset(offset).Limit(limit).Find(&relations).Order("id desc").Error; err != nil {
		return nil, 0, errors.NewCommonError(errors.CodeInternalError)
	}
	return relations, total, nil
}

func addUserRelation(c *Client, relation tables.UserRelation) error {
	_, err := getUserRelation(c, relation.BizID, relation.BizType, relation.AccountID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.client.CreateObject(&relation)
		}
		return err
	}
	return nil
}

func deleteUserRelation(c *Client, bizId string, bizType constant.BizType, accountId string) error {
	return c.client.GetDB().Where("biz_id = ? and biz_type = ? and account_id = ?", bizId, bizType, accountId).Delete(&tables.UserRelation{}).Error
}

func getUserRelation(c *Client, bizId string, bizType constant.BizType, accountId string) (relation tables.UserRelation, err error) {
	err = c.client.GetDB().Where("biz_id = ? and biz_type = ? and account_id = ?", bizId, bizType, accountId).First(&relation).Error
	return
}

func updateUserSendMessagePermission(c *Client, bizId string, bizType constant.BizType, accountId string, permission bool) error {
	return c.client.GetDB().Model(&tables.UserRelation{}).Where("biz_id = ? and biz_type = ? and account_id = ?", bizId, bizType, accountId).
		Updates(map[string]interface{}{
			"permission_send_message": permission,
		}).Error
}
