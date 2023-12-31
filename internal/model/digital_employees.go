// internal/model/user.go

package model

import (
	"time"
)

// Employee 对应于 digital_employee 数据表
type Digital_Employee struct {
	EmployeeID     string         `gorm:"primaryKey;type:char(36)"`
	Name           string         `gorm:"type:varchar(100)"`
	Type           string         `gorm:"type:varchar(100)"`
	Department     *string        `gorm:"type:varchar(100)"` 
	Salary         *float64       `gorm:"type:decimal(10,2)"`
	Status         string         `gorm:"type:varchar(50)"`
	PersonalBio    *string        `gorm:"type:text"` 
	OrganizationID string         `gorm:"type:char(36)"`
	CreatedAt      time.Time      `gorm:"type:datetime"`
	UpdatedAt      time.Time      `gorm:"type:datetime"`
	DeletedAt      *time.Time     `gorm:"type:datetime"` 
	NumOfService   int            `gorm:"type:int"`
	NumOfError     int            `gorm:"type:int"`
	AvatarURL      *string        `gorm:"type:varchar(255)"` 
}