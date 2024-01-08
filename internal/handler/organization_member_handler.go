package handler
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zcloud-bg/internal/service"
	"fmt"
)
type OrganizationMemberHandler struct {
    OrganizationMemberService *service.OrganizationMemberService
    RoleService               *service.RoleService // Add RoleService as a field
	UserService               *service.UserService 
}

func NewOrganizationMemberHandler(organizationmemberService *service.OrganizationMemberService, roleService *service.RoleService, userService *service.UserService) *OrganizationMemberHandler {
    return &OrganizationMemberHandler{
        OrganizationMemberService: organizationmemberService,
        RoleService:               roleService, // Initialize RoleService
		UserService:               userService, 
    }
}

func (h *OrganizationMemberHandler) ListMembers(c *gin.Context) {
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

	organizationmember, err := h.OrganizationMemberService.ListMembers(organizationID.(string), userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving organization"})
		return
	}
	fmt.Print("hi")
	response := []map[string]interface{}{}

	for i := range organizationmember {
		role, err := h.RoleService.GetRoleByID(organizationmember[i].RoleID) // Access RoleService from the struct
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving role"})
			return
		}
		user, err := h.UserService.GetUserByID(organizationmember[i].UserID) // Access RoleService from the struct
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user"})
			return
		}
	
		// Create a map for the member with RoleName field
		member := map[string]interface{}{
			"OrganizationID": organizationmember[i].OrganizationID,
			"UserID":         organizationmember[i].UserID,
			"UserName":         user.Username,
			"Phone":         user.Phone,
			"RoleID":         organizationmember[i].RoleID,
			"JoinedAt":       organizationmember[i].JoinedAt,
			"Status":         organizationmember[i].Status,
			"RoleName":       role.RoleName, // Add the RoleName field
		}
	
		// Append the modified member to the response slice
		response = append(response, member)
	}

	c.JSON(http.StatusOK, response)
	}