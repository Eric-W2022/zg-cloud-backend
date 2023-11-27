// pkg/router/router.go

package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"zcloud-bg/internal/handler"
	"zcloud-bg/internal/repository"
	"zcloud-bg/internal/service"
)

// Setup 初始化并返回一个配置好的 Gin 路由器
func Setup(db *gorm.DB, jwtKey []byte) *gin.Engine {
	r := gin.Default()

	userRepo := &repository.UserRepository{DB: db}
	authService := &service.AuthService{UserRepo: userRepo}
	authHandler := handler.NewAuthHandler(authService, jwtKey)

	// 公共路由
	r.POST("/login", authHandler.Login)

	// 身份验证路由组
	// ...

	return r
}
