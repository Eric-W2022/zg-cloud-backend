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

func (h *EmployeeHandler) GetEmployee(c *gin.Context) {
	organizationID := c.Param("organizationID")
    if organizationID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Organization ID not provided"})
        return
    }

    employee, err := h.EmployeeService.GetEmployeeByOrgID(organizationID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving employee"})
        return
    }

    c.JSON(http.StatusOK, employee)
}
