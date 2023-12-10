// internal/model/conversations.go

package model

import (
	"gorm.io/gorm"
	"time"
)

// Conversation 对应于 conversations 数据表
type Conversation struct {
	ConversationID string         `gorm:"type:char(36);primaryKey;comment:对话的唯一标识符" json:"conversation_id"`
	OrganizationID string         `gorm:"type:char(36);comment:组织的唯一标识符，表示对话属于哪个组织" json:"organization_id"`
	PlatformID     string         `gorm:"type:char(36);comment:平台的唯一标识符，表示对话来自哪个平台" json:"platform_id"`
	UserID         string         `gorm:"type:char(36);comment:用户的唯一标识符" json:"user_id"`
	Title          string         `gorm:"type:text;comment:对话的标题" json:"title"`
	CreatedAt      time.Time      `gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:记录的创建时间" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"type:datetime;comment:记录的最后更新时间" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"type:datetime;comment:记录的删除时间" json:"deleted_at,omitempty"`
}
