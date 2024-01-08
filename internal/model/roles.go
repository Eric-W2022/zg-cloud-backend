// internal/model/role.go

package model


// Role represents a role in the application.
type Role struct {
	RoleID      string    `gorm:"type:char(36);primaryKey"`
	RoleName    string    `gorm:"type:varchar(255)"`
	Permissions string    `gorm:"type:json"`
}