// internal/handler/user_handler.go

package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zcloud-bg/internal/service"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userID, exists := c.Get("UserID")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}

	user, err := h.UserService.GetUserByID(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"UserID":              user.UserID,
		"Username":            user.Username,
		"MemberOrganizations": user.MemberOrganizations,
	})
}

func (h *UserHandler) GetNameByID(c *gin.Context) {
	userID:= c.Param("UserID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not provided"})
		return
	}

	user, err := h.UserService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"UserID":              user.UserID,
		"Username":            user.Username,
		"MemberOrganizations": user.MemberOrganizations,
	})
}

func (h *UserHandler) UpdateUserName(c *gin.Context) {
	userID, exists := c.Get("UserID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}

	var requestBody struct {
		NewName string `json:"newName"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if len(requestBody.NewName) > 20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name too long"})
		return
	}

	err := h.UserService.UpdateUserName(userID.(string), requestBody.NewName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user name"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User name updated successfully"})
}
