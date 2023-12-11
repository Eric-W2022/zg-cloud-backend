// internal/model/messages.go

package model

import (
	"gorm.io/gorm"
	"time"
)

// Message 对应于 messages 数据表
type Message struct {
	MessageID      string         `gorm:"type:char(36);primaryKey;comment:消息的唯一标识符" json:"message_id"`
	ConversationID string         `gorm:"type:char(36);comment:对话的唯一标识符" json:"conversation_id"`
	SenderID       string         `gorm:"type:char(36);comment:发送者的唯一标识符" json:"sender_id"`
	Role           string         `gorm:"type:varchar(255);comment:发送者角色" json:"role"`
	Model          string         `gorm:"type:varchar(255);comment:处理消息的模型或系统" json:"model"`
	Content        string         `gorm:"type:text;comment:消息内容" json:"content"`
	CreatedAt      time.Time      `gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:消息创建时间" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"type:datetime;comment:消息最后更新时间" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"type:datetime;comment:消息删除时间" json:"deleted_at,omitempty"`
	InputTokens    int            `gorm:"type:int;comment:输入的令牌数量" json:"input_tokens"`
	OutputTokens   int            `gorm:"type:int;comment:输出的令牌数量" json:"output_tokens"`
	TotalTokens    int            `gorm:"type:int;comment:总令牌数量" json:"total_tokens"`
	Cost           *float64       `gorm:"type:decimal(10,6);comment:花费人民币" json:"cost,omitempty"`
}
