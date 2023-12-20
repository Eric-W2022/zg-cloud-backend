package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"zcloud-bg/internal/model"
	"zcloud-bg/pkg/db"
	"zcloud-bg/pkg/router"
)

func main() {
	// 尝试加载 .env 文件，但如果不存在则忽略错误
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	// 读取环境变量
	env := os.Getenv("ENV")

	var mysqlAddress, mysqlUsername, mysqlPassword, mysqlDatabase string

	if env == "development" {
		mysqlAddress = os.Getenv("MYSQL_ADDRESS_DEV")
		mysqlUsername = os.Getenv("MYSQL_USERNAME_DEV")
		mysqlPassword = os.Getenv("MYSQL_PASSWORD_DEV")
		mysqlDatabase = os.Getenv("MYSQL_DATABASE_DEV")
	} else {
		mysqlAddress = os.Getenv("MYSQL_ADDRESS_PROD")
		mysqlUsername = os.Getenv("MYSQL_USERNAME_PROD")
		mysqlPassword = os.Getenv("MYSQL_PASSWORD_PROD")
		mysqlDatabase = os.Getenv("MYSQL_DATABASE_PROD")
	}

	// 初始化数据库连接
	database, err := db.InitDatabase(mysqlUsername, mysqlPassword, mysqlAddress, mysqlDatabase)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 运行数据库迁移
	if err := database.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// 设置路由
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	r := router.Setup(database, []byte(jwtSecretKey))

	// 启动 Gin 服务器
	appPort := os.Getenv("APP_PORT")
	if err := r.Run(":" + appPort); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
