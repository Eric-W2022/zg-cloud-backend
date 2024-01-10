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

	// 用户相关设置
	userRepo := &repository.UserRepository{DB: db}
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// 角色相关设置
	roleRepo := &repository.RoleRepository{DB: db}
	roleService := service.NewRoleService(roleRepo)
	roleHandler := handler.NewRoleHandler(roleService)

	organizationMemberRepo := &repository.OrganizationMemberRepository{DB: db}
	organizationMemberService := service.NewOrganizationMemberService(organizationMemberRepo) // Add OrganizationMemberService
	organizationMemberHandler := handler.NewOrganizationMemberHandler(organizationMemberService, roleService, userService)

	// 授权相关设置
	authService := &service.AuthService{UserRepo: userRepo}
	authHandler := handler.NewAuthHandler(authService, userService, organizationMemberService, jwtKey) // 传递 userService

	// 对话（会话）相关设置
	conversationRepo := &repository.ConversationRepository{DB: db}
	conversationService := service.NewConversationService(conversationRepo)
	conversationHandler := handler.NewConversationHandler(conversationService)

	// 消息相关设置
	messageRepo := &repository.MessageRepository{DB: db}
	messageService := service.NewMessageService(messageRepo)
	messageHandler := handler.NewMessageHandler(messageService)

	// 员工相关设置
	employeeRepo := &repository.EmployeeRepository{DB: db}
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	// 组织相关设置
	organizationRepo := &repository.OrganizationRepository{DB: db}
	organizationService := service.NewOrganizationService(organizationRepo)
	organizationHandler := handler.NewOrganizationHandler(organizationService, organizationMemberService, employeeService)

	// 组织成员相关设置
	// organizationmemberRepo := &repository.OrganizationMemberRepository{DB: db}
	// organizationmemberService := service.NewOrganizationMemberService(organizationmemberRepo)
	// organizationmemberHandler := handler.NewOrganizationMemberHandler(organizationmemberService)

	// 公开路由
	r.POST("/login", authHandler.Login)

	// 经过认证的路由组
	authRoutes := r.Group("/").Use(jwt.AuthMiddleware(jwtKey))
	{
		// 用户路由
		authRoutes.GET("/user/info", userHandler.GetUser)
		authRoutes.GET("/username/:userID", userHandler.GetNameByID)
		authRoutes.PUT("/username", userHandler.UpdateUserName)

		// 对话（会话）路由
		authRoutes.POST("/conversation", conversationHandler.CreateConversation)
		authRoutes.GET("/conversation/:conversationID", conversationHandler.GetConversation)
		authRoutes.GET("/conversations", conversationHandler.ListConversations)
		authRoutes.PUT("/conversation/:conversationID", conversationHandler.UpdateConversation)
		authRoutes.DELETE("/conversation/:conversationID", conversationHandler.DeleteConversation)

		// 消息路由
		authRoutes.POST("/message", messageHandler.CreateMessage)
		authRoutes.POST("/messages/multiple", messageHandler.CreateMultipleMessages)
		authRoutes.GET("/message/:messageID", messageHandler.GetMessage)
		authRoutes.PUT("/message/:messageID", messageHandler.UpdateMessage)
		authRoutes.DELETE("/message/:messageID", messageHandler.DeleteMessage)
		authRoutes.GET("/messages", messageHandler.ListMessages)

		//数字员工路由
		authRoutes.GET("/employee/:organizationID", employeeHandler.GetEmployee)

		//组织员工路由
		authRoutes.GET("/members/:organizationID/:userID", organizationMemberHandler.ListMembers)

		//角色路由
		authRoutes.GET("/role/:roleID", roleHandler.GetRole)

		//组织路由
		authRoutes.GET("/organization/data/:organizationID/:userID", organizationHandler.GetOrganizationData)
		authRoutes.PUT("/organization/:organizationID/:userID", organizationHandler.UpdateOrganizationName)
		authRoutes.GET("/organization/:organizationID", organizationHandler.GetOrganization)

	}

	return r
}
