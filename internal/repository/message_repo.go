// internal/repository/message_repo.go

package repository

import (
	"gorm.io/gorm"
	"time"
	"zcloud-bg/internal/model"
)

type MessageRepository struct {
	DB *gorm.DB
}

// CreateMessage creates a new message record.
func (repo *MessageRepository) CreateMessage(message *model.Message) error {
	result := repo.DB.Create(message)
	return result.Error
}

// GetMessageByID retrieves a message by its ID.
func (repo *MessageRepository) GetMessageByID(id string) (*model.Message, error) {
	var message model.Message
	result := repo.DB.Where("message_id = ?", id).First(&message)
	if result.Error != nil {
		return nil, result.Error
	}
	return &message, nil
}

// UpdateMessage updates a given message.
func (repo *MessageRepository) UpdateMessage(message *model.Message) error {
	result := repo.DB.Model(&model.Message{}).Where("message_id = ?", message.MessageID).Updates(message)
	return result.Error
}

// DeleteMessage deletes a message by its ID.
func (repo *MessageRepository) DeleteMessage(id string) error {
	result := repo.DB.Model(&model.Message{}).Where("message_id = ?", id).Update("deleted_at", time.Now())
	return result.Error
}

// ListMessages lists messages filtered by conversation ID, sender ID, and other criteria.
func (repo *MessageRepository) ListMessages(conversationID string) ([]model.Message, error) {
	var messages []model.Message
	query := repo.DB.Order("created_at ASC") // Default sort by creation time in descending order

	// Build query based on provided filtering criteria
	if conversationID != "" {
		query = query.Where("conversation_id = ?", conversationID)
	}

	// Execute the query
	result := query.Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}

	return messages, nil
}
