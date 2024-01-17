// internal/handler/user_handler.go

package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zcloud-bg/internal/service"
)

type EmployeeHandler struct {
	EmployeeService *service.EmployeeService
}

func NewEmployeeHandler(employeeService *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		EmployeeService: employeeService,
	}
}

func (h *EmployeeHandler) ListEmployees(c *gin.Context) {
	organizationID := c.Param("organizationID")
    if organizationID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Organization ID not provided"})
        return
    }

    userID, exists := c.Get("UserID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}

    employees, err := h.EmployeeService.ListEmployees(organizationID, userID.(string))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving employee"})
        return
    }

    // response := []map[string]interface{}{}
    c.JSON(http.StatusOK, employees)
}
