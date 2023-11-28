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
