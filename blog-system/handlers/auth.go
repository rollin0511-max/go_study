package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"study/blog-system/config"
	"study/blog-system/models"
	"time"
)

// RegisterInput 注册输入结构体
type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

// LoginInput 登录输入结构体
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register 用户注册
func Register(c *gin.Context) {
	// 定义输入参数的接收变量
	var input RegisterInput
	// 绑定 JSON 请求体到 input 变量
	if err := c.ShouldBindJSON(&input); err != nil {
		// 绑定失败，返回错误
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	// 密码加密失败，返回错误
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	// 创建用户
	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
		Email:    input.Email,
	}

	// 从 Gin 上下文获取数据库实例
	db := c.MustGet("db").(*gorm.DB)
	// 新增用户到user表，如果失败返回错误
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}
	// 注册成功，返回成功消息
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Login 用户登录
func Login(c *gin.Context) {
	// 定义输入参数的接收变量
	var input LoginInput
	// 绑定 JSON 请求体到 input 变量
	if err := c.ShouldBindJSON(&input); err != nil {
		// 绑定失败，返回错误
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 从 Gin 上下文获取数据库实例
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	// 从user表中查询用户，如果失败返回错误
	if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		// 查询失败，返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	// 密码校验
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		// 密码校验失败，返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 生成 JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	// 签名 token
	tokenString, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		// 签名失败，返回错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	// 登录成功，返回 token
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
