// pkg/db/database.go

package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDatabase 初始化并返回 GORM 数据库连接
func InitDatabase(mysqlUsername, mysqlPassword, mysqlAddress string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/zg-cloud?charset=utf8mb4&parseTime=True&loc=Local", mysqlUsername, mysqlPassword, mysqlAddress)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
