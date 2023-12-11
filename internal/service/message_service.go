// internal/service/message_service.go

package service

import (
	"zcloud-bg/internal/model"
	"zcloud-bg/internal/repository"
)

type MessageService struct {
	MessageRepo *repository.MessageRepository
}

// NewMessageService creates a new instance of MessageService
func NewMessageService(messageRepo *repository.MessageRepository) *MessageService {
	return &MessageService{
		MessageRepo: messageRepo,
	}
}

// CreateMessage handles the creation of a new message
func (s *MessageService) CreateMessage(message *model.Message) error {
	return s.MessageRepo.CreateMessage(message)
}

// GetMessageByID retrieves a message by its ID
func (s *MessageService) GetMessageByID(messageID string) (*model.Message, error) {
	message, err := s.MessageRepo.GetMessageByID(messageID)
	if err != nil {
		return nil, err
	}

	// Additional business logic can be added here if necessary

	return message, nil
}

// UpdateMessage updates an existing message
func (s *MessageService) UpdateMessage(message *model.Message) error {
	return s.MessageRepo.UpdateMessage(message)
}

// DeleteMessage deletes a message by its ID
func (s *MessageService) DeleteMessage(messageID string) error {
	return s.MessageRepo.DeleteMessage(messageID)
}

// ListMessages lists messages based on filtering parameters
func (s *MessageService) ListMessages(conversationID string) ([]model.Message, error) {
	return s.MessageRepo.ListMessages(conversationID)
}
