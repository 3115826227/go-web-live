package dbclient

import (
	"github.com/3115826227/go-web-live/internal/constant"
	"github.com/3115826227/go-web-live/internal/db/tables"
)

func addLiveRoom(c *Client, liveRoom tables.LiveRoom) error {
	return c.client.CreateObject(&liveRoom)
}

func getLiveRooms(c *Client, page, pageSize int64, accountId string) (liveRooms []tables.LiveRoom, total int64, err error) {
	var (
		offset = int((page - 1) * pageSize)
		limit  = int(pageSize)
	)
	var template = c.client.GetDB().Model(&tables.LiveRoom{})
	if err = template.Count(&total).Error; err != nil {
		return
	}
	err = template.Offset(offset).Limit(limit).Find(&liveRooms).Order("id desc").Error
	return
}

func getLiveRoomById(c *Client, liveRoomId string) (liveRoom tables.LiveRoom, err error) {
	err = c.client.GetDB().Where("id = ?", liveRoomId).First(&liveRoom).Error
	return
}

func getLiveRoomByAccountId(c *Client, accountId string) (liveRoom tables.LiveRoom, err error) {
	err = c.client.GetDB().Where("account_id = ?", accountId).First(&liveRoom).Error
	return
}

func updateLiveRoomStatus(c *Client, accountId string, status constant.LiveRoomStatus) error {
	return c.client.GetDB().Model(&tables.LiveRoom{}).Where("account_id = ?", accountId).Updates(map[string]interface{}{"status": status}).Error
}

func updateLiveRoomSendMessagePermission(c *Client, accountId string, permission bool) error {
	return c.client.GetDB().Model(&tables.LiveRoom{}).Where("account_id = ?",
		accountId).Updates(map[string]interface{}{
		"permission_send_message": permission,
	}).Error
}
