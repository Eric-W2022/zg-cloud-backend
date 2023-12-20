package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDatabase 初始化并返回 GORM 数据库连接
func InitDatabase(mysqlUsername, mysqlPassword, mysqlAddress, mysqlDatabase string) (*gorm.DB, error) {
	// 使用提供的数据库名构造 DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlUsername, mysqlPassword, mysqlAddress, mysqlDatabase)

	// 使用 GORM 打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
