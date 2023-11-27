// pkg/jwt/middleware.go

package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"zcloud-bg/internal/service" // 根据你的项目结构调整
)

// AuthMiddleware 创建一个 JWT 验证中间件
func AuthMiddleware(jwtKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 提取 JWT 令牌逻辑...
		tokenString := c.GetHeader("Authorization")

		// 解析令牌
		claims := &service.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
