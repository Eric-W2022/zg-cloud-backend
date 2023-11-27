// pkg/router/router.go

package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"zcloud-bg/internal/handler"
	"zcloud-bg/internal/repository"
	"zcloud-bg/internal/service"
	"zcloud-bg/pkg/jwt"
)

// Setup 初始化并返回一个配置好的 Gin 路由器
func Setup(db *gorm.DB, jwtKey []byte) *gin.Engine {
	r := gin.Default()

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
