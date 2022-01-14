package tables

import "github.com/3115826227/go-web-live/internal/constant"

// 消息用户关联表
type UserRelation struct {
	// 业务id
	BizID string `gorm:"column:biz_id;unique_index:biz_type_account"`
	// 业务类型
	BizType constant.BizType `gorm:"biz_type;unique_index:biz_type_account"`
	// 用户id
	AccountID string `gorm:"column:account_id;unique_index:biz_type_account"`
	// 是否被禁言
	PermissionSendMessage bool `gorm:"column:permission_send_message"`
	// 加入会话的时间点
	JoinTimestamp int64 `gorm:"column:join_timestamp"`
	// 是否为访客
	Visitor bool `gorm:"column:visitor"`
}

func (table *UserRelation) TableName() string {
	return "wl_user_rel"
}

// 消息表
type Message struct {
	// 消息id 自增
	ID int64 `gorm:"column:id;pk;autoIncrement"`
	// 业务id
	BizID string `gorm:"column:biz_id;"`
	// 业务类型
	BizType constant.BizType `gorm:"column:biz_type"`
	// 消息类型
	MessageType constant.MessageType `gorm:"column:message_type;not null"`
	// 发送者id
	Send string `gorm:"column:send;not null"`
	// 消息内容
	Content string `gorm:"column:content"`
	// 消息发送时间
	SendTimestamp int64 `gorm:"column:send_timestamp;not null"`
}

func (table *Message) TableName() string {
	return "wl_message"
}
