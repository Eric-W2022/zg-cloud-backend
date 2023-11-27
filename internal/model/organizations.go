package model

import (
	"time"
)

// Organization 对应于 organizations 数据表
type Organization struct {
	OrganizationID string               `gorm:"type:char(36);primaryKey;comment:组织的唯一标识符"`
	Name           string               `gorm:"type:varchar(255);comment:组织的名称"`
	CreatedBy      string               `gorm:"type:char(36);comment:创建组织的用户的ID"`
	CreatedAt      time.Time            `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:组织的创建时间"`
	Members        []OrganizationMember `gorm:"foreignKey:OrganizationID"`
}
