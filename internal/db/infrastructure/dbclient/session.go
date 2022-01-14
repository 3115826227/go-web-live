package dbclient

import (
	"github.com/3115826227/go-web-live/internal/db/tables"
)

func addSession(c *Client, session tables.Session) error {
	return c.client.CreateObject(&session)
}

func updateSession(c *Client, session tables.Session) error {
	return c.client.GetDB().Where("id = ?", session.ID).Save(&session).Error
}

func getSessions(c *Client, page, pageSize int64, accountId string) (sessions []tables.Session, total int64, err error) {
	var (
		offset = int((page - 1) * pageSize)
		limit  = int(pageSize)
	)
	var template = c.client.GetDB().Model(&tables.Session{})
	if err = template.Count(&total).Error; err != nil {
		return
	}
	err = template.Offset(offset).Limit(limit).Find(&sessions).Order("id desc").Error
	return
}

func getSessionById(c *Client, liveRoomId string) (session tables.Session, err error) {
	err = c.client.GetDB().Where("id = ?", liveRoomId).First(&session).Error
	return
}

func updateSessionSendMessagePermission(c *Client, sessionId string, permission bool) error {
	return c.client.GetDB().Model(&tables.Session{}).Where("id = ?",
		sessionId).Updates(map[string]interface{}{
		"permission_send_message": permission,
	}).Error
}
