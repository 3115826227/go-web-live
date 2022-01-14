package interfaces

import "gorm.io/gorm"

type GORMClient interface {
	// 关闭连接
	Close()
	// 获取db
	GetDB() *gorm.DB
	// 初始化建表操作
	InitTable(dos ...DataObject) error
	// 添加数据
	CreateObject(do DataObject) error
	// 判断数据是否存在
	ExistObject(do DataObject) (bool, error)
	// 删除数据
	DeleteObject(do DataObject) error
	// 更新数据
	UpdateObject(do DataObject) error
}

type DataObject interface {
	TableName() string
}
