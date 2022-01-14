package tables

import "github.com/3115826227/go-web-live/internal/constant"

// 直播房间信息表
type LiveRoom struct {
	CommonField
	// 直播房间主播
	AccountId string `gorm:"column:account_id;unique"`
	// 直播房间状态
	Status constant.LiveRoomStatus `gorm:"column:status"`
	// 是否禁言
	PermissionSendMessage bool `gorm:"column:permission_send_message"`
	// 直播间最高人数
	MaxUserTotal int64 `gorm:"max_user_total"`
	// 举报次数
	Reports int64 `gorm:"reports"`
}

func (table *LiveRoom) TableName() string {
	return "wl_live_room"
}
