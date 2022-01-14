package tables

// 会话信息表
type Session struct {
	CommonField
	// 会话创建者AccountId
	AccountId string `gorm:"column:account_id;unique"`
	// 会话名称
	Name string `gorm:"column:name"`
	// 描述
	Description string `gorm:"column:description"`
	// 是否禁言
	PermissionSendMessage bool `gorm:"column:permission_send_message"`
	// 举报次数
	Reports int64 `gorm:"reports"`
}

func (table *Session) TableName() string {
	return "wl_session"
}
