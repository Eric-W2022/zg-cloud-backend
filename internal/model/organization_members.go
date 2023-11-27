package model

import (
	"time"
)

// OrganizationMember 对应于 organization_members 数据表
type OrganizationMember struct {
	OrganizationID string    `gorm:"type:char(36);primaryKey;column:organization_id;foreignKey;references:OrganizationID"`
	UserID         string    `gorm:"type:char(36);primaryKey;column:user_id;foreignKey;references:UserID"`
	RoleID         string    `gorm:"type:char(36);column:role_id;default:''"`
	JoinedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:joined_at"`
	Status         string    `gorm:"type:enum('invited', 'joined', 'declined');default:'invited';column:status"`
}
