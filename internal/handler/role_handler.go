// internal/handler/Role_handler.go

package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zcloud-bg/internal/service"
)

type RoleHandler struct {
	RoleService *service.RoleService
}

func NewRoleHandler(roleService *service.RoleService) *RoleHandler {
	return &RoleHandler{
		RoleService: roleService,
	}
}

func (h *RoleHandler) GetRole(c *gin.Context) {
	roleID := c.Param("roleID")
    
    if roleID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Role ID not found"})
        return
    }

    role, err := h.RoleService.GetRoleByID(roleID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving role"})
        return
    }

    c.JSON(http.StatusOK, role)
}