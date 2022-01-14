package tables

type CommonField struct {
	ID              string `gorm:"column:id;type:char(36);pk;not null"`
	CreateTimestamp int64  `gorm:"column:create_timestamp"`
	UpdateTimestamp int64  `gorm:"column:update_timestamp"`
}

type CommonIntField struct {
	ID              int64 `gorm:"column:id;AUTO_INCREMENT;"`
	CreateTimestamp int64 `gorm:"column:create_timestamp"`
	UpdateTimestamp int64 `gorm:"column:update_timestamp"`
}
