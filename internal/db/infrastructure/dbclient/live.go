package dbclient

import (
	"github.com/3115826227/go-web-live/internal/constant"
	"github.com/3115826227/go-web-live/internal/db/tables"
	"github.com/3115826227/go-web-live/internal/errors"
	"gorm.io/gorm"
)

func addLiveRoom(c *Client, liveRoom tables.LiveRoom) errors.Error {
	if err := c.client.CreateObject(&liveRoom); err != nil {
		return errors.NewCommonError(errors.CodeInternalError)
	}
	return nil
}

func getLiveRooms(c *Client, page, pageSize int64, accountId string) ([]tables.LiveRoom, int64, errors.Error) {
	var (
		liveRooms []tables.LiveRoom
		total     int64
		err       error
		offset    = int((page - 1) * pageSize)
		limit     = int(pageSize)
	)
	var template = c.client.GetDB().Model(&tables.LiveRoom{})
	if err = template.Count(&total).Error; err != nil {
		return nil, 0, errors.NewCommonError(errors.CodeInternalError)
	}
	if err = template.Offset(offset).Limit(limit).Find(&liveRooms).Order("id desc").Error; err != nil {
		return nil, 0, errors.NewCommonError(errors.CodeInternalError)
	}
	return liveRooms, total, nil
}

func getLiveRoomById(c *Client, liveRoomId string) (tables.LiveRoom, errors.Error) {
	var liveRoom tables.LiveRoom
	if err := c.client.GetDB().Where("id = ?", liveRoomId).First(&liveRoom).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return tables.LiveRoom{}, errors.NewCommonError(errors.CodeLiveRoomIdNotExistError)
		}
		return tables.LiveRoom{}, errors.NewCommonError(errors.CodeInternalError)
	}
	return liveRoom, nil
}

func getLiveRoomByAccountId(c *Client, accountId string) (tables.LiveRoom, errors.Error) {
	var liveRoom tables.LiveRoom
	if err := c.client.GetDB().Where("account_id = ?", accountId).First(&liveRoom).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return tables.LiveRoom{}, errors.NewCommonError(errors.CodeLiveRoomIdNotExistError)
		}
		return tables.LiveRoom{}, errors.NewCommonError(errors.CodeInternalError)
	}
	return liveRoom, nil
}

func updateLiveRoomStatus(c *Client, accountId string, status constant.LiveRoomStatus) errors.Error {
	if err := c.client.GetDB().Model(&tables.LiveRoom{}).
		Where("account_id = ?", accountId).
		Updates(map[string]interface{}{"status": status}).Error; err != nil {
		return errors.NewCommonError(errors.CodeInternalError)
	}
	return nil
}

func updateLiveRoomSendMessagePermission(c *Client, accountId string, permission bool) errors.Error {
	if err := c.client.GetDB().Model(&tables.LiveRoom{}).Where("account_id = ?",
		accountId).Updates(map[string]interface{}{
		"permission_send_message": permission,
	}).Error; err != nil {
		return errors.NewCommonError(errors.CodeInternalError)
	}
	return nil
}
