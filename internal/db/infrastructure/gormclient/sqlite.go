package gormclient

import (
	"errors"
	"github.com/3115826227/go-web-live/internal/db/interfaces"
	"github.com/3115826227/go-web-live/internal/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"path/filepath"
)

type GormClient struct {
	pool *gorm.DB
}

type GormWriter struct {
	lc log.Logging
}

func (gl *GormWriter) Printf(s string, values ...interface{}) {
	gl.lc.Debugf(s, values...)
}

func (c *GormClient) Close() {
}

func NewSQLiteClient(dataSource string, lc log.Logging) (interfaces.GORMClient, error) {
	var gw = &GormWriter{
		lc: lc,
	}
	// 获取配置文件，创建相应的目录
	dataSourceDir := filepath.Dir(dataSource)
	_, fileErr := os.Stat(dataSourceDir)
	if fileErr != nil || !os.IsExist(fileErr) {
		_ = os.MkdirAll(dataSourceDir, os.ModePerm)
	}
	pool, err := gorm.Open(sqlite.Open(dataSource), &gorm.Config{
		Logger: logger.New(gw, logger.Config{}),
		// issue: constraints not implemented on sqlite, consider using DisableForeignKeyConstraintWhenMigrating, more details https://gorm.io/zh_CN/docs/constraints.html#%E5%A4%96%E9%94%AE%E7%BA%A6%E6%9D%9F
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		lc.Error(err.Error())
		return nil, err
	}
	lc.Infof("sqlite open started dataSource: %s", dataSource)
	pool = pool.Debug()
	return &GormClient{pool: pool}, nil
}

func (c *GormClient) InitTable(dos ...interfaces.DataObject) error {
	var tables = make([]interface{}, 0)
	for _, do := range dos {
		tables = append(tables, do)
	}
	return c.pool.AutoMigrate(tables...)
}

func (c *GormClient) GetDB() *gorm.DB {
	return c.pool
}

// 添加
func (c *GormClient) CreateObject(object interfaces.DataObject) (err error) {
	return c.pool.Table(object.TableName()).Create(object).Error
}

// 删除
func (c *GormClient) DeleteObject(object interfaces.DataObject) (err error) {
	return c.pool.Table(object.TableName()).Delete(object).Error
}

// 获取结果
func (c *GormClient) GetObject(obCond interfaces.DataObject, object interfaces.DataObject) (err error) {
	return c.pool.Table(object.TableName()).Where(obCond).First(object).Error
}

// 更新数据
func (c *GormClient) UpdateObject(object interfaces.DataObject) (err error) {
	return c.pool.Table(object.TableName()).Select("*").Updates(object).Error
}

// 判断是否存在
func (c *GormClient) ExistObject(do interfaces.DataObject) (exist bool, err error) {
	var count int64
	template := c.pool.Table(do.TableName())
	err = template.Where(do).Count(&count).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
