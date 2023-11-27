package model

import (
	"gorm.io/gorm"
	"time"
)

// User 对应于 users 数据表
type User struct {
	UserID    string         `gorm:"type:char(36);primaryKey"`
	Username  string         `gorm:"type:varchar(50);unique"`
	Password  string         `gorm:"type:varchar(255)"`
	Phone     string         `gorm:"type:varchar(20)"`
	Email     string         `gorm:"type:varchar(100);unique"`
	CreatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp"`
	Status    string         `gorm:"type:enum('active', 'inactive', 'suspended');default:'active'"`
}
