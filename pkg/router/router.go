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

// Setup initializes and returns a configured Gin router
func Setup(db *gorm.DB, jwtKey []byte) *gin.Engine {
	r := gin.Default()

	// Use CORS middleware
	r.Use(CORS())

	// User related setup
	userRepo := &repository.UserRepository{DB: db}
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Conversation related setup
	conversationRepo := &repository.ConversationRepository{DB: db}
	conversationService := service.NewConversationService(conversationRepo)
	conversationHandler := handler.NewConversationHandler(conversationService)

	// Auth related setup
	authService := &service.AuthService{UserRepo: userRepo}
	authHandler := handler.NewAuthHandler(authService, userService, jwtKey) // 传递 userService

	// Public routes
	r.POST("/login", authHandler.Login)

	// Authenticated routes group
	authRoutes := r.Group("/").Use(jwt.AuthMiddleware(jwtKey))
	{
		// User routes
		authRoutes.GET("/user/info", userHandler.GetUser)

		// Conversation routes
		authRoutes.POST("/conversation", conversationHandler.CreateConversation)
		authRoutes.GET("/conversation/:conversationID", conversationHandler.GetConversation)
		authRoutes.GET("/conversations", conversationHandler.ListConversations)
		authRoutes.PUT("/conversation/:conversationID", conversationHandler.UpdateConversation)
		authRoutes.DELETE("/conversation/:conversationID", conversationHandler.DeleteConversation)
		// Add more conversation routes as needed
	}

	return r
}
