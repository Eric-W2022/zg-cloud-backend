// internal/handler/auth_handler.go

package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"zcloud-bg/internal/service"
)

// AuthHandler 结构体包含了需要的服务
type AuthHandler struct {
	AuthService *service.AuthService
	UserService *service.UserService // 这需要被初始化
	JwtKey      []byte
}

// 修改构造函数以接收 UserService
func NewAuthHandler(authService *service.AuthService, userService *service.UserService, jwtKey []byte) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
		UserService: userService, // 设置 UserService
		JwtKey:      jwtKey,
	}
}

// Login 处理登录请求
func (h *AuthHandler) Login(c *gin.Context) {
	var creds struct {
		Username       string  `json:"username"`
		Password       string  `json:"password"`
		OrganizationID *string `json:"organization_id,omitempty"` // 可选字段
	}

	// 尝试绑定请求体到 creds 结构体
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求格式不正确"})
		return
	}

	// 使用 AuthService 来验证用户凭据
	user, err := h.AuthService.AuthenticateUser(creds.Username, creds.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "登录失败"})
		return
	}

	// 判断是否有多企业，如果有的话让用户选择企业登录，如果只有一个企业就直接登录
	userInfo, err := h.UserService.GetUserByID(user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user"})
		return
	}

	// 如果提供了组织 ID，验证用户是否属于该组织
	if creds.OrganizationID != nil {
		found := false
		for _, org := range userInfo.Organizations {
			if *creds.OrganizationID == org.OrganizationID {
				found = true
				break
			}
		}
		if !found {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不属于指定的组织"})
			return
		}
	}

	// 创建 JWT 令牌
	expirationTime := time.Now().Add(72 * time.Hour)
	claims := &service.Claims{
		Username: user.Username,
		UserID:   user.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// 如果提供了组织 ID，将其设置为 JWT Claims 中的 OrganizationID
	if creds.OrganizationID != nil {
		claims.OrganizationID = *creds.OrganizationID
	} else if len(userInfo.Organizations) > 0 {
		// 如果没有提供组织 ID 但用户属于至少一个组织，使用第一个组织的 ID
		claims.OrganizationID = userInfo.Organizations[0].OrganizationID
	} else {
		// 用户不属于任何组织
		claims.OrganizationID = ""
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(h.JwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法生成令牌"})
		return
	}

	// 返回令牌给客户端
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
