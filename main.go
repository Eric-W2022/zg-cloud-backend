package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("iry") // 用于签名JWT的密钥
var db *sql.DB             // 数据库连接池

// 创建 JWT claims 结构
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 主函数
func main() {
	var err error
	// 初始化数据库连接
	db, err = sql.Open("mysql", "root:STC89c51$@tcp(gz-cdb-6kgcteld.sql.tencentcdb.com:63181)/zg-cloud")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// 测试数据库连接
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	r := gin.Default()

	// 登录路由
	r.POST("/login", login)
	// 新增检查令牌路由
	r.GET("/check-token", checkToken)

	// 启动服务
	r.Run(":80")
}

// 登录处理器
func login(c *gin.Context) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	// 绑定JSON到creds
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求格式不正确"})
		return
	}

	// 从数据库中获取用户密码
	var hashedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", creds.Username).Scan(&hashedPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "不存在该用户"})
		return
	}

	// 验证密码
	if !checkPasswordHash(creds.Password, hashedPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "登录失败"})
		return
	}

	// 创建JWT
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法生成令牌"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// 检查密码哈希
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func checkToken(c *gin.Context) {
	tokenString := c.Query("token") // 从请求中获取令牌

	// 解析令牌
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// 令牌有效
		c.JSON(http.StatusOK, gin.H{"status": "有效", "expires_at": claims.ExpiresAt})
	} else {
		// 令牌无效或已过期
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效或过期的令牌", "details": err})
	}
}
