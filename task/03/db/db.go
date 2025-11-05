package db

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// 加载 .env 文件
	envPath := findEnvFile()
	if envPath != "" {
		err := godotenv.Load(envPath)
		if err == nil {
			fmt.Println("加载 .env 文件成功:", envPath)
		} else {
			log.Printf("加载 .env 文件失败: %v", err)
		}
	} else {
		log.Println("未找到 .env 文件，使用系统环境变量")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if user == "" || pass == "" || host == "" || port == "" || name == "" {
		log.Println("警告: 数据库环境变量不完整，可能导致连接失败")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 连接池设置
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("获取底层连接失败: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("数据库连接成功")
}

func CloseDB() {
	if DB == nil {
		return
	}

	sqlDB, err := DB.DB()

	if err != nil {
		log.Printf("获取底层连接失败: %v", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Printf("关闭数据库连接失败: %v", err)
	} else {
		fmt.Println("数据库连接已关闭")
	}
}

// findEnvFile 从当前目录向上查找 .env 文件
func findEnvFile() string {
	dir, _ := os.Getwd()

	for i := 0; i < 6; i++ { // 最多向上 3 层目录
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			return envPath
		}
		dir = filepath.Dir(dir) // 向上一级
	}
	return ""
}
