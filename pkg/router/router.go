// router.go

package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"zcloud-bg/internal/handler"
	"zcloud-bg/internal/repository"
	"zcloud-bg/internal/service"
	"zcloud-bg/pkg/jwt"
)

// CORS 中间件设置跨域请求所需的 HTTP 头
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// Setup 初始化并返回一个配置好的 Gin 路由器
func Setup(db *gorm.DB, jwtKey []byte) *gin.Engine {
	r := gin.Default()

	// 使用 CORS 中间件
	r.Use(CORS())

	userRepo := &repository.UserRepository{DB: db}
	authService := &service.AuthService{UserRepo: userRepo}
	authHandler := handler.NewAuthHandler(authService, jwtKey)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// 公共路由
	r.POST("/login", authHandler.Login)

	// 需要身份验证的路由组
	authRoutes := r.Group("/").Use(jwt.AuthMiddleware(jwtKey))
	{
		// 添加需要验证的路由
		authRoutes.GET("/user/info", userHandler.GetUser)
	}

	return r
}
