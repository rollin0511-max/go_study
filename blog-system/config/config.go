package config

import (
	"gorm.io/gorm"
	"study/blog-system/models"
)

const SecretKey = "root123456@xyz" // JWT 密钥，建议使用环境变量

func InitDB(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
}
