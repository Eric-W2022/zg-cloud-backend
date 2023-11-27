// main.go

package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"zcloud-bg/internal/model"
	"zcloud-bg/pkg/db"
	"zcloud-bg/pkg/router"
)

func main() {
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// 读取环境变量
	mysqlUsername := os.Getenv("MYSQL_USERNAME")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlAddress := os.Getenv("MYSQL_ADDRESS")
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	// 初始化数据库连接
	database, err := db.InitDatabase(mysqlUsername, mysqlPassword, mysqlAddress)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 运行数据库迁移
	if err := database.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// 设置路由
	r := router.Setup(database, []byte(jwtSecretKey))

	// 启动 Gin 服务器
	if err := r.Run(":" + "80"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
