// internal/handler/message_handler.go

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"math"
	"net/http"
	"time"
	"zcloud-bg/internal/model"
	"zcloud-bg/internal/service"
)

type MessageHandler struct {
	MessageService *service.MessageService
}

func NewMessageHandler(messageService *service.MessageService) *MessageHandler {
	return &MessageHandler{
		MessageService: messageService,
	}
}

// GetMessage retrieves a message by its ID
func (h *MessageHandler) GetMessage(c *gin.Context) {
	messageID := c.Param("messageID")
	if messageID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Message ID not provided"})
		return
	}

	message, err := h.MessageService.GetMessageByID(messageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving message"})
		return
	}

	c.JSON(http.StatusOK, message)
}

// UpdateMessage updates an existing message
func (h *MessageHandler) UpdateMessage(c *gin.Context) {
	messageID := c.Param("messageID")
	if messageID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Message ID not provided"})
		return
	}

	type updateRequest struct {
		Content string `json:"content"`
	}

	var req updateRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	messageUpdate := model.Message{
		MessageID: messageID,
		Content:   req.Content,
	}

	err := h.MessageService.UpdateMessage(&messageUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message updated successfully"})
}

// DeleteMessage deletes a message by its ID
func (h *MessageHandler) DeleteMessage(c *gin.Context) {
	messageID := c.Param("messageID")
	if messageID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Message ID not provided"})
		return
	}

	err := h.MessageService.DeleteMessage(messageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message deleted successfully"})
}

//MessageID      string         `gorm:"type:char(36);primaryKey;comment:消息的唯一标识符" json:"message_id"`
//ConversationID string         `gorm:"type:char(36);comment:对话的唯一标识符" json:"conversation_id"`
//SenderID       string         `gorm:"type:char(36);comment:发送者的唯一标识符" json:"sender_id"`
//Role           string         `gorm:"type:varchar(255);comment:发送者角色" json:"role"`
//Model          string         `gorm:"type:varchar(255);comment:处理消息的模型或系统" json:"model"`
//Content        string         `gorm:"type:text;comment:消息内容" json:"content"`
//CreatedAt      time.Time      `gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:消息创建时间" json:"created_at"`
//UpdatedAt      time.Time      `gorm:"type:datetime;comment:消息最后更新时间" json:"updated_at"`
//DeletedAt      gorm.DeletedAt `gorm:"type:datetime;comment:消息删除时间" json:"deleted_at,omitempty"`
//InputTokens    int            `gorm:"type:int;comment:输入的令牌数量" json:"input_tokens"`
//OutputTokens   int            `gorm:"type:int;comment:输出的令牌数量" json:"output_tokens"`
//TotalTokens    int            `gorm:"type:int;comment:总令牌数量" json:"total_tokens"`

// CreateMessage 创建新消息
func (h *MessageHandler) CreateMessage(c *gin.Context) {
	var newMessage model.Message

	// 假设我们从请求体中提取 senderID、conversationID 和 content
	if err := c.BindJSON(&newMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体"})
		return
	}

	// 解析 UserID
	userID, exists := c.Get("UserID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "上下文中找不到UserID"})
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UserID 不是字符串"})
		return
	}

	newMessage.MessageID = uuid.New().String() // 为消息生成一个新的UUID
	newMessage.SenderID = userIDStr
	newMessage.CreatedAt = time.Now() // 设置当前时间为消息创建时间

	// 根据内容长度计算 handling_time
	contentLength := len(newMessage.Content)
	var handlingTime float64

	if contentLength <= 10 {
		handlingTime = float64(contentLength) * 0.5
	} else if contentLength <= 20 {
		handlingTime = float64(contentLength) * 0.2
	} else if contentLength <= 30 {
		handlingTime = float64(contentLength) * 0.2
	} else {
		handlingTime = math.Min(float64(contentLength)*0.15, 15)
	}

	// 假设你想要以秒为单位的 handling_time
	handlingTimeInSeconds := int(handlingTime)

	err := h.MessageService.CreateMessage(&newMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建消息时出错"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":       "消息创建成功",
		"message_id":    newMessage.MessageID,
		"created_at":    newMessage.CreatedAt,
		"handling_time": handlingTimeInSeconds, // 将 handlingTime 转换为秒
	})
}

type MessageCreationResponse struct {
	MessageID string    `json:"message_id"`
	Content   string    `json:"content"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func (h *MessageHandler) CreateMultipleMessages(c *gin.Context) {
	var messages []model.Message

	// 绑定请求体到消息数组
	if err := c.BindJSON(&messages); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

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

	var creationResponses []MessageCreationResponse
	currentTime := time.Now() // 设置初始时间为当前时间

	// 循环创建每个消息
	for i := range messages {
		messages[i].MessageID = uuid.New().String() // 为每个消息生成一个新的UUID
		messages[i].SenderID = userIDStr
		messages[i].CreatedAt = currentTime.Add(time.Second * time.Duration(i)) // 设置创建时间，并递增

		err := h.MessageService.CreateMessage(&messages[i])
		if err != nil {
			// 如果创建过程中出现错误，可能需要回滚已经创建的消息
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating messages"})
			return
		}

		// 将消息ID和创建时间添加到响应数组中
		creationResponses = append(creationResponses, MessageCreationResponse{
			MessageID: messages[i].MessageID,
			CreatedAt: messages[i].CreatedAt,
			Content:   messages[i].Content,
			Role:      messages[i].Role,
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Messages created successfully",
		"data":    creationResponses,
	})
}

// ListMessages lists messages based on filtering parameters
func (h *MessageHandler) ListMessages(c *gin.Context) {
	conversationID := c.Query("conversationID")

	messages, err := h.MessageService.ListMessages(conversationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error listing messages"})
		return
	}

	c.JSON(http.StatusOK, messages)
}
