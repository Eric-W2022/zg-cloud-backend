// pkg/jwt/middleware.go

package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"zcloud-bg/internal/service" // 根据你的项目结构调整
)

// AuthMiddleware 创建一个 JWT 验证中间件
func AuthMiddleware(jwtKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 提取 JWT 令牌逻辑...
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		// 移除前缀 "Bearer "
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 解析令牌
		claims := &service.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		//fmt.Print(claims)

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// 设置用户名到 Gin 上下文
		c.Set("username", claims.Username)
		c.Set("UserID", claims.UserID)
		c.Set("OrganizationID", claims.OrganizationID)
		c.Next()
	}
}
