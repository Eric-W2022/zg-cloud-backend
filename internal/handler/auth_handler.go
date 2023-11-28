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
	JwtKey      []byte
}

// NewAuthHandler 创建并返回一个新的 AuthHandler
func NewAuthHandler(authService *service.AuthService, jwtKey []byte) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
		JwtKey:      jwtKey,
	}
}

// Login 处理登录请求
func (h *AuthHandler) Login(c *gin.Context) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// 尝试绑定请求体到 creds 结构体
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求格式不正确"})
		return
	}

	// 使用 AuthService 来验证用户凭据
	user, err := h.AuthService.AuthenticateUser(creds.Username, creds.Password)
	//fmt.Print(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "登录失败"})
		return
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(h.JwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法生成令牌"})
		return
	}

	// 返回令牌给客户端
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
