// internal/handler/organization_handler.go

package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zcloud-bg/internal/service"
	"fmt"
)

type OrganizationHandler struct {
	OrganizationService *service.OrganizationService
	OrganizationMemberService *service.OrganizationMemberService
	EmployeeService *service.EmployeeService
}

func NewOrganizationHandler(organizationService *service.OrganizationService, organizationmemberService *service.OrganizationMemberService, employeeService *service.EmployeeService) *OrganizationHandler {
	return &OrganizationHandler{
		OrganizationService: organizationService,
		OrganizationMemberService: organizationmemberService,
		EmployeeService: employeeService,
	}
}

func (h *OrganizationHandler) GetOrganization(c *gin.Context) {
	organizationID, exists := c.Get("OrganizationID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Organization ID not found"})
		return
	}

	organization, err := h.OrganizationService.GetOrganizationByID(organizationID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving organization"})
		return
	}

	c.JSON(http.StatusOK, organization)
}

func (h *OrganizationHandler) UpdateOrganizationName(c *gin.Context) {
	organizationID, exists := c.Get("OrganizationID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Organization ID not found"})
		return
	}
	userID, exists := c.Get("UserID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}
	managers, err := h.OrganizationMemberService.ListManagers(organizationID.(string), userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving organization"})
		return
	}
	var userIsManager bool
	for _, manager := range managers {
		if manager.UserID == userID.(string) {
			userIsManager = true
			fmt.Println("userIsManager")
			break
		}
	}
	if !userIsManager {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not authorized"})
		return
	}

	var requestBody struct {
		NewName string `json:"newName"`
	}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = h.OrganizationService.UpdateOrganizationName(organizationID.(string), requestBody.NewName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating organization name"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Organization name updated successfully"})
}

func (h *OrganizationHandler) GetOrganizationData(c *gin.Context) {
	organizationID, exists := c.Get("OrganizationID")
	fmt.Print(organizationID)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Organization ID not found"})
		return
	}
	userID, exists := c.Get("UserID")
	fmt.Print(userID)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}
	validmember, err := h.OrganizationMemberService.GetMember(organizationID.(string), userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving organization"})
		return
	}

	if validmember == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "User is not a member of the organization"})
		return
	}

	organizationmember, err := h.OrganizationMemberService.ListMembers(organizationID.(string), userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving organization"})
		return
	}
	employees, err := h.EmployeeService.ListEmployees(organizationID.(string), userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving organization"})
		return
	}
	managers, err  := h.OrganizationMemberService.ListManagers(organizationID.(string), userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving organization"})
		return
	}
	numofmembers := len(organizationmember)
	numofemployees := len(employees)
	numofmanagers := len(managers)
	c.JSON(http.StatusOK, gin.H{
		"OrganizationID":              organizationID,
		"MemberCount":            numofmembers,
		"EmployeeCount":            numofemployees,
		"ManagerCount":            numofmanagers,
	})
}
