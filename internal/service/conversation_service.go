// internal/service/conversation_service.go

package service

import (
	"zcloud-bg/internal/model"
	"zcloud-bg/internal/repository"
)

type ConversationService struct {
	ConversationRepo *repository.ConversationRepository
}

func NewConversationService(conversationRepo *repository.ConversationRepository) *ConversationService {
	return &ConversationService{
		ConversationRepo: conversationRepo,
	}
}

func (s *ConversationService) CreateConversation(conversation *model.Conversation) error {
	// Create a new conversation
	return s.ConversationRepo.CreateConversation(conversation)
}

func (s *ConversationService) GetConversationByID(conversationID string) (*model.Conversation, error) {
	// Retrieve a conversation by its ID
	conversation, err := s.ConversationRepo.GetConversationByID(conversationID)
	if err != nil {
		return nil, err
	}

	// Additional business logic can be added here if necessary

	return conversation, nil
}

func (s *ConversationService) UpdateConversation(conversation *model.Conversation) error {
	// Update an existing conversation
	return s.ConversationRepo.UpdateConversation(conversation)
}

func (s *ConversationService) DeleteConversation(conversationID string) error {
	// Delete a conversation by its ID
	return s.ConversationRepo.DeleteConversation(conversationID)
}

// ListConversations 方法接受筛选参数
func (s *ConversationService) ListConversations(organizationID, platformID, userID string) ([]model.Conversation, error) {
	// 调用 ConversationRepo 的 ListConversations 方法，并传递筛选参数
	return s.ConversationRepo.ListConversations(organizationID, platformID, userID)
}
