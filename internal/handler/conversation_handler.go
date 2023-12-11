// internal/handler/conversation_handler.go

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"zcloud-bg/internal/model"
	"zcloud-bg/internal/service"
)

type ConversationHandler struct {
	ConversationService *service.ConversationService
}

func NewConversationHandler(conversationService *service.ConversationService) *ConversationHandler {
	return &ConversationHandler{
		ConversationService: conversationService,
	}
}

func (h *ConversationHandler) GetConversation(c *gin.Context) {
	conversationID := c.Param("conversationID")

	if conversationID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Conversation ID not provided"})
		return
	}

	conversation, err := h.ConversationService.GetConversationByID(conversationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving conversation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ConversationID": conversation.ConversationID,
		"OrganizationID": conversation.OrganizationID,
		"PlatformID":     conversation.PlatformID,
		"Title":          conversation.Title,
	})
}

func (h *ConversationHandler) UpdateConversation(c *gin.Context) {
	conversationID := c.Param("conversationID")
	if conversationID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Conversation ID not provided"})
		return
	}

	// Validate that conversationID is a valid UUID
	if _, err := uuid.Parse(conversationID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Conversation ID format"})
		return
	}

	// Define a structure to receive optional data from the request body
	type updateRequest struct {
		Title       *string  `json:"title"`
		Temperature *float32 `json:"temperature"`
		Model       *string  `json:"model"`
	}

	var req updateRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	conversationUpdate := model.Conversation{
		ConversationID: conversationID,
	}

	// Update title if provided
	if req.Title != nil {
		if len(*req.Title) > 50 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Title exceeds maximum length of 50 characters"})
			return
		}
		conversationUpdate.Title = *req.Title
	}

	// Update temperature if provided
	if req.Temperature != nil {
		conversationUpdate.Temperature = *req.Temperature
	}

	// Update model if provided
	if req.Model != nil {
		if len(*req.Model) > 100 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Model exceeds maximum length of 100 characters"})
			return
		}
		conversationUpdate.Model = *req.Model
	}

	// Fetch organizationID from context
	rawOrganizationID, orgExists := c.Get("OrganizationID")
	if !orgExists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "OrganizationID not found in context"})
		return
	}

	organizationID, ok := rawOrganizationID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "OrganizationID has an invalid type"})
		return
	}

	conversationUpdate.OrganizationID = organizationID

	// Proceed with updating the conversation
	err := h.ConversationService.UpdateConversation(&conversationUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating conversation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Conversation updated successfully"})
}

func (h *ConversationHandler) DeleteConversation(c *gin.Context) {
	conversationID := c.Param("conversationID")
	if conversationID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Conversation ID not provided"})
		return
	}

	err := h.ConversationService.DeleteConversation(conversationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting conversation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Conversation deleted successfully"})
}

func (h *ConversationHandler) CreateConversation(c *gin.Context) {
	var newConversation model.Conversation

	// 从请求中获取 title 参数
	title := c.Query("title") // 如果参数不存在，将得到空字符串

	// 解析 UserID
	userID, exists := c.Get("UserID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UserID not found in context"})
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UserID is not a string"})
		return
	}

	// 解析 OrganizationID
	organizationID, orgExists := c.Get("OrganizationID")
	if !orgExists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "OrganizationID not found in context"})
		return
	}

	organizationIDStr, orgOK := organizationID.(string)
	if !orgOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "OrganizationID is not a string"})
		return
	}

	// 默认温度 0
	newConversation.Temperature = 0

	// 如果有温度则校验
	temperatureStr := c.Query("temperature")
	if temperatureStr != "" {
		var err error
		temperature, err := strconv.ParseFloat(temperatureStr, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid temperature format"})
			return
		}
		newConversation.Temperature = float32(temperature)
	}

	// 默认模型设置空
	newConversation.Model = ""

	// 有模型则校验
	modelName := c.Query("modelName")
	if len(modelName) > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Model exceeds maximum length of 100 characters"})
		return
	} else if modelName != "" {
		newConversation.Model = modelName
	}

	newConversation.UserID = userIDStr
	newConversation.OrganizationID = organizationIDStr
	newConversation.ConversationID = uuid.New().String()
	newConversation.PlatformID = "7417dcfc-4508-4030-b709-304c7c5404dd"

	// 设置 Title 字段
	if title != "" && len(title) < 50 {
		newConversation.Title = title
	} else {
		newConversation.Title = "新对话"
	}

	err := h.ConversationService.CreateConversation(&newConversation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating conversation"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":         "Conversation created successfully",
		"conversation_id": newConversation.ConversationID,
	})
}

func (h *ConversationHandler) ListConversations(c *gin.Context) {
	// 从上下文中解析 OrganizationID
	organizationID, _ := c.Get("OrganizationID") // 假设 OrganizationID 总是存在

	// 从请求参数中解析 PlatformID 和 UserID（这些参数是可选的）
	platformID := c.Query("platformID") // 如果参数不存在，将得到空字符串
	userID := c.Query("userID")         // 如果参数不存在，将得到空字符串

	// 调用服务层方法，传递筛选参数
	conversations, err := h.ConversationService.ListConversations(organizationID.(string), platformID, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving conversations"})
		return
	}

	c.JSON(http.StatusOK, conversations)
}
