// internal/repository/conversation_repo.go

package repository

import (
	"gorm.io/gorm"
	"time"
	"zcloud-bg/internal/model"
)

type ConversationRepository struct {
	DB *gorm.DB
}

// CreateConversation creates a new conversation record.
func (repo *ConversationRepository) CreateConversation(conversation *model.Conversation) error {
	result := repo.DB.Create(conversation)
	return result.Error
}

// GetConversationByID retrieves a conversation by its ID.
func (repo *ConversationRepository) GetConversationByID(id string) (*model.Conversation, error) {
	var conversation model.Conversation
	result := repo.DB.Where("conversation_id = ?", id).First(&conversation)
	if result.Error != nil {
		return nil, result.Error
	}
	return &conversation, nil
}

// UpdateConversation updates a given conversation.
func (repo *ConversationRepository) UpdateConversation(conversation *model.Conversation) error {
	result := repo.DB.Save(conversation)
	return result.Error
}

// DeleteConversation deletes a conversation by its ID.
func (repo *ConversationRepository) DeleteConversation(id string) error {
	// Soft delete the conversation with the given ID
	result := repo.DB.Model(&model.Conversation{}).Where("conversation_id = ?", id).Update("deleted_at", time.Now())
	return result.Error
}

// ListConversations lists all conversations.
func (repo *ConversationRepository) ListConversations(organizationID, platformID, userID string) ([]model.Conversation, error) {
	var conversations []model.Conversation
	query := repo.DB.Order("created_at DESC") // 默认按创建时间降序排序

	// 根据提供的筛选条件构建查询
	if organizationID != "" {
		query = query.Where("organization_id = ?", organizationID)
	}
	if platformID != "" {
		query = query.Where("platform_id = ?", platformID)
	}
	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}

	// 执行查询
	result := query.Find(&conversations)
	if result.Error != nil {
		return nil, result.Error
	}

	return conversations, nil
}
